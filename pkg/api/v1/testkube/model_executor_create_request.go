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

// executor create request body
type ExecutorCreateRequest struct {
	// script name - Custom Resource name - must be unique, use only lowercase numbers and dashes (-)
	Name string `json:"name,omitempty"`
	// ExecutorType one of \"rest\" for rest openapi based executors or \"job\" which will be default runners for testkube soon
	ExecutorType string `json:"executor_type,omitempty"`
	// Image for kube-job
	Image string `json:"image,omitempty"`
	// Types defines what types can be handled by executor e.g. \"postman/collection\", \":curl/command\" etc
	Types []string `json:"types,omitempty"`
	// URI for rest based executors
	Uri string `json:"uri,omitempty"`
	// VolumeMountPath - where should PV be monted inside job pod for e.g. artifacts
	VolumeMountPath string `json:"volume_mount_path,omitempty"`
	// VolumeQuantity for kube-job PersistentVolume
	VolumeQuantity string `json:"volume_quantity,omitempty"`
}
