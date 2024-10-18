package v1

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	testtriggersclientv1 "github.com/kubeshop/testkube-operator/pkg/client/testtriggers/v1"
	testworkflowsv1 "github.com/kubeshop/testkube-operator/pkg/client/testworkflows/v1"
	"github.com/kubeshop/testkube/cmd/api-server/commons"
	"github.com/kubeshop/testkube/internal/config"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/log"
	repoConfig "github.com/kubeshop/testkube/pkg/repository/config"
	"github.com/kubeshop/testkube/pkg/repository/testworkflow"
	"github.com/kubeshop/testkube/pkg/secretmanager"
	"github.com/kubeshop/testkube/pkg/testworkflows/executionworker/executionworkertypes"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowexecutor"

	"k8s.io/client-go/kubernetes"

	"github.com/gofiber/fiber/v2"

	executorsclientv1 "github.com/kubeshop/testkube-operator/pkg/client/executors/v1"
	"github.com/kubeshop/testkube/internal/app/api/metrics"
	"github.com/kubeshop/testkube/pkg/event"
	"github.com/kubeshop/testkube/pkg/event/bus"
	"github.com/kubeshop/testkube/pkg/event/kind/slack"
	ws "github.com/kubeshop/testkube/pkg/event/kind/websocket"
	"github.com/kubeshop/testkube/pkg/executor/client"
	"github.com/kubeshop/testkube/pkg/featureflags"
	logsclient "github.com/kubeshop/testkube/pkg/logs/client"
	"github.com/kubeshop/testkube/pkg/scheduler"
	"github.com/kubeshop/testkube/pkg/secret"
	"github.com/kubeshop/testkube/pkg/server"
	"github.com/kubeshop/testkube/pkg/storage"
)

func NewTestkubeAPI(
	clusterId string,
	deprecatedRepositories commons.DeprecatedRepositories,
	deprecatedClients commons.DeprecatedClients,
	namespace string,
	testWorkflowResults testworkflow.Repository,
	testWorkflowOutput testworkflow.OutputRepository,
	secretClient secret.Interface,
	secretManager secretmanager.SecretManager,
	webhookClient *executorsclientv1.WebhooksClient,
	clientset kubernetes.Interface,
	testTriggersClient testtriggersclientv1.Interface,
	testWorkflowsClient testworkflowsv1.Interface,
	testWorkflowTemplatesClient testworkflowsv1.TestWorkflowTemplatesInterface,
	configMap repoConfig.Repository,
	eventsEmitter *event.Emitter,
	websocketLoader *ws.WebsocketLoader,
	executor client.Executor,
	containerExecutor client.Executor,
	testWorkflowExecutor testworkflowexecutor.TestWorkflowExecutor,
	executionWorkerClient executionworkertypes.Worker,
	metrics metrics.Metrics,
	scheduler *scheduler.Scheduler,
	slackLoader *slack.SlackLoader,
	graphqlPort int,
	artifactsStorage storage.ArtifactsStorage,
	dashboardURI string,
	helmchartVersion string,
	mode string,
	eventsBus bus.Bus,
	secretConfig testkube.SecretConfig,
	ff featureflags.FeatureFlags,
	logsStream logsclient.Stream,
	logGrpcClient logsclient.StreamGetter,
	serviceAccountNames map[string]string,
	dockerImageVersion string,
	proContext *config.ProContext,
	storageParams StorageParams,
) TestkubeAPI {

	return TestkubeAPI{
		ClusterID:                   clusterId,
		Log:                         log.DefaultLogger,
		DeprecatedRepositories:      deprecatedRepositories,
		DeprecatedClients:           deprecatedClients,
		TestWorkflowResults:         testWorkflowResults,
		TestWorkflowOutput:          testWorkflowOutput,
		SecretClient:                secretClient,
		SecretManager:               secretManager,
		Clientset:                   clientset,
		TestTriggersClient:          testTriggersClient,
		TestWorkflowsClient:         testWorkflowsClient,
		TestWorkflowTemplatesClient: testWorkflowTemplatesClient,
		Metrics:                     metrics,
		WebsocketLoader:             websocketLoader,
		Events:                      eventsEmitter,
		WebhooksClient:              webhookClient,
		Namespace:                   namespace,
		ConfigMap:                   configMap,
		Executor:                    executor,
		ContainerExecutor:           containerExecutor,
		TestWorkflowExecutor:        testWorkflowExecutor,
		ExecutionWorkerClient:       executionWorkerClient,
		storageParams:               storageParams,
		scheduler:                   scheduler,
		slackLoader:                 slackLoader,
		graphqlPort:                 graphqlPort,
		ArtifactsStorage:            artifactsStorage,
		dashboardURI:                dashboardURI,
		helmchartVersion:            helmchartVersion,
		mode:                        mode,
		eventsBus:                   eventsBus,
		secretConfig:                secretConfig,
		featureFlags:                ff,
		logsStream:                  logsStream,
		logGrpcClient:               logGrpcClient,
		ServiceAccountNames:         serviceAccountNames,
		dockerImageVersion:          dockerImageVersion,
		proContext:                  proContext,
	}
}

