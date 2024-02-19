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

// Event data
type Event struct {
	// UUID of event
	Id       string         `json:"id"`
	Resource *EventResource `json:"resource"`
	// custom stream topic for event
	StreamTopic string `json:"streamTopic,omitempty"`
	// ID of resource
	ResourceId         string              `json:"resourceId"`
	Type_              *EventType          `json:"type"`
	TestExecution      *Execution          `json:"testExecution,omitempty"`
	TestSuiteExecution *TestSuiteExecution `json:"testSuiteExecution,omitempty"`
	// cluster name of event
	ClusterName string `json:"clusterName,omitempty"`
	// environment variables
	Envs map[string]string `json:"envs,omitempty"`
}
