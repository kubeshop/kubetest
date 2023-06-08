package testsuites

import (
	corev1 "k8s.io/api/core/v1"

	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
	testsuitesv3 "github.com/kubeshop/testkube-operator/apis/testsuite/v3"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

// MapTestSuiteListKubeToAPI maps TestSuiteList CRD to list of OpenAPI spec TestSuite
func MapTestSuiteListKubeToAPI(cr testsuitesv3.TestSuiteList) (tests []testkube.TestSuite) {
	tests = make([]testkube.TestSuite, len(cr.Items))
	for i, item := range cr.Items {
		tests[i] = MapCRToAPI(item)
	}

	return
}

// MapCRToAPI maps TestSuite CRD to OpenAPI spec TestSuite
func MapCRToAPI(cr testsuitesv3.TestSuite) (test testkube.TestSuite) {
	test.Name = cr.Name
	test.Namespace = cr.Namespace
	var batches = []struct {
		source *[]testsuitesv3.TestSuiteBatchStep
		dest   *[]testkube.TestSuiteBatchStep
	}{
		{
			source: &cr.Spec.Before,
			dest:   &test.Before,
		},
		{
			source: &cr.Spec.Steps,
			dest:   &test.Steps,
		},
		{
			source: &cr.Spec.After,
			dest:   &test.After,
		},
	}

	for i := range batches {
		for _, b := range *batches[i].source {
			steps := make([]testkube.TestSuiteStep, len(b.Execute))
			for j := range b.Execute {
				steps[j] = mapCRStepToAPI(b.Execute[j])
			}

			*batches[i].dest = append(*batches[i].dest, testkube.TestSuiteBatchStep{
				StopOnFailure: b.StopOnFailure,
				Execute:       steps,
			})
		}
	}

	test.Description = cr.Spec.Description
	test.Repeats = int32(cr.Spec.Repeats)
	test.Labels = cr.Labels
	test.Schedule = cr.Spec.Schedule
	test.Created = cr.CreationTimestamp.Time
	test.ExecutionRequest = MapExecutionRequestFromSpec(cr.Spec.ExecutionRequest)
	test.Status = MapStatusFromSpec(cr.Status)
	return
}

// mapCRStepToAPI maps CRD TestSuiteStepSpec to OpenAPI spec TestSuiteStep
func mapCRStepToAPI(crstep testsuitesv3.TestSuiteStepSpec) (teststep testkube.TestSuiteStep) {

	switch true {
	case crstep.Test != nil:
		teststep = testkube.TestSuiteStep{
			Test: &testkube.TestSuiteStepExecuteTest{
				Name: crstep.Test.Name,
			},
		}

	case crstep.Delay != nil:
		teststep = testkube.TestSuiteStep{
			Delay: &testkube.TestSuiteStepDelay{
				Duration: crstep.Delay.Duration.String(),
			},
		}
	}

	return
}

// @Depracated
// MapDepratcatedParams maps old params to new variables data structure
func MapDepratcatedParams(in map[string]testkube.Variable) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		out[k] = v.Value
	}
	return out
}

// MapCRDVariables maps variables between API and operator CRDs
// TODO if we could merge operator into testkube repository we would get rid of those mappings
func MapCRDVariables(in map[string]testkube.Variable) map[string]testsuitesv3.Variable {
	out := map[string]testsuitesv3.Variable{}
	for k, v := range in {
		variable := testsuitesv3.Variable{
			Name:  v.Name,
			Type_: string(*v.Type_),
			Value: v.Value,
		}

		if v.SecretRef != nil {
			variable.ValueFrom = corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: v.SecretRef.Name,
					},
					Key: v.SecretRef.Key,
				},
			}
		}

		if v.ConfigMapRef != nil {
			variable.ValueFrom = corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: v.ConfigMapRef.Name,
					},
					Key: v.ConfigMapRef.Key,
				},
			}
		}

		out[k] = variable
	}
	return out
}

func MergeVariablesAndParams(variables map[string]testsuitesv3.Variable, params map[string]string) map[string]testkube.Variable {
	out := map[string]testkube.Variable{}
	for k, v := range params {
		out[k] = testkube.NewBasicVariable(k, v)
	}

	for k, v := range variables {
		if v.Type_ == commonv1.VariableTypeSecret {
			if v.ValueFrom.SecretKeyRef == nil {
				out[k] = testkube.NewSecretVariable(v.Name, v.Value)
			} else {
				out[k] = testkube.NewSecretVariableReference(v.Name, v.ValueFrom.SecretKeyRef.Name, v.ValueFrom.SecretKeyRef.Key)
			}
		}
		if v.Type_ == commonv1.VariableTypeBasic {
			if v.ValueFrom.ConfigMapKeyRef == nil {
				out[k] = testkube.NewBasicVariable(v.Name, v.Value)
			} else {
				out[k] = testkube.NewConfigMapVariableReference(v.Name, v.ValueFrom.ConfigMapKeyRef.Name, v.ValueFrom.ConfigMapKeyRef.Key)
			}
		}
	}

	return out
}

// MapExecutionRequestFromSpec maps CRD to OpenAPI spec ExecutionRequest
func MapExecutionRequestFromSpec(specExecutionRequest *testsuitesv3.TestSuiteExecutionRequest) *testkube.TestSuiteExecutionRequest {
	if specExecutionRequest == nil {
		return nil
	}

	return &testkube.TestSuiteExecutionRequest{
		Name:            specExecutionRequest.Name,
		Labels:          specExecutionRequest.Labels,
		ExecutionLabels: specExecutionRequest.ExecutionLabels,
		Namespace:       specExecutionRequest.Namespace,
		Variables:       MergeVariablesAndParams(specExecutionRequest.Variables, nil),
		SecretUUID:      specExecutionRequest.SecretUUID,
		Sync:            specExecutionRequest.Sync,
		HttpProxy:       specExecutionRequest.HttpProxy,
		HttpsProxy:      specExecutionRequest.HttpsProxy,
		Timeout:         specExecutionRequest.Timeout,
		CronJobTemplate: specExecutionRequest.CronJobTemplate,
	}
}

// MapStatusFromSpec maps CRD to OpenAPI spec TestSuiteStatus
func MapStatusFromSpec(specStatus testsuitesv3.TestSuiteStatus) *testkube.TestSuiteStatus {
	if specStatus.LatestExecution == nil {
		return nil
	}

	return &testkube.TestSuiteStatus{
		LatestExecution: &testkube.TestSuiteExecutionCore{
			Id:        specStatus.LatestExecution.Id,
			Status:    (*testkube.TestSuiteExecutionStatus)(specStatus.LatestExecution.Status),
			StartTime: specStatus.LatestExecution.StartTime.Time,
			EndTime:   specStatus.LatestExecution.EndTime.Time,
		},
	}
}
