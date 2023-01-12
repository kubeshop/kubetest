package pods

import (
	testtriggersv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	corev1 "k8s.io/api/core/v1"
)

// MapCRDConditionsToAPI maps Pod CRD conditions to OpenAPI spec TestTriggerConditions
func MapCRDConditionsToAPI(conditions []corev1.PodCondition) []testtriggersv1.TestTriggerCondition {
	var results []testtriggersv1.TestTriggerCondition
	for _, condition := range conditions {
		results = append(results, testtriggersv1.TestTriggerCondition{
			Type_:  string(condition.Type),
			Status: (*testtriggersv1.TestTriggerConditionStatuses)(&condition.Status),
		})
	}

	return results
}
