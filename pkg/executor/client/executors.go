package client

import (
	"fmt"
	"sync"

	v1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
	executorscr "github.com/kubeshop/testkube-operator/client/executors"
)

func NewExecutors(client *executorscr.ExecutorsClient) Executors {
	return Executors{
		ExecutorsCRClient: client,
	}
}

// Executors represents available HTTP clients for executors registered in Kubernetes API
type Executors struct {
	ExecutorsCRClient *executorscr.ExecutorsClient
	Namespace         string
	Clients           sync.Map
}

func (p *Executors) GetExecutorSpec(scriptType string) (spec v1.ExecutorSpec, err error) {
	executorCR, err := p.ExecutorsCRClient.GetByType(scriptType)
	if err != nil {
		return spec, err
	}
	spec = executorCR.Spec
	return
}

// Get gets executor based on type with a basic map.Sync cache
// TODO there is no handling of CR change
func (p *Executors) Get(scriptType string) (client ExecutorClient, err error) {

	cached, exists := p.Clients.Load(scriptType)

	if !exists {
		// get executor from kubernetes CRs
		executorCR, err := p.ExecutorsCRClient.GetByType(scriptType)
		if err != nil {
			return client, fmt.Errorf("can't get executor spec: %w", err)
		}

		// get executor based on type
		executor, err := p.GetByType(scriptType, executorCR.Spec)
		if err != nil {
			return client, err
		}

		p.Clients.Store(scriptType, executor)
		cached = executor
	}

	client = cached.(ExecutorClient)
	return
}

func (e *Executors) GetByType(scriptType string, spec v1.ExecutorSpec) (executor ExecutorClient, err error) {
	// get executor based on type
	switch spec.ExecutorType {
	case ExecutorTypeRest:
		executor, err = e.GetOpenAPIExecutor(spec.URI)
	case ExecutorTypeJob:
		executor, err = e.GetJobExecutor()
	default:
		err = fmt.Errorf("can't handle runner type '%s' for script type '%s'", spec.ExecutorType, scriptType)
	}

	return
}

func (p *Executors) GetOpenAPIExecutor(uri string) (executor RestExecutorClient, err error) {
	return NewRestExecutorClient(RestExecutorConfig{
		URI: uri,
	}), nil

}

func (p *Executors) GetJobExecutor() (executor ExecutorClient, err error) {
	return NewJobExecutorClient()
}
