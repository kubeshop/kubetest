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

type ExecutionStatus string

// List of ExecutionStatus
const (
	QUEUED_ExecutionStatus  ExecutionStatus = "queued"
	RUNNING_ExecutionStatus ExecutionStatus = "running"
	PASSED_ExecutionStatus  ExecutionStatus = "passed"
	FAILED_ExecutionStatus  ExecutionStatus = "failed"
)
