// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//	https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package testworkflowcontroller

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kubeshop/testkube/internal/common"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/tcl/testworkflowstcl/testworkflowprocessor"
)

const (
	JobRetrievalTimeout = 3 * time.Second
)

type Controller interface {
	Abort(ctx context.Context) error
	Cleanup(ctx context.Context) error
	Watch(ctx context.Context) Watcher[Notification]
}

func New(parentCtx context.Context, clientSet kubernetes.Interface, namespace, id string, scheduledAt time.Time) (Controller, error) {
	// Create local context for stopping all the processes
	ctx, ctxCancel := context.WithCancel(parentCtx)

	// Optimistically, start watching all the resources
	job := WatchJob(ctx, clientSet, namespace, id, 1)
	jobEvents := WatchJobEvents(ctx, clientSet, namespace, id, -1)
	pod := WatchMainPod(ctx, clientSet, namespace, id, 1)
	podEvents := WatchPodEventsByPodWatcher(ctx, clientSet, namespace, pod, -1)

	// Ensure the main Job exists in the cluster,
	// and obtain the signature
	var sig []testworkflowprocessor.Signature
	var err error
	select {
	case j := <-job.Any(ctx):
		if j.Error != nil {
			ctxCancel()
			return nil, j.Error
		}
		sig, err = testworkflowprocessor.GetSignatureFromJSON([]byte(j.Value.Annotations[testworkflowprocessor.SignatureAnnotationName]))
		if err != nil {
			ctxCancel()
			return nil, errors.Wrap(err, "invalid job signature")
		}
	case <-time.After(JobRetrievalTimeout):
		ctxCancel()
		return nil, errors.New("timeout retrieving job")
	}

	// Build accessible controller
	return &controller{
		id:          id,
		namespace:   namespace,
		scheduledAt: scheduledAt,
		signature:   sig,
		clientSet:   clientSet,
		ctx:         ctx,
		ctxCancel:   ctxCancel,
		job:         job,
		jobEvents:   jobEvents,
		pod:         pod,
		podEvents:   podEvents,
	}, nil
}

type controller struct {
	id          string
	namespace   string
	scheduledAt time.Time
	signature   []testworkflowprocessor.Signature
	clientSet   kubernetes.Interface
	ctx         context.Context
	ctxCancel   context.CancelFunc
	job         Watcher[*batchv1.Job]
	jobEvents   Watcher[*corev1.Event]
	pod         Watcher[*corev1.Pod]
	podEvents   Watcher[*corev1.Event]
}

func (c *controller) Abort(ctx context.Context) error {
	return c.Cleanup(ctx)
}

func (c *controller) Cleanup(ctx context.Context) error {
	return Cleanup(ctx, c.clientSet, c.namespace, c.id)
}

