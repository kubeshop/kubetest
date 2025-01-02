package runner

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/cloud"
	"github.com/kubeshop/testkube/pkg/repository/testworkflow"
	"github.com/kubeshop/testkube/pkg/testworkflows/executionworker/controller/store"
)

const (
	ExecutionSaverUpdateRetryCount = 10
	ExecutionSaverUpdateRetryDelay = 300 * time.Millisecond
)

//go:generate mockgen -destination=./mock_executionsaver.go -package=runner "github.com/kubeshop/testkube/pkg/runner" ExecutionSaver
type ExecutionSaver interface {
	UpdateResult(result testkube.TestWorkflowResult)
	AppendOutput(output ...testkube.TestWorkflowOutput)
	End(ctx context.Context, result testkube.TestWorkflowResult) error
}

type executionSaver struct {
	id                   string
	environmentId        string
	executionsRepository testworkflow.Repository
	client               cloud.TestKubeCloudAPIClient
	logs                 ExecutionLogsWriter
	newExecutionsEnabled bool

	// Intermediate data
	output       []testkube.TestWorkflowOutput
	result       *testkube.TestWorkflowResult
	resultUpdate store.Update
	resultMu     sync.Mutex

	outputSaved *atomic.Bool

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func NewExecutionSaver(
	ctx context.Context,
	executionsRepository testworkflow.Repository,
	grpcConn *grpc.ClientConn,
	id string,
	environmentId string,
	logs ExecutionLogsWriter,
	newExecutionsEnabled bool,
) (ExecutionSaver, error) {
	ctx, cancel := context.WithCancel(ctx)
	outputSaved := atomic.Bool{}
	outputSaved.Store(true)
	saver := &executionSaver{
		id:                   id,
		environmentId:        environmentId,
		executionsRepository: executionsRepository,
		client:               cloud.NewTestKubeCloudAPIClient(grpcConn),
		logs:                 logs,
		newExecutionsEnabled: newExecutionsEnabled,
		resultUpdate:         store.NewUpdate(),
		outputSaved:          &outputSaved,
		ctx:                  ctx,
		ctxCancel:            cancel,
	}
	go saver.watchResultUpdates()

	return saver, nil
}

func (s *executionSaver) watchResultUpdates() {
	defer s.resultUpdate.Close()
	ch := s.resultUpdate.Channel(s.ctx)
	var prev *testkube.TestWorkflowResult
	for {
		select {
		case <-s.ctx.Done():
			return
		case _, ok := <-ch:
			if !ok {
				return
			}
			for i := 0; i < ExecutionSaverUpdateRetryCount; i++ {
				s.resultMu.Lock()
				next := s.result
				s.resultMu.Unlock()
				if prev == next {
					break
				}
				err := s.executionsRepository.UpdateResult(s.ctx, s.id, next)
				if err == nil {
					break
				}
				select {
				case <-s.ctx.Done():
					return
				case <-time.After(ExecutionSaverUpdateRetryDelay):
				}
			}
		}
	}
}

func (s *executionSaver) UpdateResult(result testkube.TestWorkflowResult) {
	s.resultMu.Lock()
	defer s.resultMu.Unlock()
	s.result = &result
	s.resultUpdate.Emit()
}

func (s *executionSaver) AppendOutput(output ...testkube.TestWorkflowOutput) {
	s.output = append(s.output, output...)
	s.outputSaved.Store(false)
}

func (s *executionSaver) End(ctx context.Context, result testkube.TestWorkflowResult) error {
	s.ctxCancel()
	s.resultMu.Lock()
	defer s.resultMu.Unlock()

	// Save the logs and output
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		if s.outputSaved.Load() {
			return nil
		}
		// TODO: Consider AppendOutput ($push) instead
		return s.executionsRepository.UpdateOutput(ctx, s.id, s.output)
	})
	g.Go(func() error {
		if s.logs.Saved() {
			return nil
		}
		return s.logs.Save(ctx)
	})
	err := g.Wait()
	if err != nil {
		return err
	}

	// Save the final result
	if s.newExecutionsEnabled {
		err = s.saveFinalResult(ctx, &result)
	} else {
		err = s.executionsRepository.UpdateResult(ctx, s.id, &result)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *executionSaver) saveFinalResult(ctx context.Context, result *testkube.TestWorkflowResult) error {
	if result == nil {
		return errors.New("missing result")
	}
	resultBytes, err := json.Marshal(result)
	_, err = s.client.FinishExecution(ctx, &cloud.FinishExecutionRequest{
		EnvironmentId: s.environmentId,
		Id:            s.id,
		Result:        resultBytes,
	})
	return err
}
