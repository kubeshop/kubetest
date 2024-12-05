package testworkflowexecutor

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/pkg/client/testworkflows/v1"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/cloud"
	"github.com/kubeshop/testkube/pkg/log"
	"github.com/kubeshop/testkube/pkg/repository/testworkflow"
	"github.com/kubeshop/testkube/pkg/testworkflows"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowresolver"
)

const (
	SaveResultRetryMaxAttempts = 100
	SaveResultRetryBaseDelay   = 300 * time.Millisecond
)

type scheduler struct {
	logger                      *zap.SugaredLogger
	testWorkflowsClient         testworkflowsv1.Interface
	testWorkflowTemplatesClient testworkflowsv1.TestWorkflowTemplatesInterface
	resultsRepository           testworkflow.Repository
	outputRepository            testworkflow.OutputRepository
	globalTemplateName          string
	organizationId              string
}

func NewScheduler(
	testWorkflowsClient testworkflowsv1.Interface,
	testWorkflowTemplatesClient testworkflowsv1.TestWorkflowTemplatesInterface,
	resultsRepository testworkflow.Repository,
	outputRepository testworkflow.OutputRepository,
	globalTemplateName string,
	organizationId string,
) *scheduler {
	return &scheduler{
		logger:                      log.DefaultLogger,
		testWorkflowsClient:         testWorkflowsClient,
		testWorkflowTemplatesClient: testWorkflowTemplatesClient,
		resultsRepository:           resultsRepository,
		outputRepository:            outputRepository,
		globalTemplateName:          globalTemplateName,
		organizationId:              organizationId,
	}
}

func (s *scheduler) insert(ctx context.Context, execution *testkube.TestWorkflowExecution) error {
	err := retry(SaveResultRetryMaxAttempts, SaveResultRetryBaseDelay, func() error {
		err := s.resultsRepository.Insert(ctx, *execution)
		if err != nil {
			s.logger.Warnw("failed to update the TestWorkflow execution in database", "recoverable", true, "executionId", execution.Id, "error", err)
		}
		return err
	})
	if err != nil {
		s.logger.Errorw("failed to update the TestWorkflow execution in database", "recoverable", false, "executionId", execution.Id, "error", err)
	}
	return err
}

func (s *scheduler) update(ctx context.Context, execution *testkube.TestWorkflowExecution) error {
	err := retry(SaveResultRetryMaxAttempts, SaveResultRetryBaseDelay, func() error {
		err := s.resultsRepository.Update(ctx, *execution)
		if err != nil {
			s.logger.Warnw("failed to update the TestWorkflow execution in database", "recoverable", true, "executionId", execution.Id, "error", err)
		}
		return err
	})
	if err != nil {
		s.logger.Errorw("failed to update the TestWorkflow execution in database", "recoverable", false, "executionId", execution.Id, "error", err)
	}
	return err
}

func (s *scheduler) init(ctx context.Context, execution *testkube.TestWorkflowExecution) error {
	err := retry(SaveResultRetryMaxAttempts, SaveResultRetryBaseDelay, func() error {
		err := s.resultsRepository.Init(ctx, execution.Id, testworkflow.InitData{
			RunnerID:  execution.RunnerId,
			Namespace: execution.Namespace,
			Signature: execution.Signature,
		})
		if err != nil {
			s.logger.Warnw("failed to initialize the TestWorkflow execution in database", "recoverable", true, "executionId", execution.Id, "error", err)
		}
		return err
	})
	if err != nil {
		s.logger.Errorw("failed to initialize the TestWorkflow execution in database", "recoverable", false, "executionId", execution.Id, "error", err)
	}
	return err
}

func (s *scheduler) saveEmptyLogs(ctx context.Context, execution *testkube.TestWorkflowExecution) error {
	err := retry(SaveResultRetryMaxAttempts, SaveResultRetryBaseDelay, func() error {
		return s.outputRepository.SaveLog(ctx, execution.Id, execution.Workflow.Name, bytes.NewReader(nil))
	})
	if err != nil {
		s.logger.Errorw("failed to save empty log", "executionId", execution.Id, "error", err)
	}
	return err
}

