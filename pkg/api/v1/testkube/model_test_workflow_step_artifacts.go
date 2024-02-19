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

type TestWorkflowStepArtifacts struct {
	Compress *TestWorkflowStepArtifactsCompression `json:"compress,omitempty"`
	// file paths to fetch from the container
	Paths []string `json:"paths"`
}