func (c *controller) Watch(parentCtx context.Context) Watcher[Notification] {
	ctx, ctxCancel := context.WithCancel(parentCtx)
	w := newWatcher[Notification](ctx, 0)

	go func() {
		defer w.Close()
		defer ctxCancel()

		sig := make([]testkube.TestWorkflowSignature, len(c.signature))
		for i, s := range c.signature {
			sig[i] = s.ToInternal()
		}

		// Build initial result
		result := testkube.TestWorkflowResult{
			Status:          common.Ptr(testkube.QUEUED_TestWorkflowStatus),
			PredictedStatus: common.Ptr(testkube.PASSED_TestWorkflowStatus),
			Initialization: &testkube.TestWorkflowStepResult{
				Status: common.Ptr(testkube.QUEUED_TestWorkflowStepStatus),
			},
			Steps: testworkflowprocessor.MapSignatureListToStepResults(c.signature),
		}

		// Emit initial empty result
		w.SendValue(Notification{Result: result.Clone()})

		// Wait for the pod creation
		for v := range WatchJobPreEvents(ctx, c.jobEvents, 0).Stream(ctx).Channel() {
			if v.Error != nil {
				w.SendError(v.Error)
				continue
			}
			if v.Value.Reason == "SuccessfulCreate" {
				result.QueuedAt = v.Value.CreationTimestamp.Time
			}
			if v.Value.Type == "Normal" {
				continue
			}
			w.SendValue(Notification{
				Timestamp: v.Value.CreationTimestamp.Time,
				Log:       fmt.Sprintf("%s (%s) %s\n", v.Value.CreationTimestamp.Time.Format(time.RFC3339Nano), v.Value.Reason, v.Value.Message),
			})
		}

		// Emit the result with queue time
		if result.QueuedAt.IsZero() {
			w.SendError(errors.New("job is in unknown state"))
			return
		}
		w.SendValue(Notification{Result: result.Clone()})

		// Wait for the pod initialization
		for v := range WatchPodPreEvents(ctx, c.podEvents, 0).Stream(ctx).Channel() {
			if v.Error != nil {
				w.SendError(v.Error)
				continue
			}
			if v.Value.Reason == "Scheduled" {
				result.StartedAt = v.Value.CreationTimestamp.Time
				result.Status = common.Ptr(testkube.RUNNING_TestWorkflowStatus)
			}
			if v.Value.Type == "Normal" {
				continue
			}
			w.SendValue(Notification{
				Timestamp: v.Value.CreationTimestamp.Time,
				Log:       fmt.Sprintf("%s (%s) %s\n", v.Value.CreationTimestamp.Time.Format(time.RFC3339Nano), v.Value.Reason, v.Value.Message),
			})
		}

		// Emit the result with start time
		if result.StartedAt.IsZero() {
			w.SendError(errors.New("pod is in unknown state"))
			return
		}
		w.SendValue(Notification{Result: result.Clone()})

		// Wait for the initialization container
		for v := range WatchContainerPreEvents(ctx, c.podEvents, "tktw-init", 0).Stream(ctx).Channel() {
			if v.Error != nil {
				w.SendError(v.Error)
				continue
			}
			if v.Value.Reason == "Created" {
				result.Initialization.QueuedAt = v.Value.CreationTimestamp.Time
			} else if v.Value.Reason == "Started" {
				result.Initialization.StartedAt = v.Value.CreationTimestamp.Time
				result.Initialization.Status = common.Ptr(testkube.RUNNING_TestWorkflowStepStatus)
			}
			if v.Value.Type == "Normal" {
				continue
			}
			w.SendValue(Notification{
				Timestamp: v.Value.CreationTimestamp.Time,
				Log:       fmt.Sprintf("%s (%s) %s\n", v.Value.CreationTimestamp.Time.Format(time.RFC3339Nano), v.Value.Reason, v.Value.Message),
			})
		}

		// Emit the result with start time
		if result.Initialization.StartedAt.IsZero() {
			w.SendError(errors.New("init container is in unknown state"))
			return
		}
		w.SendValue(Notification{Result: result.Clone()})

		// Watch the initialization container logs
		pod := (<-c.pod.Any(ctx)).Value
		for v := range WatchContainerLogs(ctx, c.clientSet, c.podEvents, c.namespace, pod.Name, "tktw-init").Stream(ctx).Channel() {
			if v.Error != nil {
				w.SendError(v.Error)
				continue
			}
			// TODO: Calibrate clock with v.Value.Hint or just first/last timestamp here
			w.SendValue(Notification{
				Timestamp: v.Value.Time,
				Log:       fmt.Sprintf("%s %s\n", v.Value.Time.Format(time.RFC3339Nano), string(v.Value.Log)),
			})
		}

		// Update the initialization container status
		status, err := GetFinalContainerResult(ctx, c.pod, "tktw-init")
		if err != nil {
			w.SendError(err)
			return
		}
		result.Initialization.FinishedAt = status.FinishedAt
		result.Initialization.Status = common.Ptr(status.Status)
		if status.Status != testkube.PASSED_TestWorkflowStepStatus {
			result.Status = common.Ptr(testkube.FAILED_TestWorkflowStatus)
			result.PredictedStatus = result.Status
		}
		w.SendValue(Notification{Result: result.Clone()})

		// Cancel when the initialization has failed
		if status.Status != testkube.PASSED_TestWorkflowStepStatus {
			return
		}

		// Watch each of the containers
		lastTs := result.Initialization.FinishedAt
		for _, container := range append(pod.Spec.InitContainers, pod.Spec.Containers...) {
			// Ignore not-standard TestWorkflow containers
			if _, ok := result.Steps[container.Name]; !ok {
				continue
			}

			// Send the step queued time
			stepResult := result.Steps[container.Name]
			stepResult = result.UpdateStepResult(sig, container.Name, testkube.TestWorkflowStepResult{
				QueuedAt: lastTs.UTC(),
			})
			w.SendValue(Notification{Result: result.Clone()})

			// Watch for the container events
			lastEvTs := time.Time{}
			for v := range WatchContainerPreEvents(ctx, c.podEvents, container.Name, 0).Stream(ctx).Channel() {
				if v.Error != nil {
					w.SendError(v.Error)
					continue
				}
				if lastEvTs.Before(v.Value.CreationTimestamp.Time) {
					lastEvTs = v.Value.CreationTimestamp.Time
				}
				if v.Value.Reason == "Created" {
					stepResult = result.UpdateStepResult(sig, container.Name, testkube.TestWorkflowStepResult{
						QueuedAt: v.Value.CreationTimestamp.Time.UTC(),
					})
				} else if v.Value.Reason == "Started" {
					stepResult = result.UpdateStepResult(sig, container.Name, testkube.TestWorkflowStepResult{
						StartedAt: v.Value.CreationTimestamp.Time.UTC(),
						Status:    common.Ptr(testkube.RUNNING_TestWorkflowStepStatus),
					})
				}
				if v.Value.Type == "Normal" {
					continue
				}
				w.SendValue(Notification{
					Timestamp: v.Value.CreationTimestamp.Time,
					Log:       fmt.Sprintf("%s (%s) %s\n", v.Value.CreationTimestamp.Time.Format(time.RFC3339Nano), v.Value.Reason, v.Value.Message),
				})
			}

			// Emit the next result
			if stepResult.StartedAt.IsZero() {
				w.SendError(errors.New("step container is in unknown state"))
				break
			}
			w.SendValue(Notification{Result: result.Clone()})

			// Watch for the container logs, outputs and statuses
			// TODO: Calibrate clock with Hints
			for v := range WatchContainerLogs(ctx, c.clientSet, c.podEvents, c.namespace, pod.Name, container.Name).Stream(ctx).Channel() {
				if v.Error != nil {
					w.SendError(v.Error)
					continue
				}
				if v.Value.Hint != nil {
					if v.Value.Hint.Name == "start" && v.Value.Hint.Ref == container.Name {
						if v.Value.Time.After(stepResult.StartedAt) {
							stepResult = result.UpdateStepResult(sig, container.Name, testkube.TestWorkflowStepResult{
								StartedAt: v.Value.Time.UTC(),
							})
						}
					} else if v.Value.Hint.Name == "status" {
						status := testkube.TestWorkflowStepStatus(v.Value.Hint.Value.(string))
						if status == "" {
							status = testkube.PASSED_TestWorkflowStepStatus
						}
						if _, ok := result.Steps[v.Value.Hint.Ref]; ok {
							stepResult = result.UpdateStepResult(sig, v.Value.Hint.Ref, testkube.TestWorkflowStepResult{
								Status: &status,
							})
						}
					}
					continue
				}
				if v.Value.Output != nil {
					if _, ok := result.Steps[v.Value.Output.Ref]; ok {
						w.SendValue(Notification{
							Timestamp: v.Value.Time,
							Ref:       v.Value.Output.Ref,
							Output:    v.Value.Output,
						})
					}
					continue
				}
				w.SendValue(Notification{Timestamp: v.Value.Time, Log: string(v.Value.Log)})
			}

			// Watch container status
			status, err := GetFinalContainerResult(ctx, c.pod, container.Name)
			if err != nil {
				w.SendError(err)
				break
			}
			stepResult = result.UpdateStepResult(sig, container.Name, testkube.TestWorkflowStepResult{
				FinishedAt: status.FinishedAt.UTC(),
				ExitCode:   float64(status.ExitCode),
				Status:     common.Ptr(status.Status),
			})
			w.SendValue(Notification{Result: result.Clone()})

			// Update the last timestamp
			lastTs = status.FinishedAt
		}

		// Read the pod finish time
		for v := range c.job.Stream(ctx).Channel() {
			if v.Value != nil && v.Value.Status.CompletionTime != nil {
				result.FinishedAt = v.Value.Status.CompletionTime.Time
			}
		}

		// Compute the TestWorkflow status and dates
		result.Recompute(sig)
		w.SendValue(Notification{Result: result.Clone()})
	}()

	return w
}
