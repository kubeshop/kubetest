package cdevents

import (
	"fmt"
	"strings"

	cdevents "github.com/cdevents/sdk-go/pkg/api"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

// MapTestkubeEventToCDEvent maps OpenAPI spec Event to CDEvent CDEventReader
func MapTestkubeEventToCDEvent(tkEvent testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	switch tkEvent.Type_ {
	case testkube.EventStartTest:
		return MapTestkubeEventStartTestToCDEvent(tkEvent, clusterID, appVersion)
	case testkube.EventEndTestAborted, testkube.EventEndTestFailed, testkube.EventEndTestTimeout, testkube.EventEndTestSuccess:
		return MapTestkubeEventFinishTestToCDEvent(tkEvent, clusterID, appVersion)
	case testkube.EventStartTestSuite:
		return MapTestkubeEventStartTestSuiteToCDEvent(tkEvent, clusterID, appVersion)
	case testkube.EventEndTestSuiteAborted, testkube.EventEndTestSuiteFailed, testkube.EventEndTestSuiteTimeout, testkube.EventEndTestSuiteSuccess:
		return MapTestkubeEventFinishTestSuiteToCDEvent(tkEvent, clusterID, appVersion)
	}

	return nil, fmt.Errorf("not supported event type %s", tkEvent.Type_)
}

// MapTestkubeEventQueuedTestToCDEvent maps OpenAPI spec Queued Test Event to CDEvent CDEventReader
func MapTestkubeEventQueuedTestToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestCaseRunQueuedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestExecution != nil {
		ev.SetSubjectTestCase(&cdevents.TestCaseRunQueuedSubjectContentTestCase{
			Id:      event.TestExecution.Id,
			Name:    event.TestExecution.TestName,
			Type:    event.TestExecution.TestType,
			Version: appVersion,
		})

		ev.SetSubjectEnvironment(&cdevents.Reference{
			Id:     event.TestExecution.TestNamespace,
			Source: clusterID,
		})

		if event.TestExecution.RunningContext != nil {
			ev.SetSubjectTrigger(&cdevents.TestCaseRunQueuedSubjectContentTrigger{
				Type: event.TestExecution.RunningContext.Type_,
			})
		}
	}

	if event.TestSuiteExecution != nil {
		ev.SetSubjectTestSuiteRun(&cdevents.Reference{
			Id:     event.TestSuiteExecution.Id,
			Source: clusterID,
		})
	}

	return ev, nil
}

// MapTestkubeEventStartTestToCDEvent maps OpenAPI spec Start Test Event to CDEvent CDEventReader
func MapTestkubeEventStartTestToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestCaseRunStartedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestExecution != nil {
		ev.SetSubjectTestCase(&cdevents.TestCaseRunStartedSubjectContentTestCase{
			Id:      event.TestExecution.Id,
			Name:    event.TestExecution.TestName,
			Type:    event.TestExecution.TestType,
			Version: appVersion,
		})

		ev.SetSubjectEnvironment(&cdevents.Reference{
			Id:     event.TestExecution.TestNamespace,
			Source: clusterID,
		})

		if event.TestExecution.RunningContext != nil {
			ev.SetSubjectTrigger(&cdevents.TestCaseRunStartedSubjectContentTrigger{
				Type: event.TestExecution.RunningContext.Type_,
			})
		}
	}

	if event.TestSuiteExecution != nil {
		ev.SetSubjectTestSuiteRun(&cdevents.Reference{
			Id:     event.TestSuiteExecution.Id,
			Source: clusterID,
		})
	}

	return ev, nil
}

// MapTestkubeEventFinishTestToCDEvent maps OpenAPI spec Failed, Aborted, Timeout, Success Test Event to CDEvent CDEventReader
func MapTestkubeEventFinishTestToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestCaseRunFinishedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestExecution != nil {
		ev.SetSubjectTestCase(&cdevents.TestCaseRunFinishedSubjectContentTestCase{
			Id:      event.TestExecution.Id,
			Name:    event.TestExecution.TestName,
			Type:    event.TestExecution.TestType,
			Version: appVersion,
		})

		ev.SetSubjectEnvironment(&cdevents.Reference{
			Id:     event.TestExecution.TestNamespace,
			Source: clusterID,
		})

		if event.TestExecution.IsAborted() || event.TestExecution.IsTimeout() {
			ev.SetSubjectOutcome("cancel")
			if event.TestExecution.ExecutionResult != nil {
				ev.SetSubjectReason(event.TestExecution.ExecutionResult.ErrorMessage)
			}
		}

		if event.TestExecution.IsFailed() {
			ev.SetSubjectOutcome("fail")
			if event.TestExecution.ExecutionResult != nil {
				ev.SetSubjectReason(event.TestExecution.ExecutionResult.ErrorMessage)
			}
		}

		if event.TestExecution.IsPassed() {
			ev.SetSubjectOutcome("pass")
		}
	}

	if event.TestSuiteExecution != nil {
		ev.SetSubjectTestSuiteRun(&cdevents.Reference{
			Id:     event.TestSuiteExecution.Id,
			Source: clusterID,
		})
	}

	return ev, nil
}

