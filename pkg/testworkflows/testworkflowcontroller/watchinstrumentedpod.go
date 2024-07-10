package testworkflowcontroller

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kubeshop/testkube/cmd/testworkflow-init/constants"
	"github.com/kubeshop/testkube/internal/common"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor"
)

const (
	InitContainerName = "tktw-init"
	IdleTimeout       = 100 * time.Millisecond
)

type WatchInstrumentedPodOptions struct {
	JobEvents Channel[*corev1.Event]
	Job       Channel[*batchv1.Job]
	Follow    *bool
}

func WatchInstrumentedPod(parentCtx context.Context, clientSet kubernetes.Interface, signature []testworkflowprocessor.Signature, scheduledAt time.Time, pod Channel[*corev1.Pod], podEvents Channel[*corev1.Event], opts WatchInstrumentedPodOptions) (<-chan ChannelMessage[Notification], error) {
	// Avoid missing data
	if pod == nil {
		return nil, errors.New("pod watcher is required")
	}

	// Initialize controller state
	ctx, ctxCancel := context.WithCancel(parentCtx)
	s := newNotifier(ctx, signature, scheduledAt)

	// Initialize pod state
	state := initializePodState(ctx, pod, podEvents, opts.Job, opts.JobEvents, s.Error)

	// Start watching
	go func() {
		defer func() {
			s.Flush()
			ctxCancel()
		}()

		// Watch for the basic initialization warnings
		for v := range state.PreStart("") {
			if v.Value.Queued != nil {
				s.Queue("", state.QueuedAt(""))
			} else if v.Value.Started != nil {
				s.Queue("", state.QueuedAt(""))
				s.Start("", state.StartedAt(""))
			} else if v.Value.Event != nil {
				ts := maxTime(v.Value.Event.CreationTimestamp.Time, v.Value.Event.FirstTimestamp.Time, v.Value.Event.LastTimestamp.Time)
				s.Event("", ts, v.Value.Event.Type, v.Value.Event.Reason, v.Value.Event.Message)
			}
		}

		// Ensure the queue/start time has been saved
		if s.result.QueuedAt.IsZero() || s.result.StartedAt.IsZero() {
			s.Error(errors.New("missing information about scheduled pod"))
			return
		}

		// Load the namespace information
		podObj := <-pod.Peek(ctx)

		// For each container:
		lastTs := s.result.Initialization.FinishedAt
		for _, container := range append(podObj.Spec.InitContainers, podObj.Spec.Containers...) {
			// Ignore non-standard TestWorkflow containers
			ref := container.Name
			if _, ok := s.result.Steps[ref]; !(ok || ref == InitContainerName) {
				continue
			}

			// Update queue time
			s.Queue(ref, lastTs)

			// Watch the container events
			for v := range state.PreStart(ref) {
				if v.Value.Queued != nil {
					s.Queue(ref, state.QueuedAt(ref))
				} else if v.Value.Started != nil {
					s.Queue(ref, state.QueuedAt(ref))
					s.Start(ref, state.StartedAt(ref))
				} else if v.Value.Event != nil {
					ts := maxTime(v.Value.Event.CreationTimestamp.Time, v.Value.Event.FirstTimestamp.Time, v.Value.Event.LastTimestamp.Time)
					s.Event(ref, ts, v.Value.Event.Type, v.Value.Event.Reason, v.Value.Event.Message)
				}
			}

			// Ensure the queue/start time has been saved
			if s.GetStepResult(ref).QueuedAt.IsZero() || s.GetStepResult(ref).StartedAt.IsZero() {
				s.Error(fmt.Errorf("missing information about scheduled '%s' container", ref))
				return
			}

			// Watch the container logs
			follow := common.ResolvePtr(opts.Follow, true) && !state.IsFinished(ref)
			for v := range WatchContainerLogs(ctx, clientSet, podObj.Namespace, podObj.Name, ref, 10, pod).Channel() {
				if v.Error != nil {
					s.Error(v.Error)
				} else if v.Value.Output != nil {
					s.Output(v.Value.Output.Ref, v.Value.Time, v.Value.Output)
				} else if v.Value.Hint != nil {
					switch v.Value.Hint.Name {
					case constants.InstructionStart:
						s.Start(ref, v.Value.Time)
					case constants.InstructionStatus:
						status := testkube.TestWorkflowStepStatus(v.Value.Hint.Value.(string))
						if status == "" {
							status = testkube.PASSED_TestWorkflowStepStatus
						}
						s.UpdateStepStatus(ref, status)
					case constants.InstructionPause:
						ts, _ := v.Value.Hint.Value.(string)
						start, err := time.Parse(constants.PreciseTimeFormat, ts)
						if err != nil {
							start = v.Value.Time
							s.Error(fmt.Errorf("invalid timestamp provided with pausing instruction: %v", v.Value.Hint.Value))
						}
						s.Pause(ref, start)
					case constants.InstructionResume:
						ts, _ := v.Value.Hint.Value.(string)
						end, err := time.Parse(constants.PreciseTimeFormat, ts)
						if err != nil {
							end = v.Value.Time
							s.Error(fmt.Errorf("invalid timestamp provided with resuming instruction: %v", v.Value.Hint.Value))
						}
						s.Resume(ref, end)
					}
				} else {
					s.Raw(ref, v.Value.Time, string(v.Value.Log), false)
				}
			}

			// Get the final result
			if follow {
				<-state.Finished(ref)
			} else {
				select {
				case <-state.Finished(ref):
				case <-time.After(IdleTimeout):
					return
				}
			}
			status, err := state.ContainerResult(ref)
			if err != nil {
				s.Error(err)
				break
			}
			s.FinishStep(ref, status)

			// Update the last timestamp
			lastTs = s.GetLastTimestamp(ref)

			// Break the function if the step has been aborted.
			// Breaking only to the loop is not enough,
			// because due to GKE bug, the Job is still pending,
			// so it will get stuck there.
			if status.Status == testkube.ABORTED_TestWorkflowStepStatus {
				if status.Details == "" {
					status.Details = "Manual"
				}
				s.Raw(ref, s.GetLastTimestamp(ref), fmt.Sprintf("\n%s Aborted (%s)", s.GetLastTimestamp(ref).Format(KubernetesLogTimeFormat), status.Details), false)
				break
			}
		}

		// Watch the completion time
		if s.result.FinishedAt.IsZero() {
			<-state.Finished("")
			f := state.FinishedAt("")
			s.Finish(f)
		}
	}()

	return s.ch, nil
}

func maxTime(times ...time.Time) time.Time {
	var result time.Time
	for _, t := range times {
		if t.After(result) {
			result = t
		}
	}
	return result
}
