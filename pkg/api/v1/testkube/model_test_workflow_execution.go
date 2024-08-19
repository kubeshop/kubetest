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

type TestWorkflowExecution struct {
	// unique execution identifier
	Id string `json:"id"`
	// execution name
	Name string `json:"name"`
	// execution namespace
	Namespace string `json:"namespace,omitempty"`
	// sequence number for the execution
	Number int32 `json:"number,omitempty"`
	// when the execution has been scheduled to run
	ScheduledAt time.Time `json:"scheduledAt,omitempty"`
	// when the execution result's status has changed last time (queued, passed, failed)
	StatusAt time.Time `json:"statusAt,omitempty"`
	// structured tree of steps
	Signature []TestWorkflowSignature `json:"signature,omitempty"`
	Result    *TestWorkflowResult     `json:"result,omitempty"`
	// additional information from the steps, like referenced executed tests or artifacts
	Output []TestWorkflowOutput `json:"output,omitempty"`
	// generated reports from the steps, like junit
	Reports          []TestWorkflowReport `json:"reports,omitempty"`
	Workflow         *TestWorkflow        `json:"workflow"`
	ResolvedWorkflow *TestWorkflow        `json:"resolvedWorkflow,omitempty"`
	// test workflow execution name started the test workflow execution
	TestWorkflowExecutionName string `json:"testWorkflowExecutionName,omitempty"`
	// whether webhooks on the execution of this test workflow are disabled
	DisableWebhooks bool              `json:"disableWebhooks,omitempty"`
	Tags            map[string]string `json:"tags,omitempty"`
}