// MapTestkubeArtifactToCDEvent maps OpenAPI spec Artifact to CDEvent CDEventReader
func MapTestkubeArtifactToCDEvent(execution *testkube.Execution, clusterID, appVersion, format, outputType string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestOutputPublishedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(execution.Id)
	ev.SetSubjectSource(clusterID)
	ev.SetSubjectTestCaseRun(&cdevents.Reference{
		Id:     execution.Id,
		Source: clusterID,
	})

	ev.SetSubjectFormat(format)
	ev.SetSubjectOutputType(outputType)

	return ev, nil
}

// MapTestkubeEventQueuedTestSuiteToCDEvent maps OpenAPI spec Queued Test Suite Event to CDEvent CDEventReader
func MapTestkubeEventQueuedTestSuiteToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestSuiteRunQueuedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestSuiteExecution != nil {
		if event.TestSuiteExecution.TestSuite != nil {
			ev.SetSubjectTestSuite(&cdevents.TestSuiteRunQueuedSubjectContentTestSuite{
				Id:      event.TestSuiteExecution.Id,
				Name:    event.TestSuiteExecution.TestSuite.Name,
				Version: appVersion,
			})

			ev.SetSubjectEnvironment(&cdevents.Reference{
				Id:     event.TestSuiteExecution.TestSuite.Namespace,
				Source: clusterID,
			})
		}

		if event.TestSuiteExecution.RunningContext != nil {
			ev.SetSubjectTrigger(&cdevents.TestSuiteRunQueuedSubjectContentTrigger{
				Type: event.TestSuiteExecution.RunningContext.Type_,
			})
		}
	}

	return ev, nil
}

// MapTestkubeEventStartTestSuiteToCDEvent maps OpenAPI spec Start Test Suite Event to CDEvent CDEventReader
func MapTestkubeEventStartTestSuiteToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestSuiteRunStartedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestSuiteExecution != nil {
		if event.TestSuiteExecution.TestSuite != nil {
			ev.SetSubjectTestSuite(&cdevents.TestSuiteRunStartedSubjectContentTestSuite{
				Id:      event.TestSuiteExecution.Id,
				Name:    event.TestSuiteExecution.TestSuite.Name,
				Version: appVersion,
			})

			ev.SetSubjectEnvironment(&cdevents.Reference{
				Id:     event.TestSuiteExecution.TestSuite.Namespace,
				Source: clusterID,
			})
		}

		if event.TestSuiteExecution.RunningContext != nil {
			ev.SetSubjectTrigger(&cdevents.TestSuiteRunStartedSubjectContentTrigger{
				Type: event.TestSuiteExecution.RunningContext.Type_,
			})
		}
	}

	return ev, nil
}

// MapTestkubeEventFinishTestSuiteToCDEvent maps OpenAPI spec Failed, Aborted, Timeout, Success Test Event to CDEvent CDEventReader
func MapTestkubeEventFinishTestSuiteToCDEvent(event testkube.Event, clusterID, appVersion string) (cdevents.CDEventReader, error) {
	// Create the base event
	ev, err := cdevents.NewTestSuiteRunFinishedEvent()
	if err != nil {
		return nil, err
	}

	ev.SetSubjectId(event.Id)
	ev.SetSubjectSource(clusterID)
	if event.TestSuiteExecution != nil {
		if event.TestSuiteExecution.TestSuite != nil {
			ev.SetSubjectTestSuite(&cdevents.TestSuiteRunFinishedSubjectContentTestSuite{
				Id:      event.TestSuiteExecution.Id,
				Name:    event.TestSuiteExecution.TestSuite.Name,
				Version: appVersion,
			})

			ev.SetSubjectEnvironment(&cdevents.Reference{
				Id:     event.TestSuiteExecution.TestSuite.Namespace,
				Source: clusterID,
			})
		}

		if event.TestSuiteExecution.IsAborted() || event.TestSuiteExecution.IsTimeout() {
			ev.SetSubjectOutcome("cancel")
			var errs []string
			for _, result := range event.TestSuiteExecution.StepResults {
				if result.Execution != nil && result.Execution.ExecutionResult != nil {
					errs = append(errs, result.Execution.ExecutionResult.ErrorMessage)
				}
			}

			ev.SetSubjectReason(strings.Join(errs, ","))
		}

		if event.TestSuiteExecution.IsFailed() {
			ev.SetSubjectOutcome("fail")
			var errs []string
			for _, result := range event.TestSuiteExecution.StepResults {
				if result.Execution != nil && result.Execution.ExecutionResult != nil {
					errs = append(errs, result.Execution.ExecutionResult.ErrorMessage)
				}
			}

			ev.SetSubjectReason(strings.Join(errs, ","))
		}

		if event.TestSuiteExecution.IsPassed() {
			ev.SetSubjectOutcome("pass")
		}
	}

	return ev, nil
}