func (s *scheduler) Schedule(ctx context.Context, sensitiveDataHandler SensitiveDataHandler, req *cloud.ScheduleRequest) (<-chan *testkube.TestWorkflowExecution, error) {
	// Prepare the channel
	ch := make(chan *testkube.TestWorkflowExecution, 1)

	// Set up context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Validate the execution request
	if err := ValidateExecutionRequest(req); err != nil {
		close(ch)
		return ch, err
	}

	// Check if there is anything to run
	if len(req.Selectors) == 0 {
		close(ch)
		return ch, nil
	}

	// Initialize execution template
	now := time.Now().UTC()
	base := NewIntermediateExecution().
		SetGroupID(primitive.NewObjectIDFromTimestamp(now).Hex()).
		SetScheduledAt(now).
		AppendTags(req.Tags).
		SetDisabledWebhooks(req.DisableWebhooks).
		SetKubernetesObjectName(req.KubernetesObjectName).
		SetRunningContext(GetLegacyRunningContext(req)).
		PrependTemplate(s.globalTemplateName)

	// Initialize fetchers
	testWorkflows := NewTestWorkflowFetcher(s.testWorkflowsClient)
	testWorkflowTemplates := NewTestWorkflowTemplateFetcher(s.testWorkflowTemplatesClient)

	// Prefetch all the Test Workflows
	err := testWorkflows.PrefetchMany(req.Selectors)
	if err != nil {
		close(ch)
		return ch, err
	}

	// Prefetch all the Test Workflow Templates.
	// Don't fail immediately - it should be execution's error message if it's missing.
	tplNames := testWorkflows.TemplateNames()
	if s.globalTemplateName != "" {
		tplNames[testworkflowresolver.GetInternalTemplateName(s.globalTemplateName)] = struct{}{}
	}
	_ = testWorkflowTemplates.PrefetchMany(tplNames)

	// Flatten selectors
	selectors := make([]*cloud.ScheduleSelector, 0, len(req.Selectors))
	for i := range req.Selectors {
		list, _ := testWorkflows.Get(req.Selectors[i])
		for _, w := range list {
			selectors = append(selectors, &cloud.ScheduleSelector{
				Name:          w.Name,
				Config:        req.Selectors[i].Config,
				ExecutionName: req.Selectors[i].ExecutionName, // TODO: what to do when execution name is configured, but multiple requested?
				Tags:          req.Selectors[i].Tags,
			})
		}
	}

	// Resolve executions for each selector
	intermediate := make([]*IntermediateExecution, 0, len(req.Selectors))
	for _, v := range selectors {
		workflow, _ := testWorkflows.GetByName(v.Name)
		current := base.Clone().
			AutoGenerateID().
			SetName(v.ExecutionName).
			AppendTags(v.Tags).
			SetWorkflow(workflow)
		intermediate = append(intermediate, current)

		// Inject configuration
		storeConfig := true
		schema := workflow.Spec.Config
		for k := range v.Config {
			if s, ok := schema[k]; ok && s.Sensitive {
				storeConfig = false
			}
		}
		if storeConfig && testworkflows.CountMapBytes(v.Config) < ConfigSizeLimit {
			current.StoreConfig(v.Config)
		}

		// Apply the configuration
		if err := current.ApplyConfig(v.Config); err != nil {
			current.SetError("Cannot inline Test Workflow configuration", err)
			continue
		}

		// Load the required Test Workflow Templates
		tpls, err := testWorkflowTemplates.GetMany(current.TemplateNames())
		if err != nil {
			current.SetError("Cannot fetch required Test Workflow Templates", err)
			continue
		}

		// Apply the Test Workflow Templates
		if err = current.ApplyTemplates(tpls); err != nil {
			current.SetError("Cannot inline Test Workflow Templates", err)
			continue
		}
	}

	// Simplify group ID in case of single execution
	if len(intermediate) == 1 {
		intermediate[0].SetGroupID(intermediate[0].ID())
	}

	// Validate if there are no execution name duplicates initially
	if err = ValidateExecutionNameDuplicates(intermediate); err != nil {
		close(ch)
		return ch, err
	}

	// Validate if the static execution names are not reserved in the database already
	for i := range intermediate {
		if intermediate[i].Name() == "" {
			continue
		}
		if err = ValidateExecutionNameRemoteDuplicate(ctx, s.resultsRepository, intermediate[i]); err != nil {
			close(ch)
			return ch, err
		}
	}

	// Ensure the rest of operations won't be stopped if started
	if ctx.Err() != nil {
		close(ch)
		return ch, ctx.Err()
	}
	cancel()
	ctx = context.Background()

	// Generate execution names and sequence numbers
	for i := range intermediate {
		// Load execution identifier data
		number, err := s.resultsRepository.GetNextExecutionNumber(context.Background(), intermediate[i].WorkflowName())
		if err != nil {
			close(ch)
			return ch, errors.Wrap(err, "registering next execution sequence number")
		}
		intermediate[i].SetSequenceNumber(number)

		// Generating the execution name
		if intermediate[i].Name() == "" {
			name := fmt.Sprintf("%s-%d", intermediate[i].WorkflowName(), number)
			intermediate[i].SetName(name)

			// Edge case: Check for local duplicates, if there is no clash between static and auto-generated one
			if err = ValidateExecutionNameDuplicates(intermediate); err != nil {
				return ch, err
			}

			// Ensure the execution name is unique
			if err = ValidateExecutionNameRemoteDuplicate(context.Background(), s.resultsRepository, intermediate[i]); err != nil {
				close(ch)
				return ch, err
			}
		}

		// Resolve it finally
		err = intermediate[i].Resolve(s.organizationId, req.EnvironmentId, req.ParentExecutionIds, false)
		if err != nil {
			intermediate[i].SetError("Cannot process Test Workflow specification", err)
			continue
		}
	}

	go func() {
		defer close(ch)
		for i := range intermediate {
			// Prepare sensitive data
			if err = sensitiveDataHandler.Process(intermediate[i]); err != nil {
				intermediate[i].SetError("Cannot store the sensitive data", err)
			}

			// Save empty logs if the execution is already finished
			if intermediate[i].Finished() {
				_ = s.saveEmptyLogs(context.Background(), intermediate[i].Execution())
			}

			// Insert the execution
			if err = s.insert(context.Background(), intermediate[i].Execution()); err != nil {
				sensitiveDataHandler.Rollback(intermediate[i].ID())
				// TODO: notify API about problem (?)
				continue
			}

			// Inform about the next execution
			ch <- intermediate[i].Execution()
		}
	}()

	return ch, nil
}

func (s *scheduler) CriticalError(execution *testkube.TestWorkflowExecution, name string, err error) error {
	execution.InitializationError(name, err)
	_ = s.saveEmptyLogs(context.Background(), execution)
	return s.update(context.Background(), execution)
}

func (s *scheduler) Start(execution *testkube.TestWorkflowExecution) error {
	return s.init(context.Background(), execution)
}
