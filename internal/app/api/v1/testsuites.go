package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	testsuitesv1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	"github.com/kubeshop/testkube/internal/pkg/api/datefilter"
	"github.com/kubeshop/testkube/internal/pkg/api/repository/testresult"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/crd"
	"github.com/kubeshop/testkube/pkg/cronjob"
	testsuitesmapper "github.com/kubeshop/testkube/pkg/mapper/testsuites"
	"github.com/kubeshop/testkube/pkg/types"
	"github.com/kubeshop/testkube/pkg/workerpool"
)

// GetTestSuiteHandler for getting test object
func (s TestkubeAPI) CreateTestSuiteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request testkube.TestSuiteUpsertRequest
		err := c.BodyParser(&request)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		if c.Accepts(mediaTypeJSON, mediaTypeYAML) == mediaTypeYAML {
			if request.Description != "" {
				request.Description = fmt.Sprintf("%q", request.Description)
			}

			data, err := crd.GenerateYAML(crd.TemplateTestSuite, []testkube.TestSuiteUpsertRequest{request})
			return s.getCRDs(c, data, err)
		}

		testSuite := mapTestSuiteUpsertRequestToTestCRD(request)
		testSuite.Namespace = s.Namespace

		s.Log.Infow("creating test suite", "testSuite", testSuite)

		created, err := s.TestsSuitesClient.Create(&testSuite)

		s.Metrics.IncCreateTestSuite(err)

		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		c.Status(http.StatusCreated)
		return c.JSON(created)
	}
}

// UpdateTestSuiteHandler updates an existing TestSuite CR based on TestSuite content
func (s TestkubeAPI) UpdateTestSuiteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request testkube.TestSuiteUpsertRequest
		err := c.BodyParser(&request)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		// we need to get resource first and load its metadata.ResourceVersion
		testSuite, err := s.TestsSuitesClient.Get(request.Name)
		if err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		// delete cron job, if schedule is cleaned
		if testSuite.Spec.Schedule != "" {
			cronJob, err := s.CronJobClient.Get(cronjob.GetMetadataName(request.Name, testSuiteResourceURI))
			if err != nil && !errors.IsNotFound(err) {
				return s.Error(c, http.StatusBadGateway, err)
			}

			if cronJob != nil {
				if request.Schedule == "" {
					if err = s.CronJobClient.Delete(cronjob.GetMetadataName(request.Name, testSuiteResourceURI)); err != nil {
						return s.Error(c, http.StatusBadGateway, err)
					}
				} else {
					if err = s.CronJobClient.UpdateLabels(cronJob, testSuite.Labels, request.Labels); err != nil {
						return s.Error(c, http.StatusBadGateway, err)
					}
				}
			}
		}

		// map TestSuite but load spec only to not override metadata.ResourceVersion
		testSuiteSpec := mapTestSuiteUpsertRequestToTestCRD(request)
		testSuite.Spec = testSuiteSpec.Spec
		testSuite.Labels = request.Labels
		testSuite, err = s.TestsSuitesClient.Update(testSuite)

		s.Metrics.IncUpdateTestSuite(err)

		if err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		return c.JSON(testSuite)
	}
}

// GetTestSuiteHandler for getting TestSuite object
func (s TestkubeAPI) GetTestSuiteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		crTestSuite, err := s.TestsSuitesClient.Get(name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		testSuite := testsuitesmapper.MapCRToAPI(*crTestSuite)
		if c.Accepts(mediaTypeJSON, mediaTypeYAML) == mediaTypeYAML {
			if testSuite.Description != "" {
				testSuite.Description = fmt.Sprintf("%q", testSuite.Description)
			}

			data, err := crd.GenerateYAML(crd.TemplateTestSuite, []testkube.TestSuite{testSuite})
			return s.getCRDs(c, data, err)
		}

		return c.JSON(testSuite)
	}
}

