package testworkflowprocessor

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowconfig"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/action/actiontypes"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/action/actiontypes/lite"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/constants"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/stage"
)

type BundleOptions struct {
	Secrets []corev1.Secret
	Config  testworkflowconfig.InternalConfig
}

type Bundle struct {
	Secrets       []corev1.Secret
	ConfigMaps    []corev1.ConfigMap
	Job           batchv1.Job
	Signature     []stage.Signature
	FullSignature []stage.Signature
}

func (b *Bundle) Actions() (actions actiontypes.ActionGroups) {
	_ = json.Unmarshal([]byte(b.Job.Spec.Template.Annotations[constants.SpecAnnotationName]), &actions)
	return
}

func (b *Bundle) LiteActions() (actions lite.LiteActionGroups) {
	_ = json.Unmarshal([]byte(b.Job.Spec.Template.Annotations[constants.SpecAnnotationName]), &actions)
	return
}

func (b *Bundle) Deploy(ctx context.Context, clientSet kubernetes.Interface, namespace string) (err error) {
	if b.Job.Namespace != "" {
		namespace = b.Job.Namespace
	}
	for _, item := range b.Secrets {
		_, err = clientSet.CoreV1().Secrets(namespace).Create(ctx, &item, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(err, "failed to deploy secrets")
		}
	}
	for _, item := range b.ConfigMaps {
		_, err = clientSet.CoreV1().ConfigMaps(namespace).Create(ctx, &item, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(err, "failed to deploy config maps")
		}
	}
	_, err = clientSet.BatchV1().Jobs(namespace).Create(ctx, &b.Job, metav1.CreateOptions{})
	return errors.Wrap(err, "failed to deploy job")
}
