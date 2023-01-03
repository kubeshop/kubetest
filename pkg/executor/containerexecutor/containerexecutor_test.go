package containerexecutor

import (
	"context"
	"testing"
	"time"

	v3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/executor/client"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func TestExecuteAsync(t *testing.T) {
	ce := ContainerExecutor{
		clientSet:  getFakeClient("1"),
		log:        logger(),
		repository: FakeResultRepository{},
		metrics:    FakeMetricCounter{},
		emitter:    FakeEmitter{},
		namespace:  "default",
		configMap:  FakeConfigRepository{},
	}

	execution := &testkube.Execution{Id: "1"}
	options := client.ExecuteOptions{}
	res, err := ce.Execute(execution, options)
	assert.NoError(t, err)

	// Status is either running or passed, depending if async goroutine managed to finish
	assert.Contains(t,
		[]testkube.ExecutionStatus{testkube.RUNNING_ExecutionStatus, testkube.PASSED_ExecutionStatus},
		*res.Status)
}

func TestExecuteSync(t *testing.T) {
	ce := ContainerExecutor{
		clientSet:  getFakeClient("1"),
		log:        logger(),
		repository: FakeResultRepository{},
		metrics:    FakeMetricCounter{},
		emitter:    FakeEmitter{},
		namespace:  "default",
		configMap:  FakeConfigRepository{},
	}

	execution := &testkube.Execution{Id: "1"}
	options := client.ExecuteOptions{ImagePullSecretNames: []string{"secret-name1"}}
	res, err := ce.ExecuteSync(execution, options)
	assert.NoError(t, err)
	assert.Equal(t, testkube.PASSED_ExecutionStatus, *res.Status)
}

func TestNewExecutorJobSpecEmptyArgs(t *testing.T) {
	jobOptions := &JobOptions{
		Name:        "name",
		Namespace:   "namespace",
		InitImage:   "kubeshop/testkube-executor-init:0.7.10",
		Image:       "ubuntu",
		JobTemplate: defaultJobTemplate,
		Args:        []string{},
	}
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)
}

func TestNewExecutorJobSpecWithArgs(t *testing.T) {
	jobOptions := &JobOptions{
		Name:                  "name",
		Namespace:             "namespace",
		InitImage:             "kubeshop/testkube-executor-init:0.7.10",
		Image:                 "curl",
		JobTemplate:           defaultJobTemplate,
		ImagePullSecrets:      []string{"secret-name"},
		Command:               []string{"/bin/curl"},
		Args:                  []string{"-v", "https://testkube.kubeshop.io"},
		ActiveDeadlineSeconds: 100,
		Envs:                  map[string]string{"key": "value"},
		Variables:             map[string]testkube.Variable{"aa": testkube.Variable{Name: "name", Value: "value", Type_: testkube.VariableTypeBasic}},
	}
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	wantEnvs := []corev1.EnvVar{
		{Name: "DEBUG", Value: ""}, {Name: "RUNNER_ENDPOINT", Value: ""},
		{Name: "RUNNER_ACCESSKEYID", Value: ""}, {Name: "RUNNER_SECRETACCESSKEY", Value: ""},
		{Name: "RUNNER_LOCATION", Value: ""}, {Name: "RUNNER_TOKEN", Value: ""},
		{Name: "RUNNER_SSL", Value: ""}, {Name: "RUNNER_SCRAPPERENABLED", Value: ""},
		{Name: "RUNNER_DATADIR", Value: "/data"}, {Name: "NAME", Value: "value"},
		{Name: "key", Value: "value"},
	}

	assert.Equal(t, wantEnvs, spec.Spec.Template.Spec.Containers[0].Env)
}

func TestNewExecutorJobSpecWithoutInitImage(t *testing.T) {
	jobOptions := &JobOptions{
		Name:        "name",
		Namespace:   "namespace",
		InitImage:   "",
		Image:       "ubuntu",
		JobTemplate: defaultJobTemplate,
		Args:        []string{},
	}
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)
}

