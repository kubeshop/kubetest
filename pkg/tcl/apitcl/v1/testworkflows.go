// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//	https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	"github.com/kubeshop/testkube/internal/common"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/tcl/expressionstcl"
	testworkflowmappers "github.com/kubeshop/testkube/pkg/tcl/mapperstcl/testworkflows"
	"github.com/kubeshop/testkube/pkg/tcl/testworkflowstcl/testworkflowcontroller"
	"github.com/kubeshop/testkube/pkg/tcl/testworkflowstcl/testworkflowprocessor"
	"github.com/kubeshop/testkube/pkg/tcl/testworkflowstcl/testworkflowresolver"
)

func (s *apiTCL) ListTestWorkflowsHandler() fiber.Handler {
	errPrefix := "failed to list test workflows"
	return func(c *fiber.Ctx) (err error) {
		workflows, err := s.getFilteredTestWorkflowList(c)
		if err != nil {
			return s.BadGateway(c, errPrefix, "client problem", err)
		}
		err = SendResourceList(c, "TestWorkflow", testworkflowsv1.GroupVersion, testworkflowmappers.MapTestWorkflowKubeToAPI, workflows.Items...)
		if err != nil {
			return s.InternalError(c, errPrefix, "serialization problem", err)
		}
		return
	}
}

func (s *apiTCL) GetTestWorkflowHandler() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		name := c.Params("id")
		errPrefix := fmt.Sprintf("failed to get test workflow '%s'", name)
		workflow, err := s.TestWorkflowsClient.Get(name)
		if err != nil {
			return s.ClientError(c, errPrefix, err)
		}
		err = SendResource(c, "TestWorkflow", testworkflowsv1.GroupVersion, testworkflowmappers.MapKubeToAPI, workflow)
		if err != nil {
			return s.InternalError(c, errPrefix, "serialization problem", err)
		}
		return
	}
}

func (s *apiTCL) DeleteTestWorkflowHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		errPrefix := fmt.Sprintf("failed to delete test workflow '%s'", name)
		err := s.TestWorkflowsClient.Delete(name)
		s.Metrics.IncDeleteTestWorkflow(err)
		if err != nil {
			return s.ClientError(c, errPrefix, err)
		}
		skipExecutions := c.Query("skipDeleteExecutions", "")
		if skipExecutions != "true" {
			err = s.TestWorkflowResults.DeleteByTestWorkflow(context.Background(), name)
		}
		return c.SendStatus(http.StatusNoContent)
	}
}

