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

type TestWorkflowResultSummary struct {
	Status          *TestWorkflowStatus `json:"status"`
	PredictedStatus *TestWorkflowStatus `json:"predictedStatus"`
	// when the pod was created
	QueuedAt time.Time `json:"queuedAt,omitempty"`
	// when the pod has been successfully assigned
	StartedAt time.Time `json:"startedAt,omitempty"`
	// when the pod has been completed
	FinishedAt time.Time `json:"finishedAt,omitempty"`
	// Go-formatted (human-readable) duration
	Duration string `json:"duration,omitempty"`
}
