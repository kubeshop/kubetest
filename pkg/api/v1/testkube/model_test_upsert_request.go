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

// test create request body
type TestUpsertRequest struct {
	// test name
	Name string `json:"name,omitempty"`
	// test namespace
	Namespace string `json:"namespace,omitempty"`
	// test type
	Type_   string       `json:"type,omitempty"`
	Content *TestContent `json:"content,omitempty"`
	Created time.Time    `json:"created,omitempty"`
	// test tags
	Tags []string `json:"tags,omitempty"`
}