// GetTestSuiteWithExecutionHandler for getting TestSuite object with execution
func (s TestkubeAPI) GetTestSuiteWithExecutionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		crTestSuite, err := s.TestsSuitesClient.Get(name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		testSuite := testsuitesmapper.MapCRToAPI(*crTestSuite)
		if c.Accepts(mediaTypeJSON, mediaTypeYAML) == mediaTypeYAML {
			if testSuite.Description != "" {
				testSuite.Description = fmt.Sprintf("%q", testSuite.Description)
			}

			data, err := crd.GenerateYAML(crd.TemplateTestSuite, []testkube.TestSuite{testSuite})
			return s.getCRDs(c, data, err)
		}

		ctx := c.Context()
		startExecution, startErr := s.TestExecutionResults.GetLatestByTestSuite(ctx, name, "starttime")
		if startErr != nil && startErr != mongo.ErrNoDocuments {
			return s.Error(c, http.StatusInternalServerError, startErr)
		}

		endExecution, endErr := s.TestExecutionResults.GetLatestByTestSuite(ctx, name, "endtime")
		if endErr != nil && endErr != mongo.ErrNoDocuments {
			return s.Error(c, http.StatusInternalServerError, endErr)
		}

		testSuiteWithExecution := testkube.TestSuiteWithExecution{
			TestSuite: &testSuite,
		}
		if startErr == nil && endErr == nil {
			if startExecution.StartTime.After(endExecution.EndTime) {
				testSuiteWithExecution.LatestExecution = &startExecution
			} else {
				testSuiteWithExecution.LatestExecution = &endExecution
			}
		} else if startErr == nil {
			testSuiteWithExecution.LatestExecution = &startExecution
		} else if endErr == nil {
			testSuiteWithExecution.LatestExecution = &endExecution
		}

		return c.JSON(testSuiteWithExecution)
	}
}

// DeleteTestSuiteHandler for deleting a TestSuite with id
func (s TestkubeAPI) DeleteTestSuiteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("id")
		err := s.TestsSuitesClient.Delete(name)
		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		// delete cron job for test suite
		if err = s.CronJobClient.Delete(cronjob.GetMetadataName(name, testSuiteResourceURI)); err != nil {
			if !errors.IsNotFound(err) {
				return s.Error(c, http.StatusBadGateway, err)
			}
		}

		// delete executions for test
		if err = s.ExecutionResults.DeleteByTestSuite(c.Context(), name); err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		// delete executions for test suite
		if err = s.TestExecutionResults.DeleteByTestSuite(c.Context(), name); err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// DeleteTestSuitesHandler for deleting all TestSuites
func (s TestkubeAPI) DeleteTestSuitesHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var testSuiteNames []string
		selector := c.Query("selector")
		if selector == "" {
			err = s.TestsSuitesClient.DeleteAll()
		} else {
			testSuiteList, err := s.TestsSuitesClient.List(selector)
			if err != nil {
				if !errors.IsNotFound(err) {
					return s.Error(c, http.StatusBadGateway, err)
				}
			} else {
				for _, item := range testSuiteList.Items {
					testSuiteNames = append(testSuiteNames, item.Name)
				}
			}

			err = s.TestsSuitesClient.DeleteByLabels(selector)
		}

		if err != nil {
			if errors.IsNotFound(err) {
				return s.Warn(c, http.StatusNotFound, err)
			}

			return s.Error(c, http.StatusBadGateway, err)
		}

		// delete all cron jobs for test suites
		if err = s.CronJobClient.DeleteAll(testSuiteResourceURI, selector); err != nil {
			if !errors.IsNotFound(err) {
				return s.Error(c, http.StatusBadGateway, err)
			}
		}

		// delete all executions for tests
		if selector == "" {
			err = s.ExecutionResults.DeleteForAllTestSuites(c.Context())
		} else {
			err = s.ExecutionResults.DeleteByTestSuites(c.Context(), testSuiteNames)
		}

		if err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		// delete all executions for test suites
		if selector == "" {
			err = s.TestExecutionResults.DeleteAll(c.Context())
		} else {
			err = s.TestExecutionResults.DeleteByTestSuites(c.Context(), testSuiteNames)
		}

		if err != nil {
			return s.Error(c, http.StatusBadGateway, err)
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (s TestkubeAPI) getFilteredTestSuitesList(c *fiber.Ctx) (*testsuitesv1.TestSuiteList, error) {
	crTestSuites, err := s.TestsSuitesClient.List(c.Query("selector"))
	if err != nil {
		return nil, err
	}

	search := c.Query("textSearch")
	if search != "" {
		// filter items array
		for i := len(crTestSuites.Items) - 1; i >= 0; i-- {
			if !strings.Contains(crTestSuites.Items[i].Name, search) {
				crTestSuites.Items = append(crTestSuites.Items[:i], crTestSuites.Items[i+1:]...)
			}
		}
	}

	return crTestSuites, nil
}

// ListTestSuitesHandler for getting list of all available TestSuites
func (s TestkubeAPI) ListTestSuitesHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		crTestSuites, err := s.getFilteredTestSuitesList(c)
		if err != nil {
			return s.Error(c, http.StatusInternalServerError, err)
		}

		testSuites := testsuitesmapper.MapTestSuiteListKubeToAPI(*crTestSuites)
		if c.Accepts(mediaTypeJSON, mediaTypeYAML) == mediaTypeYAML {
			for i := range testSuites {
				if testSuites[i].Description != "" {
					testSuites[i].Description = fmt.Sprintf("%q", testSuites[i].Description)
				}
			}

			data, err := crd.GenerateYAML(crd.TemplateTestSuite, testSuites)
			return s.getCRDs(c, data, err)
		}

		return c.JSON(testSuites)
	}
}

