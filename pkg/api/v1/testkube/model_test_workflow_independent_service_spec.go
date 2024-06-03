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

type TestWorkflowIndependentServiceSpec struct {
	WorkingDir *BoxedString `json:"workingDir,omitempty"`
	// image to be used for the container
	Image           string           `json:"image,omitempty"`
	ImagePullPolicy *ImagePullPolicy `json:"imagePullPolicy,omitempty"`
	// environment variables to append to the container
	Env []EnvVar `json:"env,omitempty"`
	// external environment variables to append to the container
	EnvFrom         []EnvFromSource        `json:"envFrom,omitempty"`
	Command         *BoxedStringList       `json:"command,omitempty"`
	Args            *BoxedStringList       `json:"args,omitempty"`
	Shell           *BoxedString           `json:"shell,omitempty"`
	Resources       *TestWorkflowResources `json:"resources,omitempty"`
	SecurityContext *SecurityContext       `json:"securityContext,omitempty"`
	// volumes to mount to the container
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
	// service description to display
	Description string `json:"description,omitempty"`
	// maximum time until reaching readiness
	Timeout string `json:"timeout,omitempty"`
	// list of files to send to parallel steps
	Transfer       []TestWorkflowStepParallelTransfer `json:"transfer,omitempty"`
	Content        *TestWorkflowContent               `json:"content,omitempty"`
	Pod            *TestWorkflowPodConfig             `json:"pod,omitempty"`
	Logs           *BoxedString                       `json:"logs,omitempty"`
	RestartPolicy  string                             `json:"restartPolicy,omitempty"`
	ReadinessProbe *Probe                             `json:"readinessProbe,omitempty"`
	Count          *BoxedString                       `json:"count,omitempty"`
	MaxCount       *BoxedString                       `json:"maxCount,omitempty"`
	// matrix of parameters to spawn instances
	Matrix map[string]interface{} `json:"matrix,omitempty"`
	// parameters that should be distributed across sharded instances
	Shards map[string]interface{} `json:"shards,omitempty"`
}
