package webhook

import (
	"fmt"

	"go.uber.org/zap"

	executorsclientv1 "github.com/kubeshop/testkube-operator/pkg/client/executors/v1"
	"github.com/kubeshop/testkube/cmd/api-server/commons"
	v1 "github.com/kubeshop/testkube/internal/app/api/metrics"
	"github.com/kubeshop/testkube/internal/config"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/event/kind/common"
	"github.com/kubeshop/testkube/pkg/mapper/webhooks"
	"github.com/kubeshop/testkube/pkg/repository/testworkflow"
)

var _ common.ListenerLoader = (*WebhooksLoader)(nil)

func NewWebhookLoader(log *zap.SugaredLogger, webhooksClient executorsclientv1.WebhooksInterface, deprecatedClients commons.DeprecatedClients,
	deprecatedRepositories commons.DeprecatedRepositories, testWorkflowExecutionResults testworkflow.Repository,
	metrics v1.Metrics, proContext *config.ProContext, envs map[string]string,
) *WebhooksLoader {
	return &WebhooksLoader{
		log:                          log,
		WebhooksClient:               webhooksClient,
		deprecatedClients:            deprecatedClients,
		deprecatedRepositories:       deprecatedRepositories,
		testWorkflowExecutionResults: testWorkflowExecutionResults,
		metrics:                      metrics,
		proContext:                   proContext,
		envs:                         envs,
	}
}

type WebhooksLoader struct {
	log                          *zap.SugaredLogger
	WebhooksClient               executorsclientv1.WebhooksInterface
	deprecatedClients            commons.DeprecatedClients
	deprecatedRepositories       commons.DeprecatedRepositories
	testWorkflowExecutionResults testworkflow.Repository
	metrics                      v1.Metrics
	proContext                   *config.ProContext
	envs                         map[string]string
}

func (r WebhooksLoader) Kind() string {
	return "webhook"
}

func (r WebhooksLoader) Load() (listeners common.Listeners, err error) {
	// load all webhooks from kubernetes CRDs
	webhookList, err := r.WebhooksClient.List("")
	if err != nil {
		return listeners, err
	}

	// and create listeners for each webhook spec
	for _, webhook := range webhookList.Items {
		payloadTemplate := ""
		if webhook.Spec.PayloadTemplateReference != "" {
			if r.deprecatedClients == nil {
				r.log.Errorw("webhook using deprecated PayloadTemplateReference", "name", webhook.Name, "template", webhook.Spec.PayloadTemplateReference)
				continue
			}
			template, err := r.deprecatedClients.Templates().Get(webhook.Spec.PayloadTemplateReference)
			if err != nil {
				return listeners, err
			}

			if template.Spec.Type_ != nil && testkube.TemplateType(*template.Spec.Type_) == testkube.WEBHOOK_TemplateType {
				payloadTemplate = template.Spec.Body
			} else {
				r.log.Warnw("not matching template type", "template", webhook.Spec.PayloadTemplateReference)
			}
		}

		if webhook.Spec.PayloadTemplate != "" {
			payloadTemplate = webhook.Spec.PayloadTemplate
		}

		types := webhooks.MapEventArrayToCRDEvents(webhook.Spec.Events)
		name := fmt.Sprintf("%s.%s", webhook.ObjectMeta.Namespace, webhook.ObjectMeta.Name)
		listeners = append(
			listeners,
			NewWebhookListener(
				name, webhook.Spec.Uri, webhook.Spec.Selector, types,
				webhook.Spec.PayloadObjectField, payloadTemplate, webhook.Spec.Headers, webhook.Spec.Disabled,
				r.deprecatedRepositories, r.testWorkflowExecutionResults,
				r.metrics, r.proContext, r.envs,
			),
		)
	}

	return listeners, nil
}
