/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

type EventType string

// List of EventType
const (
	START_TEST_EventType               EventType = "start-test"
	END_TEST_SUCCESS_EventType         EventType = "end-test-success"
	END_TEST_FAILED_EventType          EventType = "end-test-failed"
	END_TEST_ABORTED_EventType         EventType = "end-test-aborted"
	END_TEST_TIMEOUT_EventType         EventType = "end-test-timeout"
	START_TESTSUITE_EventType          EventType = "start-testsuite"
	END_TESTSUITE_SUCCESS_EventType    EventType = "end-testsuite-success"
	END_TESTSUITE_FAILED_EventType     EventType = "end-testsuite-failed"
	END_TESTSUITE_ABORTED_EventType    EventType = "end-testsuite-aborted"
	END_TESTSUITE_TIMEOUT_EventType    EventType = "end-testsuite-timeout"
	QUEUE_TESTWORKFLOW_EventType       EventType = "queue-testworkflow"
	START_TESTWORKFLOW_EventType       EventType = "start-testworkflow"
	END_TESTWORKFLOW_SUCCESS_EventType EventType = "end-testworkflow-success"
	END_TESTWORKFLOW_FAILED_EventType  EventType = "end-testworkflow-failed"
	END_TESTWORKFLOW_ABORTED_EventType EventType = "end-testworkflow-aborted"
	CREATED_EventType                  EventType = "created"
	UPDATED_EventType                  EventType = "updated"
	DELETED_EventType                  EventType = "deleted"
)