type TestkubeAPI struct {
	ClusterID                   string
	Log                         *zap.SugaredLogger
	TestWorkflowResults         testworkflow.Repository
	TestWorkflowOutput          testworkflow.OutputRepository
	Executor                    client.Executor
	ContainerExecutor           client.Executor
	TestWorkflowExecutor        testworkflowexecutor.TestWorkflowExecutor
	ExecutionWorkerClient       executionworkertypes.Worker
	DeprecatedRepositories      commons.DeprecatedRepositories
	DeprecatedClients           commons.DeprecatedClients
	SecretClient                secret.Interface
	SecretManager               secretmanager.SecretManager
	WebhooksClient              *executorsclientv1.WebhooksClient
	TestTriggersClient          testtriggersclientv1.Interface
	TestWorkflowsClient         testworkflowsv1.Interface
	TestWorkflowTemplatesClient testworkflowsv1.TestWorkflowTemplatesInterface
	Metrics                     metrics.Metrics
	storageParams               StorageParams
	Namespace                   string
	WebsocketLoader             *ws.WebsocketLoader
	Events                      *event.Emitter
	ConfigMap                   repoConfig.Repository
	scheduler                   *scheduler.Scheduler
	Clientset                   kubernetes.Interface
	slackLoader                 *slack.SlackLoader
	graphqlPort                 int
	ArtifactsStorage            storage.ArtifactsStorage
	dashboardURI                string
	helmchartVersion            string
	mode                        string
	eventsBus                   bus.Bus
	secretConfig                testkube.SecretConfig
	featureFlags                featureflags.FeatureFlags
	logsStream                  logsclient.Stream
	logGrpcClient               logsclient.StreamGetter
	proContext                  *config.ProContext
	ServiceAccountNames         map[string]string
	dockerImageVersion          string
}

type StorageParams struct {
	SSL             bool   `envconfig:"STORAGE_SSL" default:"false"`
	SkipVerify      bool   `envconfig:"STORAGE_SKIP_VERIFY" default:"false"`
	CertFile        string `envconfig:"STORAGE_CERT_FILE"`
	KeyFile         string `envconfig:"STORAGE_KEY_FILE"`
	CAFile          string `envconfig:"STORAGE_CA_FILE"`
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	Token           string
	Bucket          string
}

