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

// test execution start request body
type ExecutorStartRequest struct {
	// ID of test execution to handle by executor, execution need to be able to return execution info based on this ID
	Id string `json:"id,omitempty"`
	// test type
	Type_ string `json:"type,omitempty"`
	// test execution custom name
	Name string `json:"name,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
	// params file content - need to be in format for particular executor (e.g. postman envs file)
	ParamsFile string       `json:"paramsFile,omitempty"`
	Content    *TestContent `json:"content,omitempty"`
}
