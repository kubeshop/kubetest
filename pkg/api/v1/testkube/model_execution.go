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

// test execution
type Execution struct {
	// execution id
	Id string `json:"id,omitempty"`
	// unique test name (CRD Test name)
	TestName string `json:"testName,omitempty"`
	// unique test suite name (CRD Test suite name), if it's run as a part of test suite
	TestSuiteName string `json:"testSuiteName,omitempty"`
	// test namespace
	TestNamespace string `json:"testNamespace,omitempty"`
	// test type e.g. postman/collection
	TestType string `json:"testType,omitempty"`
	// execution name
	Name string `json:"name,omitempty"`
	// execution number
	Number int32 `json:"number,omitempty"`
	// environment variables passed to executor
	Envs map[string]string `json:"envs,omitempty"`
	// additional arguments/flags passed to executor binary
	Args      []string            `json:"args,omitempty"`
	Variables map[string]Variable `json:"variables,omitempty"`
	// variables file content - need to be in format for particular executor (e.g. postman envs file)
	VariablesFile string `json:"variablesFile,omitempty"`
	// test secret uuid
	TestSecretUUID string `json:"testSecretUUID,omitempty"`
	// test suite secret uuid, if it's run as a part of test suite
	TestSuiteSecretUUID string       `json:"testSuiteSecretUUID,omitempty"`
	Content             *TestContent `json:"content,omitempty"`
	// test start time
	StartTime time.Time `json:"startTime,omitempty"`
	// test end time
	EndTime time.Time `json:"endTime,omitempty"`
	// test duration
	Duration string `json:"duration,omitempty"`
	// test duration in milliseconds
	DurationMs      int32            `json:"durationMs,omitempty"`
	ExecutionResult *ExecutionResult `json:"executionResult,omitempty"`
	// test and execution labels
	Labels map[string]string `json:"labels,omitempty"`
	// list of file paths that need to be copied into the test from uploads
	Uploads []string `json:"uploads,omitempty"`
}
