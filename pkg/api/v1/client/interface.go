package client

import (
	"io"
	"net/http"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Do(req *http.Request) (resp *http.Response, err error)
}

type Client interface {
	GetScript(id string) (script testkube.Script, err error)
	GetExecution(scriptID, executionID string) (execution testkube.Execution, err error)
	ListExecutions(scriptID string) (executions testkube.ExecutionsResult, err error)
	AbortExecution(script string, id string) error
	CreateScript(options CreateScriptOptions) (script testkube.Script, err error)
	DeleteScript(name string, namespace string) error
	DeleteScripts(namespace string) error
	ExecuteScript(id, namespace, executionName string, executionParams map[string]string) (execution testkube.Execution, err error)
	ListScripts(namespace string) (scripts testkube.Scripts, err error)
	GetServerInfo() (scripts testkube.ServerInfo, err error)

	CreateExecutor(options CreateExecutorOptions) (executor testkube.ExecutorDetails, err error)
	GetExecutor(name string) (executor testkube.ExecutorDetails, err error)
	ListExecutors() (executors testkube.ExecutorsDetails, err error)
	DeleteExecutor(name string) (err error)
}

// CreateScriptOptions - is mapping for now to OpenAPI schema for creating request
// if needed can beextended to custom struct
type CreateScriptOptions testkube.ScriptCreateRequest

// CreateExectorOptions - is mapping for now to OpenAPI schema for creating request
type CreateExecutorOptions testkube.ExecutorCreateRequest