// getLatestTestSuiteExecutions return latest test suite executions either by starttime or endtine for tests
func (s TestkubeAPI) getLatestTestSuiteExecutions(ctx context.Context, testSuiteNames []string) (map[string]testkube.TestSuiteExecution, error) {
	executions, err := s.TestExecutionResults.GetLatestByTestSuites(ctx, testSuiteNames, "starttime")
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	startExecutionMap := make(map[string]testkube.TestSuiteExecution, len(executions))
	for i := range executions {
		if executions[i].TestSuite == nil {
			continue
		}

		startExecutionMap[executions[i].TestSuite.Name] = executions[i]
	}

	executions, err = s.TestExecutionResults.GetLatestByTestSuites(ctx, testSuiteNames, "endtime")
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	endExecutionMap := make(map[string]testkube.TestSuiteExecution, len(executions))
	for i := range executions {
		if executions[i].TestSuite == nil {
			continue
		}

		endExecutionMap[executions[i].TestSuite.Name] = executions[i]
	}

	executionMap := make(map[string]testkube.TestSuiteExecution)
	for _, testSuiteName := range testSuiteNames {
		startExecution, okStart := startExecutionMap[testSuiteName]
		endExecution, okEnd := endExecutionMap[testSuiteName]
		if !okStart && !okEnd {
			continue
		}

		if okStart && !okEnd {
			executionMap[testSuiteName] = startExecution
			continue
		}

		if !okStart && okEnd {
			executionMap[testSuiteName] = endExecution
			continue
		}

		if startExecution.StartTime.After(endExecution.EndTime) {
			executionMap[testSuiteName] = startExecution
		} else {
			executionMap[testSuiteName] = endExecution
		}
	}

	return executionMap, nil
}