func (s *TestkubeAPI) Init(server server.HTTPServer) {
	// TODO: Consider extracting outside?
	server.Routes.Get("/info", s.InfoHandler())
	server.Routes.Get("/debug", s.DebugHandler())

	root := server.Routes

	webhooks := root.Group("/webhooks")

	webhooks.Post("/", s.CreateWebhookHandler())
	webhooks.Patch("/:name", s.UpdateWebhookHandler())
	webhooks.Get("/", s.ListWebhooksHandler())
	webhooks.Get("/:name", s.GetWebhookHandler())
	webhooks.Delete("/:name", s.DeleteWebhookHandler())
	webhooks.Delete("/", s.DeleteWebhooksHandler())

	testWorkflows := root.Group("/test-workflows")
	testWorkflows.Get("/", s.ListTestWorkflowsHandler())
	testWorkflows.Post("/", s.CreateTestWorkflowHandler())
	testWorkflows.Delete("/", s.DeleteTestWorkflowsHandler())
	testWorkflows.Get("/:id", s.GetTestWorkflowHandler())
	testWorkflows.Put("/:id", s.UpdateTestWorkflowHandler())
	testWorkflows.Delete("/:id", s.DeleteTestWorkflowHandler())
	testWorkflows.Get("/:id/executions", s.ListTestWorkflowExecutionsHandler())
	testWorkflows.Post("/:id/executions", s.ExecuteTestWorkflowHandler())
	testWorkflows.Get("/:id/tags", s.ListTagsHandler())
	testWorkflows.Get("/:id/metrics", s.GetTestWorkflowMetricsHandler())
	testWorkflows.Get("/:id/executions/:executionID", s.GetTestWorkflowExecutionHandler())
	testWorkflows.Post("/:id/abort", s.AbortAllTestWorkflowExecutionsHandler())
	testWorkflows.Post("/:id/executions/:executionID/abort", s.AbortTestWorkflowExecutionHandler())
	testWorkflows.Post("/:id/executions/:executionID/pause", s.PauseTestWorkflowExecutionHandler())
	testWorkflows.Post("/:id/executions/:executionID/resume", s.ResumeTestWorkflowExecutionHandler())
	testWorkflows.Get("/:id/executions/:executionID/logs", s.GetTestWorkflowExecutionLogsHandler())

	testWorkflowExecutions := root.Group("/test-workflow-executions")
	testWorkflowExecutions.Get("/", s.ListTestWorkflowExecutionsHandler())
	testWorkflowExecutions.Get("/:executionID", s.GetTestWorkflowExecutionHandler())
	testWorkflowExecutions.Get("/:executionID/notifications", s.StreamTestWorkflowExecutionNotificationsHandler())
	testWorkflowExecutions.Get("/:executionID/notifications/stream", s.StreamTestWorkflowExecutionNotificationsWebSocketHandler())
	testWorkflowExecutions.Post("/:executionID/abort", s.AbortTestWorkflowExecutionHandler())
	testWorkflowExecutions.Post("/:executionID/pause", s.PauseTestWorkflowExecutionHandler())
	testWorkflowExecutions.Post("/:executionID/resume", s.ResumeTestWorkflowExecutionHandler())
	testWorkflowExecutions.Get("/:executionID/logs", s.GetTestWorkflowExecutionLogsHandler())
	testWorkflowExecutions.Get("/:executionID/artifacts", s.ListTestWorkflowExecutionArtifactsHandler())
	testWorkflowExecutions.Get("/:executionID/artifacts/:filename", s.GetTestWorkflowArtifactHandler())
	testWorkflowExecutions.Get("/:executionID/artifact-archive", s.GetTestWorkflowArtifactArchiveHandler())

	testWorkflowWithExecutions := root.Group("/test-workflow-with-executions")
	testWorkflowWithExecutions.Get("/", s.ListTestWorkflowWithExecutionsHandler())
	testWorkflowWithExecutions.Get("/:id", s.GetTestWorkflowWithExecutionHandler())
	testWorkflowWithExecutions.Get("/:id/tags", s.ListTagsHandler())

	root.Post("/preview-test-workflow", s.PreviewTestWorkflowHandler())

	testWorkflowTemplates := root.Group("/test-workflow-templates")
	testWorkflowTemplates.Get("/", s.ListTestWorkflowTemplatesHandler())
	testWorkflowTemplates.Post("/", s.CreateTestWorkflowTemplateHandler())
	testWorkflowTemplates.Delete("/", s.DeleteTestWorkflowTemplatesHandler())
	testWorkflowTemplates.Get("/:id", s.GetTestWorkflowTemplateHandler())
	testWorkflowTemplates.Put("/:id", s.UpdateTestWorkflowTemplateHandler())
	testWorkflowTemplates.Delete("/:id", s.DeleteTestWorkflowTemplateHandler())

	testTriggers := root.Group("/triggers")
	testTriggers.Get("/", s.ListTestTriggersHandler())
	testTriggers.Post("/", s.CreateTestTriggerHandler())
	testTriggers.Patch("/", s.BulkUpdateTestTriggersHandler())
	testTriggers.Delete("/", s.DeleteTestTriggersHandler())
	testTriggers.Get("/:id", s.GetTestTriggerHandler())
	testTriggers.Patch("/:id", s.UpdateTestTriggerHandler())
	testTriggers.Delete("/:id", s.DeleteTestTriggerHandler())

	keymap := root.Group("/keymap")
	keymap.Get("/triggers", s.GetTestTriggerKeyMapHandler())

	labels := root.Group("/labels")
	labels.Get("/", s.ListLabelsHandler())

	tags := root.Group("/tags")
	tags.Get("/", s.ListTagsHandler())

	events := root.Group("/events")
	events.Post("/flux", s.FluxEventHandler())
	events.Get("/stream", s.EventsStreamHandler())

	configs := root.Group("/config")
	configs.Get("/", s.GetConfigsHandler())
	configs.Patch("/", s.UpdateConfigsHandler())

	debug := root.Group("/debug")
	debug.Get("/listeners", s.GetDebugListenersHandler())

	secrets := root.Group("/secrets")
	secrets.Get("/", s.ListSecretsHandler())
	secrets.Post("/", s.CreateSecretHandler())
	secrets.Get("/:id", s.GetSecretHandler())
	secrets.Delete("/:id", s.DeleteSecretHandler())
	secrets.Patch("/:id", s.UpdateSecretHandler())

	repositories := root.Group("/repositories")
	repositories.Post("/", s.ValidateRepositoryHandler())

	// set up proxy for the internal GraphQL server
	server.Mux.All("/graphql", func(c *fiber.Ctx) error {
		// Connect to server
		serverConn, err := net.Dial("tcp", fmt.Sprintf(":%d", s.graphqlPort))
		if err != nil {
			s.Log.Errorw("could not connect to GraphQL server as a proxy", "error", err)
			return err
		}

		// Resend headers to the server
		_, err = serverConn.Write(c.Request().Header.Header())
		if err != nil {
			serverConn.Close()
			s.Log.Errorw("error while sending headers to GraphQL server", "error", err)
			return err
		}

		// Resend body to the server
		_, err = serverConn.Write(c.Body())
		if err != nil && err != io.EOF {
			serverConn.Close()
			s.Log.Errorw("error while reading GraphQL client data", "error", err)
			return err
		}

		// Handle optional WebSocket connection
		c.Context().HijackSetNoResponse(true)
		c.Context().Hijack(func(clientConn net.Conn) {
			// Close the connection afterward
			defer serverConn.Close()
			defer clientConn.Close()

			// Extract Unix connection
			serverSock, ok := serverConn.(*net.TCPConn)
			if !ok {
				s.Log.Errorw("error while building TCPConn out ouf serverConn", "error", err)
				return
			}
			clientSock, ok := reflect.Indirect(reflect.ValueOf(clientConn)).FieldByName("Conn").Interface().(*net.TCPConn)
			if !ok {
				s.Log.Errorw("error while building TCPConn out of hijacked connection", "error", err)
				return
			}

			// Duplex communication between client and GraphQL server
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				_, err := io.Copy(clientSock, serverSock)
				if err != nil && err != io.EOF && !errors.Is(err, syscall.ECONNRESET) && !errors.Is(err, syscall.EPIPE) {
					s.Log.Errorw("error while reading GraphQL client data", "error", err)
				}
				serverSock.CloseWrite()
			}()
			go func() {
				defer wg.Done()
				_, err = io.Copy(serverSock, clientSock)
				if err != nil && err != io.EOF {
					s.Log.Errorw("error while reading GraphQL server data", "error", err)
				}
				clientSock.CloseWrite()
			}()
			wg.Wait()
		})
		return nil
	})
}
