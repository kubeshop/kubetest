package env

import (
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/constants"
	"github.com/kubeshop/testkube/pkg/ui"
)

var (
	UseProxyValue = false
)

type envObjectStorageConfig struct {
	Endpoint        string `envconfig:"TK_OS_ENDPOINT"`
	AccessKeyID     string `envconfig:"TK_OS_ACCESSKEY"`
	SecretAccessKey string `envconfig:"TK_OS_SECRETKEY"`
	Region          string `envconfig:"TK_OS_REGION"`
	Token           string `envconfig:"TK_OS_TOKEN"`
	Bucket          string `envconfig:"TK_OS_BUCKET"`
	Ssl             bool   `envconfig:"TK_OS_SSL" default:"false"`
	SkipVerify      bool   `envconfig:"TK_OS_SSL_SKIP_VERIFY" default:"false"`
	CertFile        string `envconfig:"TK_OS_CERT_FILE"`
	KeyFile         string `envconfig:"TK_OS_KEY_FILE"`
	CAFile          string `envconfig:"TK_OS_CA_FILE"`
}

type envCloudConfig struct {
	Url         string `envconfig:"TK_C_URL"`
	ApiKey      string `envconfig:"TK_C_KEY"`
	SkipVerify  bool   `envconfig:"TK_C_SKIP_VERIFY" default:"false"`
	TlsInsecure bool   `envconfig:"TK_C_TLS_INSECURE" default:"false"`
}

type envExecutionConfig struct {
	WorkflowName    string `envconfig:"TK_WF"`
	Id              string `envconfig:"TK_EX"`
	Name            string `envconfig:"TK_EXN"`
	Number          int64  `envconfig:"TK_EXC"`
	ScheduledAt     string `envconfig:"TK_EXS"`
	ResourceId      string `envconfig:"TK_EXI"`
	RootResourceId  string `envconfig:"TK_EXR"`
	FSPrefix        string `envconfig:"TK_FS"`
	disableWebhooks bool   `envconfig:"TK_DWH"`
}

type envSystemConfig struct {
	Debug                 string `envconfig:"DEBUG"`
	Ref                   string `envconfig:"TK_REF"`
	Namespace             string `envconfig:"TK_NS"`
	DefaultRegistry       string `envconfig:"TK_R"`
	DefaultServiceAccount string `envconfig:"TK_SA"`
	Ip                    string `envconfig:"TK_IP"`
	DashboardUrl          string `envconfig:"TK_DASH"`
	ApiUrl                string `envconfig:"TK_API"`
	ClusterID             string `envconfig:"TK_CLU"`
	CDEventsTarget        string `envconfig:"TK_CDE"`
}

type envImagesConfig struct {
	Init                         string `envconfig:"TESTKUBE_TW_INIT_IMAGE"`
	Toolkit                      string `envconfig:"TESTKUBE_TW_TOOLKIT_IMAGE"`
	InspectorPersistenceEnabled  bool   `envconfig:"TK_IMG_P" default:"false"`
	InspectorPersistenceCacheKey string `envconfig:"TK_IMG_PK"`
}

type featuresConfig struct {
	EnableJUnitParser bool `envconfig:"TK_FF_JUNIT_REPORT" default:"false"`
}

type envConfig struct {
	System        envSystemConfig
	ObjectStorage envObjectStorageConfig
	Cloud         envCloudConfig
	Execution     envExecutionConfig
	Images        envImagesConfig
	Features      featuresConfig
}

var cfg envConfig
var cfgLoaded = false

func Config() *envConfig {
	if !cfgLoaded {
		err := envconfig.Process("", &cfg.System)
		ui.ExitOnError("configuring environment", err)
		err = envconfig.Process("", &cfg.ObjectStorage)
		ui.ExitOnError("configuring environment", err)
		err = envconfig.Process("", &cfg.Cloud)
		ui.ExitOnError("configuring environment", err)
		err = envconfig.Process("", &cfg.Execution)
		ui.ExitOnError("configuring environment", err)
		err = envconfig.Process("", &cfg.Images)
		ui.ExitOnError("configuring environment", err)
		err = envconfig.Process("", &cfg.Features)
		ui.ExitOnError("configuring environment", err)
	}
	cfgLoaded = true
	return &cfg
}

func Debug() bool {
	return Config().System.Debug == "1"
}

func CloudEnabled() bool {
	return Config().Cloud.ApiKey != ""
}

func UseProxy() bool {
	return UseProxyValue
}

func Ref() string {
	return Config().System.Ref
}

func Namespace() string {
	return Config().System.Namespace
}

func IP() string {
	return Config().System.Ip
}

func WorkflowName() string {
	return Config().Execution.WorkflowName
}

func ExecutionId() string {
	return Config().Execution.Id
}

func ExecutionName() string {
	return Config().Execution.Name
}

func ExecutionNumber() int64 {
	return Config().Execution.Number
}

func ExecutionScheduledAt() time.Time {
	t, _ := time.Parse(constants.RFC3339Millis, Config().Execution.ScheduledAt)
	return t
}

func ExecutionDisableWebhooks() bool {
	return Config().Execution.disableWebhooks
}

func JUnitParserEnabled() bool {
	return Config().Features.EnableJUnitParser
}
