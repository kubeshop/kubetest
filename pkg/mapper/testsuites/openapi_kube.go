package testsuites

import (
	testsuitesv2 "github.com/kubeshop/testkube-operator/apis/testsuite/v2"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO move to testuites mapper
func MapToTestExecutionSummary(executions []testkube.TestSuiteExecution) []testkube.TestSuiteExecutionSummary {
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
			DurationMs:    types.FormatDurationMs(execution.Duration),
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

func MapTestSuiteUpsertRequestToTestCRD(request testkube.TestSuiteUpsertRequest) testsuitesv2.TestSuite {
	return testsuitesv2.TestSuite{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
			Labels:    request.Labels,
		},
		Spec: testsuitesv2.TestSuiteSpec{
			Repeats:          int(request.Repeats),
			Description:      request.Description,
			Before:           mapTestStepsToCRD(request.Before),
			Steps:            mapTestStepsToCRD(request.Steps),
			After:            mapTestStepsToCRD(request.After),
			Schedule:         request.Schedule,
			ExecutionRequest: MapExecutionRequestToSpecExecutionRequest(request.ExecutionRequest),
		},
	}
}

func mapTestStepsToCRD(steps []testkube.TestSuiteStep) (out []testsuitesv2.TestSuiteStepSpec) {
	for _, step := range steps {
		out = append(out, mapTestStepToCRD(step))
	}

	return out
}

func mapTestStepToCRD(step testkube.TestSuiteStep) (stepSpec testsuitesv2.TestSuiteStepSpec) {
	switch step.Type() {

	case testkube.TestSuiteStepTypeDelay:
		stepSpec.Delay = &testsuitesv2.TestSuiteStepDelay{
			Duration: step.Delay.Duration,
		}

	case testkube.TestSuiteStepTypeExecuteTest:
		s := step.Execute
		stepSpec.Execute = &testsuitesv2.TestSuiteStepExecute{
			Namespace: s.Namespace,
			Name:      s.Name,
			// TODO move StopOnFailure level up in operator model to mimic this one
			StopOnFailure: step.StopTestOnFailure,
		}
	}

	return
}

// MapExecutionRequestToSpecExecutionRequest maps ExecutionRequest OpenAPI spec to ExecutionRequest CRD spec
func MapExecutionRequestToSpecExecutionRequest(executionRequest *testkube.TestSuiteExecutionRequest) *testsuitesv2.TestSuiteExecutionRequest {
	if executionRequest == nil {
		return nil
	}

	return &testsuitesv2.TestSuiteExecutionRequest{
		Name:            executionRequest.Name,
		Labels:          executionRequest.Labels,
		ExecutionLabels: executionRequest.ExecutionLabels,
		Namespace:       executionRequest.Namespace,
		Variables:       MapCRDVariables(executionRequest.Variables),
		SecretUUID:      executionRequest.SecretUUID,
		Sync:            executionRequest.Sync,
		HttpProxy:       executionRequest.HttpProxy,
		HttpsProxy:      executionRequest.HttpsProxy,
		Timeout:         executionRequest.Timeout,
	}
}

//MapTestSuiteUpsertRequestToTestCRD maps TestSuiteUpdateRequest OpenAPI spec to TestSuite CRD spec
func MapTestSuiteUpdateRequestToTestCRD(request testkube.TestSuiteUpdateRequest, testSuite *testsuitesv2.TestSuite) *testsuitesv2.TestSuite {
	var fields = []struct {
		source      *string
		destination *string
	}{
		{
			request.Name,
			&testSuite.Name,
		},
		{
			request.Namespace,
			&testSuite.Namespace,
		},
		{
			request.Description,
			&testSuite.Spec.Description,
		},
		{
			request.Schedule,
			&testSuite.Spec.Schedule,
		},
	}

	for _, field := range fields {
		if field.source != nil {
			*field.destination = *field.source
		}
	}

	if request.Before != nil {
		testSuite.Spec.Before = mapTestStepsToCRD(*request.Before)
	}

	if request.Steps != nil {
		testSuite.Spec.Steps = mapTestStepsToCRD(*request.Steps)
	}

	if request.After != nil {
		testSuite.Spec.After = mapTestStepsToCRD(*request.After)
	}

	if request.Labels != nil {
		testSuite.Labels = *request.Labels
	}

	if request.Repeats != nil {
		testSuite.Spec.Repeats = int(*request.Repeats)
	}

	if request.ExecutionRequest != nil {
		testSuite.Spec.ExecutionRequest = MapExecutionUpdateRequestToSpecExecutionRequest(*request.ExecutionRequest, testSuite.Spec.ExecutionRequest)
	}

	return testSuite
}

// MapExecutionUpdateRequestToSpecExecutionRequest maps ExecutionUpdateRequest OpenAPI spec to ExecutionRequest CRD spec
func MapExecutionUpdateRequestToSpecExecutionRequest(executionRequest *testkube.TestSuiteExecutionUpdateRequest,
	request *testsuitesv2.TestSuiteExecutionRequest) *testsuitesv2.TestSuiteExecutionRequest {
	if executionRequest == nil {
		return nil
	}

	if request == nil {
		request = &testsuitesv2.TestSuiteExecutionRequest{}
	}

	var fields = []struct {
		source      *string
		destination *string
	}{
		{
			executionRequest.Name,
			&request.Name,
		},
		{
			executionRequest.Namespace,
			&request.Namespace,
		},
		{
			executionRequest.SecretUUID,
			&request.SecretUUID,
		},
		{
			executionRequest.HttpProxy,
			&request.HttpProxy,
		},
		{
			executionRequest.HttpsProxy,
			&request.HttpsProxy,
		},
	}

	for _, field := range fields {
		if field.source != nil {
			*field.destination = *field.source
		}
	}

	if executionRequest.Labels != nil {
		request.Labels = *executionRequest.Labels
	}

	if executionRequest.ExecutionLabels != nil {
		request.ExecutionLabels = *executionRequest.ExecutionLabels
	}

	if executionRequest.Sync != nil {
		request.Sync = *executionRequest.Sync
	}

	if executionRequest.Timeout != nil {
		request.Timeout = *executionRequest.Timeout
	}

	if executionRequest.Variables != nil {
		request.Variables = MapCRDVariables(*executionRequest.Variables)
	}

	return request
}