// ListTestSuiteWithExecutionsHandler for getting list of all available TestSuite with latest executions
func (s TestkubeAPI) ListTestSuiteWithExecutionsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		crTestSuites, err := s.getFilteredTestSuitesList(c)
		if err != nil {
			return s.Error(c, http.StatusInternalServerError, err)
		}

		testSuites := testsuitesmapper.MapTestSuiteListKubeToAPI(*crTestSuites)
		if c.Accepts(mediaTypeJSON, mediaTypeYAML) == mediaTypeYAML {
			for i := range testSuites {
				if testSuites[i].Description != "" {
					testSuites[i].Description = fmt.Sprintf("%q", testSuites[i].Description)
				}
			}

			data, err := crd.GenerateYAML(crd.TemplateTestSuite, testSuites)
			return s.getCRDs(c, data, err)
		}

		ctx := c.Context()
		testSuiteWithExecutions := make([]testkube.TestSuiteWithExecution, 0, len(testSuites))
		results := make([]testkube.TestSuiteWithExecution, 0, len(testSuites))
		testSuiteNames := make([]string, len(testSuites))
		for i := range testSuites {
			testSuiteNames[i] = testSuites[i].Name
		}

		executionMap, err := s.getLatestTestSuiteExecutions(ctx, testSuiteNames)
		if err != nil {
			return s.Error(c, http.StatusInternalServerError, err)
		}

		for i := range testSuites {
			if execution, ok := executionMap[testSuites[i].Name]; ok {
				results = append(results, testkube.TestSuiteWithExecution{
					TestSuite:       &testSuites[i],
					LatestExecution: &execution,
				})
			} else {
				testSuiteWithExecutions = append(testSuiteWithExecutions, testkube.TestSuiteWithExecution{
					TestSuite: &testSuites[i],
				})
			}
		}

		sort.Slice(testSuiteWithExecutions, func(i, j int) bool {
			return testSuiteWithExecutions[i].TestSuite.Created.After(testSuiteWithExecutions[j].TestSuite.Created)
		})

		sort.Slice(results, func(i, j int) bool {
			iTime := results[i].LatestExecution.EndTime
			if results[i].LatestExecution.StartTime.After(results[i].LatestExecution.EndTime) {
				iTime = results[i].LatestExecution.StartTime
			}

			jTime := results[j].LatestExecution.EndTime
			if results[j].LatestExecution.StartTime.After(results[j].LatestExecution.EndTime) {
				jTime = results[j].LatestExecution.StartTime
			}

			return iTime.After(jTime)
		})

		testSuiteWithExecutions = append(testSuiteWithExecutions, results...)
		status := c.Query("status")
		if status != "" {
			statusList, err := testkube.ParseTestSuiteExecutionStatusList(status, ",")
			if err != nil {
				return s.Error(c, http.StatusBadRequest, fmt.Errorf("test suite execution status filter invalid: %w", err))
			}

			statusMap := statusList.ToMap()
			// filter items array
			for i := len(testSuiteWithExecutions) - 1; i >= 0; i-- {
				if testSuiteWithExecutions[i].LatestExecution != nil && testSuiteWithExecutions[i].LatestExecution.Status != nil {
					if _, ok := statusMap[*testSuiteWithExecutions[i].LatestExecution.Status]; ok {
						continue
					}
				}

				testSuiteWithExecutions = append(testSuiteWithExecutions[:i], testSuiteWithExecutions[i+1:]...)
			}
		}

		return c.JSON(testSuiteWithExecutions)
	}
}

func (s TestkubeAPI) ExecuteTestSuitesHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		var request testkube.TestSuiteExecutionRequest
		err := c.BodyParser(&request)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, fmt.Errorf("test execution request body invalid: %w", err))
		}

		name := c.Params("id")
		namespace := c.Query("namespace", "testkube")
		selector := c.Query("selector")
		s.Log.Debugw("getting test suite", "name", name, "selector", selector)

		var testSuites []testsuitesv1.TestSuite
		if name != "" {
			testSuite, err := s.TestsSuitesClient.Get(name)
			if err != nil {
				if errors.IsNotFound(err) {
					return s.Warn(c, http.StatusNotFound, err)
				}

				return s.Error(c, http.StatusBadGateway, err)
			}

			testSuites = append(testSuites, *testSuite)
		} else {
			testSuiteList, err := s.TestsSuitesClient.List(selector)
			if err != nil {
				return s.Error(c, http.StatusInternalServerError, fmt.Errorf("can't get test suites: %w", err))
			}

			testSuites = append(testSuites, testSuiteList.Items...)
		}

		var results []testkube.TestSuiteExecution
		var work []testsuitesv1.TestSuite
		for _, testSuite := range testSuites {
			if testSuite.Spec.Schedule == "" || c.Query("callback") != "" {
				work = append(work, testSuite)
				continue
			}

			data, err := json.Marshal(request)
			if err != nil {
				return s.Error(c, http.StatusBadRequest, fmt.Errorf("can't prepare test suite request: %w", err))
			}

			options := cronjob.CronJobOptions{
				Schedule: testSuite.Spec.Schedule,
				Resource: testSuiteResourceURI,
				Data:     string(data),
				Labels:   testSuite.Labels,
			}
			if err = s.CronJobClient.Apply(testSuite.Name, cronjob.GetMetadataName(testSuite.Name, testSuiteResourceURI), options); err != nil {
				return s.Error(c, http.StatusInternalServerError, fmt.Errorf("can't create scheduled test suite: %w", err))
			}

			results = append(results, testkube.NewQueuedTestSuiteExecution(name, namespace))
		}

		if len(work) != 0 {
			concurrencyLevel, err := strconv.Atoi(c.Query("concurrency", defaultConcurrencyLevel))
			if err != nil {
				return s.Error(c, http.StatusBadRequest, fmt.Errorf("can't detect concurrency level: %w", err))
			}

			workerpoolService := workerpool.New[testkube.TestSuite, testkube.TestSuiteExecutionRequest, testkube.TestSuiteExecution](concurrencyLevel)

			go workerpoolService.SendRequests(s.prepareTestSuiteRequests(work, request))
			go workerpoolService.Run(ctx)

			for r := range workerpoolService.GetResponses() {
				results = append(results, r.Result)
			}
		}

		s.Log.Debugw("executing test", "name", name, "selector", selector)
		if name != "" && len(results) != 0 {
			if results[0].IsFailed() {
				return s.Error(c, http.StatusInternalServerError, fmt.Errorf("Test suite failed %v", name))
			}

			c.Status(http.StatusCreated)
			return c.JSON(results[0])
		}

		c.Status(http.StatusCreated)
		return c.JSON(results)
	}
}

