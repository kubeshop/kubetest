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

// TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace
type TypedLocalObjectReference struct {
	ApiGroup *BoxedString `json:"apiGroup,omitempty"`
	// kind is the type of resource being referenced
	Kind string `json:"kind,omitempty"`
	// name is the name of resource being referenced
	Name string `json:"name,omitempty"`
}
