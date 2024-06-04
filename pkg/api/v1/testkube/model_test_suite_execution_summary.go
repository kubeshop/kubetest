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

import (
	"time"
)

// Test execution summary
type TestSuiteExecutionSummary struct {
	// execution id
	Id string `json:"id"`
	// execution name
	Name string `json:"name"`
	// name of the test suite
	TestSuiteName string                    `json:"testSuiteName"`
	Status        *TestSuiteExecutionStatus `json:"status"`
	// test suite execution start time
	StartTime time.Time `json:"startTime,omitempty"`
	// test suite execution end time
	EndTime time.Time `json:"endTime,omitempty"`
	// test suite execution duration
	Duration string `json:"duration,omitempty"`
	// test suite execution duration in ms
	DurationMs int32                                `json:"durationMs,omitempty"`
	Execution  []TestSuiteBatchStepExecutionSummary `json:"execution,omitempty"`
	// test suite and execution labels
	Labels map[string]string `json:"labels,omitempty"`
}