func (s TestkubeAPI) prepareTestSuiteRequests(work []testsuitesv1.TestSuite, request testkube.TestSuiteExecutionRequest) []workerpool.Request[
	testkube.TestSuite, testkube.TestSuiteExecutionRequest, testkube.TestSuiteExecution] {
	requests := make([]workerpool.Request[testkube.TestSuite, testkube.TestSuiteExecutionRequest, testkube.TestSuiteExecution], len(work))
	for i := range work {
		requests[i] = workerpool.Request[testkube.TestSuite, testkube.TestSuiteExecutionRequest, testkube.TestSuiteExecution]{
			Object:  testsuitesmapper.MapCRToAPI(work[i]),
			Options: request,
			ExecFn:  s.executeTestSuite,
		}
	}
	return requests
}

func (s TestkubeAPI) ListTestSuiteExecutionsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		filter := getExecutionsFilterFromRequest(c)

		executionsTotals, err := s.TestExecutionResults.GetExecutionsTotals(ctx, filter)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}
		allExecutionsTotals, err := s.TestExecutionResults.GetExecutionsTotals(ctx)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		executions, err := s.TestExecutionResults.GetExecutions(ctx, filter)
		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		return c.JSON(testkube.TestSuiteExecutionsResult{
			Totals:   &allExecutionsTotals,
			Filtered: &executionsTotals,
			Results:  mapToTestExecutionSummary(executions),
		})
	}
}

func (s TestkubeAPI) GetTestSuiteExecutionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		id := c.Params("executionID")
		execution, err := s.TestExecutionResults.Get(ctx, id)

		if err != nil {
			return s.Error(c, http.StatusBadRequest, err)
		}

		execution.Duration = types.FormatDuration(execution.Duration)

		secretMap := make(map[string]string)
		if execution.SecretUUID != "" && execution.TestSuite != nil {
			secretMap, err = s.TestsSuitesClient.GetSecretTestSuiteVars(execution.TestSuite.Name, execution.SecretUUID)
			if err != nil {
				return s.Error(c, http.StatusInternalServerError, err)
			}
		}

		for key, value := range secretMap {
			if variable, ok := execution.Variables[key]; ok {
				variable.Value = string(value)
				variable.SecretRef = nil
				execution.Variables[key] = variable
			}
		}

		return c.JSON(execution)
	}
}

