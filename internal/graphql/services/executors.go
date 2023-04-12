package services

import (
	"context"

	executorsclientv1 "github.com/kubeshop/testkube-operator/client/executors/v1"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	executorsmapper "github.com/kubeshop/testkube/pkg/mapper/executors"
)

type ExecutorsService struct {
	*Service
	Client *executorsclientv1.ExecutorsClient
}

func (s *ExecutorsService) List() ([]testkube.ExecutorDetails, error) {
	execs, err := s.Client.List("")
	if err != nil {
		return nil, err
	}
	return Map(execs.Items, executorsmapper.MapExecutorCRDToExecutorDetails), nil
}

func (s *ExecutorsService) SubscribeList(ctx context.Context) (<-chan []testkube.ExecutorDetails, error) {
	return HandleSubscription(ctx, "events.executor.>", s, s.List)
}
