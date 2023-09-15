package triggers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/event/bus"
)

func (s *Service) runExecutionScraper(ctx context.Context) {
	ticker := time.NewTicker(s.scraperInterval)
	s.logger.Debugf("trigger service: starting execution scraper")

	for {
		select {
		case <-ctx.Done():
			s.logger.Infof("trigger service: stopping scraper component")
			return
		case <-ticker.C:
			s.logger.Debugf("trigger service: execution scraper component: starting new ticker iteration")
			for triggerName, status := range s.triggerStatus {
				if status.hasActiveTests() {
					s.checkForRunningTestExecutions(ctx, status)
					s.checkForRunningTestSuiteExecutions(ctx, status)
					if !status.hasActiveTests() {
						s.logger.Debugf("marking status as finished for testtrigger %s", triggerName)
						status.done()
					}
				}
			}
		}
	}
}

func (s *Service) checkForRunningTestExecutions(ctx context.Context, status *triggerStatus) {
	testExecutionIDs := status.getExecutionIDs()

	for _, id := range testExecutionIDs {
		execution, err := s.resultRepository.Get(ctx, id)
		if err == mongo.ErrNoDocuments {
			s.logger.Warnf("trigger service: execution scraper component: no test execution found for id %s", id)
			status.removeExecutionID(id)
			continue
		} else if err != nil {
			s.logger.Errorf("trigger service: execution scraper component: error fetching test execution result: %v", err)
			continue
		}
		if !execution.IsRunning() && !execution.IsQueued() {
			s.logger.Debugf("trigger service: execution scraper component: test execution %s is finished", id)
			status.removeExecutionID(id)
		}
	}
}

func (s *Service) checkForRunningTestSuiteExecutions(ctx context.Context, status *triggerStatus) {
	testSuiteExecutionIDs := status.getTestSuiteExecutionIDs()

	for _, id := range testSuiteExecutionIDs {
		execution, err := s.testResultRepository.Get(ctx, id)
		if err == mongo.ErrNoDocuments {
			s.logger.Warnf("trigger service: execution scraper component: no testsuite execution found for id %s", id)
			status.removeTestSuiteExecutionID(id)
			continue
		} else if err != nil {
			s.logger.Errorf("trigger service: execution scraper component: error fetching testsuite execution result: %v", err)
			continue
		}
		if !execution.IsRunning() && !execution.IsQueued() {
			s.logger.Debugf("trigger service: execution scraper component: testsuite execution %s is finished", id)
			status.removeTestSuiteExecutionID(id)
		}
	}
}

func (s *Service) abortExecutions(ctx context.Context, testTriggerName string, status *triggerStatus) {
	s.logger.Debugf("trigger service: abort executions")
	s.abortRunningTestExecutions(ctx, status)
	s.abortRunningTestSuiteExecutions(ctx, status)
	if !status.hasActiveTests() {
		s.logger.Debugf("marking status as finished for testtrigger %s", testTriggerName)
		status.done()
	}
}

func (s *Service) abortRunningTestExecutions(ctx context.Context, status *triggerStatus) {
	testExecutionIDs := status.getExecutionIDs()

	for _, id := range testExecutionIDs {
		execution, err := s.resultRepository.Get(ctx, id)
		if err == mongo.ErrNoDocuments {
			s.logger.Warnf("trigger service: execution scraper component: no test execution found for id %s", id)
			status.removeExecutionID(id)
			continue
		} else if err != nil {
			s.logger.Errorf("trigger service: execution scraper component: error fetching test execution result: %v", err)
			continue
		}
		if execution.IsRunning() || execution.IsQueued() {
			res, err := s.testExecutor.Abort(ctx, &execution)
			if err != nil {
				s.logger.Errorf("trigger service: execution scraper component: error aborting test execution: %v", err)
				continue
			}
			s.metrics.IncAbortTest(execution.TestType, res.IsFailed())

			s.logger.Debugf("trigger service: execution scraper component: test execution %s is aborted", id)
			status.removeExecutionID(id)
		}
	}
}

func (s *Service) abortRunningTestSuiteExecutions(ctx context.Context, status *triggerStatus) {
	testSuiteExecutionIDs := status.getTestSuiteExecutionIDs()

	for _, id := range testSuiteExecutionIDs {
		execution, err := s.testResultRepository.Get(ctx, id)
		if err == mongo.ErrNoDocuments {
			s.logger.Warnf("trigger service: execution scraper component: no testsuite execution found for id %s", id)
			status.removeTestSuiteExecutionID(id)
			continue
		} else if err != nil {
			s.logger.Errorf("trigger service: execution scraper component: error fetching testsuite execution result: %v", err)
			continue
		}
		if execution.IsRunning() || execution.IsQueued() {
			err := s.eventsBus.PublishTopic(bus.InternalPublishTopic, testkube.NewEventEndTestSuiteAborted(&execution))
			if err != nil {
				s.logger.Errorf("trigger service: execution scraper component: error aborting test suite execution: %v", err)
				continue
			}

			s.logger.Debugf("trigger service: execution scraper component: testsuite execution %s is aborted", id)
			status.removeTestSuiteExecutionID(id)
		}
	}
}