func (s TestkubeAPI) executeTestSuite(ctx context.Context, testSuite testkube.TestSuite, request testkube.TestSuiteExecutionRequest) (
	testsuiteExecution testkube.TestSuiteExecution, err error) {
	s.Log.Debugw("Got test to execute", "test", testSuite)
	secretUUID, err := s.TestsSuitesClient.GetCurrentSecretUUID(testSuite.Name)
	if err != nil {
		return testsuiteExecution, err
	}

	request.SecretUUID = secretUUID
	testsuiteExecution = testkube.NewStartedTestSuiteExecution(testSuite, request)
	err = s.TestExecutionResults.Insert(ctx, testsuiteExecution)
	if err != nil {
		s.Log.Infow("Inserting test execution", "error", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(testsuiteExecution *testkube.TestSuiteExecution, request testkube.TestSuiteExecutionRequest) {
		defer func(testExecution *testkube.TestSuiteExecution) {
			duration := testExecution.CalculateDuration()
			testExecution.EndTime = time.Now()
			testExecution.Duration = duration.String()

			err = s.TestExecutionResults.EndExecution(ctx, testExecution.Id, testExecution.EndTime, duration)
			if err != nil {
				s.Log.Errorw("error setting end time", "error", err.Error())
			}

			wg.Done()
		}(testsuiteExecution)

		hasFailedSteps := false
		cancellSteps := false
		for i := range testsuiteExecution.StepResults {
			if cancellSteps {
				testsuiteExecution.StepResults[i].Execution.ExecutionResult.Cancel()
				continue
			}

			// start execution of given step
			testsuiteExecution.StepResults[i].Execution.ExecutionResult.InProgress()
			err = s.TestExecutionResults.Update(ctx, *testsuiteExecution)
			if err != nil {
				s.Log.Infow("Updating test execution", "error", err)
			}

			s.executeTestStep(ctx, *testsuiteExecution, request, &testsuiteExecution.StepResults[i])

			err := s.TestExecutionResults.Update(ctx, *testsuiteExecution)
			if err != nil {
				hasFailedSteps = true
				s.Log.Errorw("saving test suite execution results error", "error", err)
				continue
			}

			if testsuiteExecution.StepResults[i].IsFailed() {
				hasFailedSteps = true
				if testsuiteExecution.StepResults[i].Step.StopTestOnFailure {
					cancellSteps = true
					continue
				}
			}
		}

		testsuiteExecution.Status = testkube.TestSuiteExecutionStatusPassed
		if hasFailedSteps {
			testsuiteExecution.Status = testkube.TestSuiteExecutionStatusFailed
		}

		s.Metrics.IncExecuteTestSuite(*testsuiteExecution)

		err := s.TestExecutionResults.Update(ctx, *testsuiteExecution)
		if err != nil {
			s.Log.Errorw("saving final test suite execution result error", "error", err)
		}

	}(&testsuiteExecution, request)

	// wait for sync test suite execution
	if request.Sync {
		wg.Wait()
	}

	return testsuiteExecution, nil
}

func (s TestkubeAPI) executeTestStep(ctx context.Context, testsuiteExecution testkube.TestSuiteExecution,
	request testkube.TestSuiteExecutionRequest, result *testkube.TestSuiteStepExecutionResult) {

	var testSuiteName string
	if testsuiteExecution.TestSuite != nil {
		testSuiteName = testsuiteExecution.TestSuite.Name
	}

	step := result.Step

	l := s.Log.With("type", step.Type(), "testSuiteName", testSuiteName, "name", step.FullName())

	switch step.Type() {

	case testkube.TestSuiteStepTypeExecuteTest:
		executeTestStep := step.Execute
		request := testkube.ExecutionRequest{
			Name:                fmt.Sprintf("%s-%s", testSuiteName, executeTestStep.Name),
			TestSuiteName:       testSuiteName,
			Namespace:           executeTestStep.Namespace,
			Variables:           testsuiteExecution.Variables,
			TestSuiteSecretUUID: request.SecretUUID,
			Sync:                true,
			HttpProxy:           request.HttpProxy,
			HttpsProxy:          request.HttpsProxy,
			ExecutionLabels:     request.ExecutionLabels,
		}

		l.Info("executing test", "variables", testsuiteExecution.Variables, "request", request)
		execution, err := s.executeTest(ctx, testkube.Test{Name: executeTestStep.Name}, request)
		if err != nil {
			result.Err(err)
			return
		}
		result.Execution = &execution

	case testkube.TestSuiteStepTypeDelay:
		l.Debug("delaying execution")
		time.Sleep(time.Millisecond * time.Duration(step.Delay.Duration))
		result.Execution.ExecutionResult.Success()

	default:
		result.Err(fmt.Errorf("can't find handler for execution step type: '%v'", step.Type()))
	}
}

func getExecutionsFilterFromRequest(c *fiber.Ctx) testresult.Filter {

	filter := testresult.NewExecutionsFilter()
	name := c.Query("id", "")
	if name != "" {
		filter = filter.WithName(name)
	}

	textSearch := c.Query("textSearch", "")
	if textSearch != "" {
		filter = filter.WithTextSearch(textSearch)
	}

	page, err := strconv.Atoi(c.Query("page", ""))
	if err == nil {
		filter = filter.WithPage(page)
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", ""))
	if err == nil && pageSize != 0 {
		filter = filter.WithPageSize(pageSize)
	}

	status := c.Query("status", "")
	if status != "" {
		filter = filter.WithStatus(status)
	}

	dFilter := datefilter.NewDateFilter(c.Query("startDate", ""), c.Query("endDate", ""))
	if dFilter.IsStartValid {
		filter = filter.WithStartDate(dFilter.Start)
	}

	if dFilter.IsEndValid {
		filter = filter.WithEndDate(dFilter.End)
	}

	selector := c.Query("selector")
	if selector != "" {
		filter = filter.WithSelector(selector)
	}

	return filter
}

// TODO move to testuites mapper
func mapToTestExecutionSummary(executions []testkube.TestSuiteExecution) []testkube.TestSuiteExecutionSummary {
	result := make([]testkube.TestSuiteExecutionSummary, len(executions))

	for i, execution := range executions {
		executionsSummary := make([]testkube.TestSuiteStepExecutionSummary, len(execution.StepResults))
		for j, stepResult := range execution.StepResults {
			executionsSummary[j] = mapStepResultToExecutionSummary(stepResult)
		}

		result[i] = testkube.TestSuiteExecutionSummary{
			Id:            execution.Id,
			Name:          execution.Name,
			TestSuiteName: execution.TestSuite.Name,
			Status:        execution.Status,
			StartTime:     execution.StartTime,
			EndTime:       execution.EndTime,
			Duration:      types.FormatDuration(execution.Duration),
			Execution:     executionsSummary,
			Labels:        execution.Labels,
		}
	}

	return result
}

func mapStepResultToExecutionSummary(r testkube.TestSuiteStepExecutionResult) testkube.TestSuiteStepExecutionSummary {
	var id, testName, name string
	var status *testkube.ExecutionStatus = testkube.ExecutionStatusPassed
	var stepType *testkube.TestSuiteStepType

	if r.Test != nil {
		testName = r.Test.Name
	}

	if r.Execution != nil {
		id = r.Execution.Id
		if r.Execution.ExecutionResult != nil {
			status = r.Execution.ExecutionResult.Status
		}
	}

	if r.Step != nil {
		stepType = r.Step.Type()
		name = r.Step.FullName()
	}

	return testkube.TestSuiteStepExecutionSummary{
		Id:       id,
		Name:     name,
		TestName: testName,
		Status:   status,
		Type_:    stepType,
	}
}

func mapTestSuiteUpsertRequestToTestCRD(request testkube.TestSuiteUpsertRequest) testsuitesv1.TestSuite {
	return testsuitesv1.TestSuite{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
			Labels:    request.Labels,
		},
		Spec: testsuitesv1.TestSuiteSpec{
			Repeats:     int(request.Repeats),
			Description: request.Description,
			Before:      mapTestStepsToCRD(request.Before),
			Steps:       mapTestStepsToCRD(request.Steps),
			After:       mapTestStepsToCRD(request.After),
			Schedule:    request.Schedule,
			Variables:   testsuitesmapper.MapCRDVariables(request.Variables),
		},
	}
}

func mapTestStepsToCRD(steps []testkube.TestSuiteStep) (out []testsuitesv1.TestSuiteStepSpec) {
	for _, step := range steps {
		out = append(out, mapTestStepToCRD(step))
	}

	return
}

func mapTestStepToCRD(step testkube.TestSuiteStep) (stepSpec testsuitesv1.TestSuiteStepSpec) {
	switch step.Type() {

	case testkube.TestSuiteStepTypeDelay:
		stepSpec.Delay = &testsuitesv1.TestSuiteStepDelay{
			Duration: step.Delay.Duration,
		}

	case testkube.TestSuiteStepTypeExecuteTest:
		s := step.Execute
		stepSpec.Execute = &testsuitesv1.TestSuiteStepExecute{
			Namespace: s.Namespace,
			Name:      s.Name,
			// TODO move StopOnFailure level up in operator model to mimic this one
			StopOnFailure: step.StopTestOnFailure,
		}
	}

	return
}