func (s *apiTCL) DeleteTestWorkflowsHandler() fiber.Handler {
	errPrefix := "failed to delete test workflows"
	return func(c *fiber.Ctx) error {
		selector := c.Query("selector")
		workflows, err := s.TestWorkflowsClient.List(selector)
		if err != nil {
			return s.BadGateway(c, errPrefix, "client problem", err)
		}

		// Delete
		err = s.TestWorkflowsClient.DeleteByLabels(selector)
		if err != nil {
			return s.ClientError(c, errPrefix, err)
		}

		// Mark as deleted
		for range workflows.Items {
			s.Metrics.IncDeleteTestWorkflow(err)
		}

		// Delete the executions
		skipExecutions := c.Query("skipDeleteExecutions", "")
		if skipExecutions != "true" {
			names := common.MapSlice(workflows.Items, func(t testworkflowsv1.TestWorkflow) string {
				return t.Name
			})
			err = s.TestWorkflowResults.DeleteByTestWorkflows(context.Background(), names)
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (s *apiTCL) CreateTestWorkflowHandler() fiber.Handler {
	errPrefix := "failed to create test workflow"
	return func(c *fiber.Ctx) (err error) {
		// Deserialize resource
		obj := new(testworkflowsv1.TestWorkflow)
		if HasYAML(c) {
			err = common.DeserializeCRD(obj, c.Body())
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
		} else {
			var v *testkube.TestWorkflow
			err = c.BodyParser(&v)
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
			obj = testworkflowmappers.MapAPIToKube(v)
		}

		// Validate resource
		if obj == nil || obj.Name == "" {
			return s.BadRequest(c, errPrefix, "invalid body", errors.New("name is required"))
		}
		obj.Namespace = s.Namespace

		// Create the resource
		obj, err = s.TestWorkflowsClient.Create(obj)
		s.Metrics.IncCreateTestWorkflow(err)
		if err != nil {
			return s.BadRequest(c, errPrefix, "client error", err)
		}

		err = SendResource(c, "TestWorkflow", testworkflowsv1.GroupVersion, testworkflowmappers.MapKubeToAPI, obj)
		if err != nil {
			return s.InternalError(c, errPrefix, "serialization problem", err)
		}
		return
	}
}

func (s *apiTCL) UpdateTestWorkflowHandler() fiber.Handler {
	errPrefix := "failed to update test workflow"
	return func(c *fiber.Ctx) (err error) {
		name := c.Params("id")

		// Deserialize resource
		obj := new(testworkflowsv1.TestWorkflow)
		if HasYAML(c) {
			err = common.DeserializeCRD(obj, c.Body())
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
		} else {
			var v *testkube.TestWorkflow
			err = c.BodyParser(&v)
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
			obj = testworkflowmappers.MapAPIToKube(v)
		}

		// Read existing resource
		workflow, err := s.TestWorkflowsClient.Get(name)
		if err != nil {
			return s.ClientError(c, errPrefix, err)
		}

		// Validate resource
		if obj == nil {
			return s.BadRequest(c, errPrefix, "invalid body", errors.New("body is required"))
		}
		obj.Namespace = workflow.Namespace
		obj.Name = workflow.Name
		obj.ResourceVersion = workflow.ResourceVersion

		// Update the resource
		obj, err = s.TestWorkflowsClient.Update(obj)
		s.Metrics.IncUpdateTestWorkflow(err)
		if err != nil {
			return s.BadRequest(c, errPrefix, "client error", err)
		}

		err = SendResource(c, "TestWorkflow", testworkflowsv1.GroupVersion, testworkflowmappers.MapKubeToAPI, obj)
		if err != nil {
			return s.InternalError(c, errPrefix, "serialization problem", err)
		}
		return
	}
}

func (s *apiTCL) PreviewTestWorkflowHandler() fiber.Handler {
	errPrefix := "failed to resolve test workflow"
	return func(c *fiber.Ctx) (err error) {
		// Deserialize resource
		obj := new(testworkflowsv1.TestWorkflow)
		if HasYAML(c) {
			err = common.DeserializeCRD(obj, c.Body())
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
		} else {
			var v *testkube.TestWorkflow
			err = c.BodyParser(&v)
			if err != nil {
				return s.BadRequest(c, errPrefix, "invalid body", err)
			}
			obj = testworkflowmappers.MapAPIToKube(v)
		}

		// Validate resource
		if obj == nil {
			return s.BadRequest(c, errPrefix, "invalid body", errors.New("name is required"))
		}
		obj.Namespace = s.Namespace

		// Fetch the templates
		tpls := testworkflowresolver.ListTemplates(obj)
		tplsMap := make(map[string]testworkflowsv1.TestWorkflowTemplate, len(tpls))
		for name := range tpls {
			tpl, err := s.TestWorkflowTemplatesClient.Get(name)
			if err != nil {
				return s.BadRequest(c, errPrefix, "fetching error", err)
			}
			tplsMap[name] = *tpl
		}

		// Resolve the TestWorkflow
		err = testworkflowresolver.ApplyTemplates(obj, tplsMap)
		if err != nil {
			return s.BadRequest(c, errPrefix, "resolving error", err)
		}

		err = SendResource(c, "TestWorkflow", testworkflowsv1.GroupVersion, testworkflowmappers.MapKubeToAPI, obj)
		if err != nil {
			return s.InternalError(c, errPrefix, "serialization problem", err)
		}
		return
	}
}

// TODO: Add metrics
func (s *apiTCL) ExecuteTestWorkflowHandler() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		ctx := c.Context()
		name := c.Params("id")
		errPrefix := fmt.Sprintf("failed to execute test workflow '%s'", name)
		workflow, err := s.TestWorkflowsClient.Get(name)
		if err != nil {
			return s.ClientError(c, errPrefix, err)
		}

		// Delete unnecessary data
		delete(workflow.Annotations, "kubectl.kubernetes.io/last-applied-configuration")

		// Preserve initial workflow
		initialWorkflow := workflow.DeepCopy()

		// Load the execution request
		var request testkube.TestWorkflowExecutionRequest
		err = c.BodyParser(&request)
		if err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
			return s.BadRequest(c, errPrefix, "invalid body", err)
		}

		// Fetch the templates
		tpls := testworkflowresolver.ListTemplates(workflow)
		tplsMap := make(map[string]testworkflowsv1.TestWorkflowTemplate, len(tpls))
		for tplName := range tpls {
			tpl, err := s.TestWorkflowTemplatesClient.Get(tplName)
			if err != nil {
				return s.BadRequest(c, errPrefix, "fetching error", err)
			}
			tplsMap[tplName] = *tpl
		}

		// Apply the configuration
		_, err = testworkflowresolver.ApplyWorkflowConfig(workflow, testworkflowmappers.MapConfigValueAPIToKube(request.Config))
		if err != nil {
			return s.BadRequest(c, errPrefix, "configuration", err)
		}

		// Resolve the TestWorkflow
		err = testworkflowresolver.ApplyTemplates(workflow, tplsMap)
		if err != nil {
			return s.BadRequest(c, errPrefix, "resolving error", err)
		}

		// Build the basic Execution data
		id := primitive.NewObjectID().Hex()
		now := time.Now()
		machine := expressionstcl.NewMachine().
			Register("execution.id", id)

		// Preserve resolved TestWorkflow
		resolvedWorkflow := workflow.DeepCopy()

		// Process the TestWorkflow
		bundle, err := testworkflowprocessor.NewFullFeatured(s.ImageInspector).
			Bundle(c.Context(), workflow, machine)
		if err != nil {
			return s.BadRequest(c, errPrefix, "processing error", err)
		}

		// Load execution identifier data
		// TODO: Consider if that should not be shared (as now it is between Tests and Test Suites)
		number, _ := s.ExecutionResults.GetNextExecutionNumber(context.Background(), workflow.Name)
		executionName := request.Name
		if executionName == "" {
			executionName = fmt.Sprintf("%s-%d", workflow.Name, number)
		}

		// Ensure it is unique name
		// TODO: Consider if we shouldn't make name unique across all TestWorkflows
		next, _ := s.TestWorkflowResults.GetByNameAndTestWorkflow(ctx, executionName, workflow.Name)
		if next.Name == executionName {
			return s.BadRequest(c, errPrefix, "execution name already exists", errors.New(executionName))
		}

		// Build Execution entity
		// TODO: Consider storing "config" as well
		execution := testkube.TestWorkflowExecution{
			Id:          id,
			Name:        executionName,
			Number:      number,
			ScheduledAt: now,
			StatusAt:    now,
			Signature:   testworkflowprocessor.MapSignatureListToInternal(bundle.Signature),
			Result: &testkube.TestWorkflowResult{
				Status:          common.Ptr(testkube.QUEUED_TestWorkflowStatus),
				PredictedStatus: common.Ptr(testkube.PASSED_TestWorkflowStatus),
				Initialization: &testkube.TestWorkflowStepResult{
					Status: common.Ptr(testkube.QUEUED_TestWorkflowStepStatus),
				},
				Steps: testworkflowprocessor.MapSignatureListToStepResults(bundle.Signature),
			},
			Output:           []testkube.TestWorkflowOutput{},
			Workflow:         testworkflowmappers.MapKubeToAPI(initialWorkflow),
			ResolvedWorkflow: testworkflowmappers.MapKubeToAPI(resolvedWorkflow),
		}
		err = s.TestWorkflowResults.Insert(ctx, execution)
		if err != nil {
			return s.InternalError(c, errPrefix, "inserting execution to storage", err)
		}

		// Inform about execution start
		s.Events.Notify(testkube.NewEventQueueTestWorkflow(&execution))

		// Deploy the resources
		for _, item := range bundle.Secrets {
			_, err = s.Clientset.CoreV1().Secrets(s.Namespace).Create(context.Background(), &item, metav1.CreateOptions{})
			if err != nil {
				execution.Result.Fatal(err)
				err = s.TestWorkflowResults.UpdateResult(context.Background(), execution.Id, execution.Result)
				go testworkflowcontroller.Cleanup(context.Background(), s.Clientset, s.Namespace, execution.Id)
				if err != nil {
					return s.BadRequest(c, errPrefix, "saving execution", err)
				}
				return c.JSON(execution)
			}
		}
		for _, item := range bundle.ConfigMaps {
			_, err = s.Clientset.CoreV1().ConfigMaps(s.Namespace).Create(context.Background(), &item, metav1.CreateOptions{})
			if err != nil {
				execution.Result.Fatal(err)
				err = s.TestWorkflowResults.UpdateResult(context.Background(), execution.Id, execution.Result)
				go testworkflowcontroller.Cleanup(context.Background(), s.Clientset, s.Namespace, execution.Id)
				if err != nil {
					return s.BadRequest(c, errPrefix, "saving execution", err)
				}
				return c.JSON(execution)
			}
		}
		_, err = s.Clientset.BatchV1().Jobs(s.Namespace).Create(context.Background(), &bundle.Job, metav1.CreateOptions{})
		if err != nil {
			execution.Result.Fatal(err)
			err = s.TestWorkflowResults.UpdateResult(context.Background(), execution.Id, execution.Result)
			go testworkflowcontroller.Cleanup(context.Background(), s.Clientset, s.Namespace, execution.Id)
			if err != nil {
				return s.BadRequest(c, errPrefix, "saving execution", err)
			}
			return c.JSON(execution)
		}

		// Start to control the results
		// TODO: Move it outside of the API when persistence will be there
		go func() {
			ctrl, err := testworkflowcontroller.New(context.Background(), s.Clientset, s.Namespace, execution.Id, execution.ScheduledAt)
			if err != nil {
				execution.Result.Fatal(errors.Wrap(err, "reading job"))
				_ = s.TestWorkflowResults.UpdateResult(context.Background(), execution.Id, execution.Result)
				testworkflowcontroller.Cleanup(context.Background(), s.Clientset, s.Namespace, execution.Id)
				return
			}
			for v := range ctrl.Watch(context.Background()).Stream(context.Background()).Channel() {
				if v.Error != nil {
					continue
				}
				if v.Value.Output != nil {
					execution.Output = append(execution.Output, *v.Value.Output.ToInternal())
					// TODO: Consider AppendOutput ($push) instead
					_ = s.TestWorkflowResults.UpdateOutput(context.Background(), execution.Id, execution.Output)
				}
				if v.Value.Result != nil {
					execution.Result = v.Value.Result
					if execution.Result.IsFinished() {
						execution.StatusAt = execution.Result.FinishedAt
					}
					_ = s.TestWorkflowResults.UpdateResult(context.Background(), execution.Id, execution.Result)
				}
			}
			if execution.Result.IsFinished() {
				if execution.Result.IsPassed() {
					s.Events.Notify(testkube.NewEventEndTestWorkflowSuccess(&execution))
				} else if execution.Result.IsAborted() {
					s.Events.Notify(testkube.NewEventEndTestWorkflowAborted(&execution))
				} else {
					s.Events.Notify(testkube.NewEventEndTestWorkflowFailed(&execution))
				}
			}
			testworkflowcontroller.Cleanup(context.Background(), s.Clientset, s.Namespace, execution.Id)
		}()

		return c.JSON(execution)
	}
}

func (s *apiTCL) getFilteredTestWorkflowList(c *fiber.Ctx) (*testworkflowsv1.TestWorkflowList, error) {
	crWorkflows, err := s.TestWorkflowsClient.List(c.Query("selector"))
	if err != nil {
		return nil, err
	}

	search := c.Query("textSearch")
	if search != "" {
		// filter items array
		for i := len(crWorkflows.Items) - 1; i >= 0; i-- {
			if !strings.Contains(crWorkflows.Items[i].Name, search) {
				crWorkflows.Items = append(crWorkflows.Items[:i], crWorkflows.Items[i+1:]...)
			}
		}
	}

	return crWorkflows, nil
}
