/*
 * TestKube API
 *
 * TestKube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// test execution summary
type TestStepExecutionSummary struct {
	Id string `json:"id"`
	// execution name
	Name string `json:"name"`
	// script name
	ScriptName string           `json:"scriptName,omitempty"`
	Status     *ExecutionStatus `json:"status"`
	Type_      *TestStepType    `json:"type,omitempty"`
}