func TestNewExecutorJobSpecWithWorkingDirRelative(t *testing.T) {
	jobOptions, _ := NewJobOptions(
		executor.Images{},
		executor.Templates{},
		"",
		testkube.Execution{
			Id:            "name",
			TestName:      "name-test-1",
			TestNamespace: "namespace",
		},
		client.ExecuteOptions{
			TestSpec: v3.TestSpec{
				ExecutionRequest: &v3.ExecutionRequest{
					Image: "ubuntu",
				},
				Content: &v3.TestContent{
					Repository: &v3.Repository{
						WorkingDir: "relative/path",
					},
				},
			},
		},
	)
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	assert.Equal(t, repoPath+"/relative/path", spec.Spec.Template.Spec.Containers[0].WorkingDir)
}

func TestNewExecutorJobSpecWithWorkingDirAbsolute(t *testing.T) {
	jobOptions, _ := NewJobOptions(
		executor.Images{},
		executor.Templates{},
		"",
		testkube.Execution{
			Id:            "name",
			TestName:      "name-test-1",
			TestNamespace: "namespace",
		},
		client.ExecuteOptions{
			TestSpec: v3.TestSpec{
				ExecutionRequest: &v3.ExecutionRequest{
					Image: "ubuntu",
				},
				Content: &v3.TestContent{
					Repository: &v3.Repository{
						WorkingDir: "/absolute/path",
					},
				},
			},
		},
	)
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	assert.Equal(t, "/absolute/path", spec.Spec.Template.Spec.Containers[0].WorkingDir)
}

func TestNewExecutorJobSpecWithoutWorkingDir(t *testing.T) {
	jobOptions, _ := NewJobOptions(
		executor.Images{},
		executor.Templates{},
		"",
		testkube.Execution{
			Id:            "name",
			TestName:      "name-test-1",
			TestNamespace: "namespace",
		},
		client.ExecuteOptions{
			Namespace: "namespace",
			TestSpec: v3.TestSpec{
				ExecutionRequest: &v3.ExecutionRequest{
					Image: "ubuntu",
				},
				Content: &v3.TestContent{
					Repository: &v3.Repository{},
				},
			},
		},
	)
	spec, err := NewExecutorJobSpec(logger(), jobOptions)
	assert.NoError(t, err)
	assert.NotNil(t, spec)

	assert.Empty(t, spec.Spec.Template.Spec.Containers[0].WorkingDir)
}

func logger() *zap.SugaredLogger {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level = atomicLevel

	z, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}
	return z.Sugar()
}

func getFakeClient(executionID string) *fake.Clientset {
	initObjects := []runtime.Object{
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      executionID,
				Namespace: "default",
				Labels: map[string]string{
					"job-name": executionID,
				},
			},
			Status: corev1.PodStatus{
				Phase: corev1.PodSucceeded,
			},
		},
	}
	fakeClient := fake.NewSimpleClientset(initObjects...)
	return fakeClient
}

type FakeMetricCounter struct {
}

func (FakeMetricCounter) IncExecuteTest(execution testkube.Execution) {
	return
}

type FakeEmitter struct {
}

func (FakeEmitter) Notify(event testkube.Event) {
	return
}

type FakeResultRepository struct {
}

func (FakeResultRepository) Get(ctx context.Context, id string) (testkube.Execution, error) {
	return testkube.Execution{}, nil
}

func (FakeResultRepository) UpdateResult(ctx context.Context, id string, execution testkube.ExecutionResult) error {
	return nil
}
func (FakeResultRepository) StartExecution(ctx context.Context, id string, startTime time.Time) error {
	return nil
}
func (FakeResultRepository) EndExecution(ctx context.Context, execution testkube.Execution) error {
	return nil
}

type FakeConfigRepository struct {
}

func (FakeConfigRepository) GetUniqueClusterId(ctx context.Context) (string, error) {
	return "", nil
}

func (FakeConfigRepository) GetTelemetryEnabled(ctx context.Context) (ok bool, err error) {
	return false, nil
}

func (FakeConfigRepository) Get(ctx context.Context) (testkube.Config, error) {
	return testkube.Config{}, nil
}

func (FakeConfigRepository) Upsert(ctx context.Context, config testkube.Config) error {
	return nil
}
