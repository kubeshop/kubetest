package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	testsv1 "github.com/kubeshop/testkube-operator/apis/tests/v1"
	"github.com/kubeshop/testkube/internal/pkg/api/datefilter"
	"github.com/kubeshop/testkube/internal/pkg/api/repository/testresult"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	testsmapper "github.com/kubeshop/testkube/pkg/mapper/tests"
	"github.com/kubeshop/testkube/pkg/rand"
)

// GetTestHandler for getting test object
func (s TestKubeAPI) CreateTestHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request testkube.TestUpsertRequest
		err := c.BodyParser(&request)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		test := mapTestUpsertRequestToTestCRD(request)
		created, err := s.TestsClient.Create(&test)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		c.Status(201)
		return c.JSON(created)
	}
}

// GetTestHandler for getting test object
func (s TestKubeAPI) GetTestHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		namespace := c.Query("namespace", "testkube")
		crTest, err := s.TestsClient.Get(namespace, name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		test := testsmapper.MapCRToAPI(*crTest)

		return c.JSON(test)
	}
}

// DeleteTestHandler for deleting a test with id
func (s TestKubeAPI) DeleteTestHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		namespace := c.Query("namespace", "testkube")
		err := s.TestsClient.Delete(namespace, name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

// DeleteTestsHandler for deleting all Tests
func (s TestKubeAPI) DeleteTestsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		namespace := c.Query("namespace", "testkube")
		err := s.TestsClient.DeleteAll(namespace)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

// ListTestsHandler for getting list of all available tests
func (s TestKubeAPI) ListTestsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s.Log.Debug("Getting scripts list")
		namespace := c.Query("namespace", "testkube")

		rawTags := c.Query("tags")
		var tags []string
		if rawTags != "" {
			tags = strings.Split(rawTags, ",")
		}

		crTests, err := s.TestsClient.List(namespace, tags)

		if err != nil {
			return s.Error(c, http.StatusInternalServerError, err)
		}

		search := c.Query("textSearch")
		if search != "" {
			// filter items array
			for i := len(crTests.Items) - 1; i >= 0; i-- {
				if !strings.Contains(crTests.Items[i].Name, search) {
					crTests.Items = append(crTests.Items[:i], crTests.Items[i+1:]...)
				}
			}
		}

		tests := testsmapper.MapTestListKubeToAPI(*crTests)

		return c.JSON(tests)
	}
}

func (s TestKubeAPI) ExecuteTestHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		name := c.Params("id")
		namespace := c.Query("namespace", "testkube")
		s.Log.Debugw("getting test", "name", name)

		crTest, err := s.TestsClient.Get(namespace, name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		var request testkube.TestExecutionRequest
		err = c.BodyParser(&request)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, fmt.Errorf("test execution request body invalid: %w", err))
		}

		test := testsmapper.MapCRToAPI(*crTest)
		s.Log.Debugw("executing test", "name", name, "test", test, "cr", crTest)
		results := s.executeTest(ctx, request, test)

		c.Response().SetStatusCode(fiber.StatusCreated)
		return c.JSON(results)
	}
}

func (s TestKubeAPI) ListTestExecutionsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		filter := getExecutionsFilterFromRequest(c)

		executionsTotals, err := s.TestExecutionResults.GetExecutionsTotals(ctx, filter)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}
		allExecutionsTotals, err := s.TestExecutionResults.GetExecutionsTotals(ctx)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		executions, err := s.TestExecutionResults.GetExecutions(ctx, filter)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		return c.JSON(testkube.TestExecutionsResult{
			Totals:   &allExecutionsTotals,
			Filtered: &executionsTotals,
			Results:  mapToTestExecutionSummary(executions),
		})
	}
}

func (s TestKubeAPI) GetTestExecutionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		id := c.Params("executionID")
		execution, err := s.TestExecutionResults.Get(ctx, id)

		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		return c.JSON(execution)
	}
}

func (s TestKubeAPI) executeTest(ctx context.Context, request testkube.TestExecutionRequest, test testkube.Test) (testExecution testkube.TestExecution) {
	s.Log.Debugw("Got test to execute", "test", test)

	testExecution = testkube.NewStartedTestExecution(test, request)
	s.TestExecutionResults.Insert(ctx, testExecution)

	go func(testExecution testkube.TestExecution) {

		defer func(testExecution *testkube.TestExecution) {
			duration := testExecution.CalculateDuration()
			testExecution.EndTime = time.Now()
			testExecution.Duration = duration.String()

			err := s.TestExecutionResults.EndExecution(ctx, testExecution.Id, testExecution.EndTime, duration)
			if err != nil {
				s.Log.Errorw("error setting end time", "error", err.Error())
			}
		}(&testExecution)

		hasFailedSteps := false
		for i := range testExecution.StepResults {

			// start execution of given step
			testExecution.StepResults[i].Execution.ExecutionResult.InProgress()
			s.TestExecutionResults.Update(ctx, testExecution)

			s.executeTestStep(ctx, testExecution, &testExecution.StepResults[i])
			if testExecution.StepResults[i].IsFailed() {
				hasFailedSteps = true
				if testExecution.StepResults[i].Step.StopTestOnFailure {
					break
				}
			}

			s.TestExecutionResults.Update(ctx, testExecution)
		}

		testExecution.Status = testkube.TestStatusSuccess
		if hasFailedSteps {
			testExecution.Status = testkube.TestStatusError
		}

		s.TestExecutionResults.Update(ctx, testExecution)

	}(testExecution)

	return

}

func (s TestKubeAPI) executeTestStep(ctx context.Context, testExecution testkube.TestExecution, result *testkube.TestStepExecutionResult) {

	var testName string
	if testExecution.Test != nil {
		testName = testExecution.Test.Name
	}

	step := result.Step

	l := s.Log.With("type", step.Type(), "testName", testName, "name", step.FullName())

	switch step.Type() {

	case testkube.TestStepTypeExecuteScript:
		executeScriptStep := step.Execute
		options, err := s.GetExecuteOptions(executeScriptStep.Namespace, executeScriptStep.Name, testkube.ExecutionRequest{
			Name:      fmt.Sprintf("%s-%s-%s", testName, executeScriptStep.Name, rand.String(5)),
			Namespace: executeScriptStep.Namespace,
			Params:    testExecution.Params,
		})

		if err != nil {
			result.Err(err)
		}

		l.Debug("executing script", "params", testExecution.Params)
		options.Sync = true
		execution := s.executeScript(ctx, options)
		result.Execution = &execution

	case testkube.TestStepTypeDelay:
		l.Debug("delaying execution")
		time.Sleep(time.Millisecond * time.Duration(step.Delay.Duration))
		result.Execution.ExecutionResult.Success()

	default:
		result.Err(fmt.Errorf("can't find handler for execution step type: '%v'", step.Type()))
	}
}

func getExecutionsFilterFromRequest(c *fiber.Ctx) testresult.Filter {

	filter := testresult.NewExecutionsFilter()
	name := c.Query("id", "")
	if name != "" {
		filter = filter.WithName(name)
	}

	textSearch := c.Query("textSearch", "")
	if textSearch != "" {
		filter = filter.WithTextSearch(textSearch)
	}

	page, err := strconv.Atoi(c.Query("page", ""))
	if err == nil {
		filter = filter.WithPage(page)
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", ""))
	if err == nil && pageSize != 0 {
		filter = filter.WithPageSize(pageSize)
	}

	status := c.Query("status", "")
	if status != "" {
		filter = filter.WithStatus(testkube.ExecutionStatus(status))
	}

	dFilter := datefilter.NewDateFilter(c.Query("startDate", ""), c.Query("endDate", ""))
	if dFilter.IsStartValid {
		filter = filter.WithStartDate(dFilter.Start)
	}

	if dFilter.IsEndValid {
		filter = filter.WithEndDate(dFilter.End)
	}

	return filter
}

func mapToTestExecutionSummary(executions []testkube.TestExecution) []testkube.TestExecutionSummary {
	result := make([]testkube.TestExecutionSummary, len(executions))

	for i, execution := range executions {
		executionsSummary := make([]testkube.TestStepExecutionSummary, len(execution.StepResults))
		for j, stepResult := range execution.StepResults {
			executionsSummary[j] = mapStepResultToExecutionSummary(stepResult)
		}

		result[i] = testkube.TestExecutionSummary{
			Id:        execution.Id,
			Name:      execution.Name,
			TestName:  execution.Test.Name,
			Status:    execution.Status,
			StartTime: execution.StartTime,
			EndTime:   execution.EndTime,
			Duration:  execution.Duration,
			Execution: executionsSummary,
		}
	}

	return result
}

func mapStepResultToExecutionSummary(r testkube.TestStepExecutionResult) testkube.TestStepExecutionSummary {
	var id, scriptName, name string
	var status *testkube.ExecutionStatus = testkube.ExecutionStatusSuccess
	var stepType *testkube.TestStepType

	if r.Script != nil {
		scriptName = r.Script.Name
	}

	if r.Execution != nil {
		id = r.Execution.Id
		if r.Execution.ExecutionResult != nil {
			status = r.Execution.ExecutionResult.Status
		}
	}

	if r.Step != nil {
		stepType = r.Step.Type()
		name = r.Step.FullName()
	}

	return testkube.TestStepExecutionSummary{
		Id:         id,
		Name:       name,
		ScriptName: scriptName,
		Status:     status,
		Type_:      stepType,
	}
}

func mapTestUpsertRequestToTestCRD(request testkube.TestUpsertRequest) testsv1.Test {
	return testsv1.Test{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
		},
		Spec: testsv1.TestSpec{
			Repeats:     int(request.Repeats),
			Description: request.Description,
			Tags:        request.Tags,
			Before:      mapTestStepsToCRD(request.Before),
			Steps:       mapTestStepsToCRD(request.Steps),
			After:       mapTestStepsToCRD(request.After),
		},
	}
}

func mapTestStepsToCRD(steps []testkube.TestStep) (out []testsv1.TestStepSpec) {
	for _, step := range steps {
		out = append(out, mapTestStepToCRD(step))
	}

	return
}

func mapTestStepToCRD(step testkube.TestStep) (stepSpec testsv1.TestStepSpec) {
	switch step.Type() {

	case testkube.TestStepTypeDelay:
		stepSpec.Delay = &testsv1.TestStepDelay{
			Duration: step.Delay.Duration,
		}

	case testkube.TestStepTypeExecuteScript:
		s := step.Execute
		stepSpec.Execute = &testsv1.TestStepExecute{
			Namespace: s.Namespace,
			Name:      s.Name,
			// TODO move StopOnFailure level up in operator model to mimic this one
			StopOnFailure: step.StopTestOnFailure,
		}
	}

	return
}
