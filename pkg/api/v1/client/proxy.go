package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/problem"
)

// check in compile time if interface is implemented
var _ Client = (*ProxyAPIClient)(nil)

func GetClientSet() (clientset kubernetes.Interface, err error) {
	clcfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return clientset, err
	}

	restcfg, err := clientcmd.NewNonInteractiveClientConfig(
		*clcfg, "", &clientcmd.ConfigOverrides{}, nil).ClientConfig()
	if err != nil {
		return clientset, err
	}

	return kubernetes.NewForConfig(restcfg)
}

func NewProxyAPIClient(client kubernetes.Interface, config ProxyConfig) ProxyAPIClient {
	return ProxyAPIClient{
		client: client,
		config: config,
	}
}

func NewProxyConfig(namespace string) ProxyConfig {
	return ProxyConfig{
		Namespace:   namespace,
		ServiceName: "testkube-api-server",
		ServicePort: 8088,
	}
}

type ProxyConfig struct {
	// Namespace where testkube is installed
	Namespace string
	// API Server service name
	ServiceName string
	// API Server service port
	ServicePort int
}

type ProxyAPIClient struct {
	client kubernetes.Interface
	config ProxyConfig
}

// tests and executions -----------------------------------------------------------------------------

func (c ProxyAPIClient) GetTest(id, namespace string) (test testkube.Test, err error) {
	uri := c.getURI("/tests/%s", id)
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("namespace", namespace)

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return test, fmt.Errorf("api/get-test returned error: %w", err)
	}

	return c.getTestFromResponse(resp)
}

func (c ProxyAPIClient) GetExecution(executionID string) (execution testkube.Execution, err error) {

	uri := c.getURI("/executions/%s", executionID)

	req := c.GetProxy("GET").Suffix(uri)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/get-execution returned error: %w", err)
	}

	return c.getExecutionFromResponse(resp)
}

// ListExecutions list all executions for given test name
func (c ProxyAPIClient) ListExecutions(id string, limit int, tags []string) (executions testkube.ExecutionsResult, err error) {

	uri := c.getURI("/executions/")

	if id != "" {
		uri = fmt.Sprintf("/tests/%s/executions", id)
	}

	req := c.GetProxy("GET").
		Suffix(uri).
		Param("pageSize", fmt.Sprintf("%d", limit))

	if len(tags) > 0 {
		req.Param("tags", strings.Join(tags, ","))
	}

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return executions, fmt.Errorf("api/get-executions returned error: %w", err)
	}

	return c.getExecutionsFromResponse(resp)
}

func (c ProxyAPIClient) DeleteTests(namespace string) error {
	uri := c.getURI("/tests")
	return c.makeDeleteRequest(uri, namespace, true)
}

func (c ProxyAPIClient) DeleteTest(name string, namespace string) error {
	if name == "" {
		return fmt.Errorf("test name '%s' is not valid", name)
	}
	uri := c.getURI("/tests/%s", name)
	return c.makeDeleteRequest(uri, namespace, true)
}

// CreateTest creates new Test Custom Resource
func (c ProxyAPIClient) CreateTest(options UpsertTestOptions) (test testkube.Test, err error) {
	uri := c.getURI("/tests")

	request := testkube.TestUpsertRequest(options)

	body, err := json.Marshal(request)
	if err != nil {
		return test, err
	}

	req := c.GetProxy("POST").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return test, fmt.Errorf("api/create-test returned error: %w", err)
	}

	return c.getTestFromResponse(resp)
}

// UpdateTest Test Custom Resource
func (c ProxyAPIClient) UpdateTest(options UpsertTestOptions) (test testkube.Test, err error) {
	uri := c.getURI("/tests/%s", options.Name)

	request := testkube.TestUpsertRequest(options)

	body, err := json.Marshal(request)
	if err != nil {
		return test, err
	}

	req := c.GetProxy("PATCH").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return test, fmt.Errorf("api/udpate-test returned error: %w", err)
	}

	return c.getTestFromResponse(resp)
}

// ExecuteTest starts test execution, reads data and returns ID
// Execution is started asynchronously client can check later for results
func (c ProxyAPIClient) ExecuteTest(id, namespace, executionName string, executionParams map[string]string, executionParamsFileContent string, args []string) (execution testkube.Execution, err error) {
	uri := c.getURI("/tests/%s/executions", id)

	// get test to get test tags
	test, err := c.GetTest(id, namespace)
	if err != nil {
		return execution, nil
	}

	request := testkube.ExecutionRequest{
		Name:       executionName,
		Namespace:  namespace,
		ParamsFile: executionParamsFileContent,
		Params:     executionParams,
		Tags:       test.Tags,
		Args:       args,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return execution, err
	}

	req := c.GetProxy("POST").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/execute-test returned error: %w", err)
	}

	return c.getExecutionFromResponse(resp)
}

func (c ProxyAPIClient) Logs(id string) (logs chan output.Output, err error) {
	logs = make(chan output.Output)
	uri := c.getURI("/executions/%s/logs", id)

	resp, err := c.GetProxy("GET").
		Suffix(uri).
		SetHeader("Accept", "text/event-stream").
		Stream(context.Background())

	go func() {
		defer close(logs)
		defer resp.Close()

		StreamToLogsChannel(resp, logs)
	}()

	return
}

// GetExecutions list all executions in given test
func (c ProxyAPIClient) ListTests(namespace string, tags []string) (tests testkube.Tests, err error) {
	uri := c.getURI("/tests")
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("namespace", namespace)

	if len(tags) > 0 {
		req.Param("tags", strings.Join(tags, ","))
	}

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return tests, fmt.Errorf("api/list-tests returned error: %w", err)
	}

	return c.getTestsFromResponse(resp)
}

// GetExecutions list all executions in given test
func (c ProxyAPIClient) AbortExecution(testID, id string) error {
	uri := c.getURI("/tests/%s/executions/%s", testID, id)
	return c.makeDeleteRequest(uri, "testkube", false)
}

// executor --------------------------------------------------------------------------------

func (c ProxyAPIClient) CreateExecutor(options CreateExecutorOptions) (executor testkube.ExecutorDetails, err error) {
	uri := c.getURI("/executors")

	request := testkube.ExecutorCreateRequest(options)

	body, err := json.Marshal(request)
	if err != nil {
		return executor, err
	}

	req := c.GetProxy("POST").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return executor, fmt.Errorf("api/create-test returned error: %w", err)
	}

	return c.getExecutorDetailsFromResponse(resp)
}

func (c ProxyAPIClient) GetExecutor(name string) (executor testkube.ExecutorDetails, err error) {
	uri := c.getURI("/executors/%s", name)
	req := c.GetProxy("GET").Suffix(uri)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return executor, fmt.Errorf("api/get-executor returned error: %w", err)
	}

	return c.getExecutorDetailsFromResponse(resp)
}

func (c ProxyAPIClient) ListExecutors() (executors testkube.ExecutorsDetails, err error) {
	uri := c.getURI("/executors")
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("namespace", "testkube")

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return executors, fmt.Errorf("api/list-executors returned error: %w", err)
	}

	return c.getExecutorsDetailsFromResponse(resp)
}

func (c ProxyAPIClient) DeleteExecutor(name string) (err error) {
	uri := c.getURI("/executors/%s", name)
	return c.makeDeleteRequest(uri, "testkube", false)
}

// maintenance --------------------------------------------------------------------------------

func (c ProxyAPIClient) GetServerInfo() (info testkube.ServerInfo, err error) {
	uri := c.getURI("/info")
	req := c.GetProxy("GET").Suffix(uri)
	resp := req.Do(context.Background())
	if resp.Error() != nil {
		return info, resp.Error()
	}

	bytes, err := resp.Raw()
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(bytes, &info)

	return

}

func (c ProxyAPIClient) GetProxy(requestType string) *rest.Request {
	return c.client.CoreV1().RESTClient().Verb(requestType).
		Namespace(c.config.Namespace).
		Resource("services").
		SetHeader("Content-Type", "application/json").
		Name(fmt.Sprintf("%s:%d", c.config.ServiceName, c.config.ServicePort)).
		SubResource("proxy")
}

func (c ProxyAPIClient) getExecutionFromResponse(resp rest.Result) (execution testkube.Execution, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return execution, err
	}

	err = json.Unmarshal(bytes, &execution)

	return execution, err
}

func (c ProxyAPIClient) getExecutionsFromResponse(resp rest.Result) (executions testkube.ExecutionsResult, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return executions, err
	}

	err = json.Unmarshal(bytes, &executions)

	return executions, err
}

func (c ProxyAPIClient) getTestsFromResponse(resp rest.Result) (tests testkube.Tests, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return tests, err
	}

	err = json.Unmarshal(bytes, &tests)

	return tests, err
}

func (c ProxyAPIClient) getExecutorsDetailsFromResponse(resp rest.Result) (executors testkube.ExecutorsDetails, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return executors, err
	}

	err = json.Unmarshal(bytes, &executors)

	return executors, err
}

func (c ProxyAPIClient) getTestFromResponse(resp rest.Result) (test testkube.Test, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return test, err
	}

	err = json.Unmarshal(bytes, &test)

	return test, err
}

func (c ProxyAPIClient) getExecutorDetailsFromResponse(resp rest.Result) (executor testkube.ExecutorDetails, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return executor, err
	}

	err = json.Unmarshal(bytes, &executor)

	return executor, err
}

func (c ProxyAPIClient) getProblemFromResponse(resp rest.Result) (problem.Problem, error) {
	bytes, respErr := resp.Raw()

	problemResponse := problem.Problem{}
	err := json.Unmarshal(bytes, &problemResponse)

	// add kubeAPI client error to details
	if respErr != nil {
		problemResponse.Detail += ";\nresp error:" + respErr.Error()
	}

	return problemResponse, err
}

// responseError tries to lookup if response is of Problem type
func (c ProxyAPIClient) responseError(resp rest.Result) error {
	if resp.Error() != nil {
		pr, err := c.getProblemFromResponse(resp)

		// if can't process response return content from response
		if err != nil {
			content, _ := resp.Raw()
			return fmt.Errorf("api server response: '%s'\nerror: %w", content, resp.Error())
		}

		return fmt.Errorf("api server problem: %s", pr.Detail)
	}

	return nil
}

func (c ProxyAPIClient) getURI(pathTemplate string, params ...interface{}) string {
	path := fmt.Sprintf(pathTemplate, params...)
	return fmt.Sprintf("%s%s", Version, path)
}

func (c ProxyAPIClient) makeDeleteRequest(uri string, namespace string, isContentExpected bool) error {

	req := c.GetProxy("DELETE").
		Suffix(uri).
		Param("namespace", namespace)
	resp := req.Do(context.Background())

	if resp.Error() != nil {
		return resp.Error()
	}

	if err := c.responseError(resp); err != nil {
		return err
	}

	if isContentExpected {
		var code int
		resp.StatusCode(&code)
		if code != http.StatusNoContent {
			respBody, err := resp.Raw()
			if err != nil {
				return err
			}
			return fmt.Errorf("request returned error: %s", respBody)
		}
	}

	return nil
}

func (c ProxyAPIClient) GetExecutionArtifacts(executionID string) (artifacts testkube.Artifacts, err error) {
	uri := c.getURI("/executions/%s/artifacts", executionID)
	req := c.GetProxy("GET").
		Suffix(uri)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return artifacts, fmt.Errorf("api/list-artifacts returned error: %w", err)
	}

	return c.getArtifactsFromResponse(resp)

}

func (c ProxyAPIClient) DownloadFile(executionID, fileName, destination string) (artifact string, err error) {
	uri := c.getURI("/executions/%s/artifacts/%s", executionID, url.QueryEscape(fileName))
	req, err := c.GetProxy("GET").
		Suffix(uri).
		SetHeader("Accept", "text/event-stream").
		Stream(context.Background())
	if err != nil {
		return "", err
	}

	defer req.Close()

	f, err := os.Create(filepath.Join(destination, filepath.Base(fileName)))
	if err != nil {
		return "", err
	}

	if _, err := f.ReadFrom(req); err != nil {
		return "", err
	}

	defer f.Close()
	return f.Name(), err
}

func (c ProxyAPIClient) getArtifactsFromResponse(resp rest.Result) (artifacts []testkube.Artifact, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return artifacts, err
	}

	err = json.Unmarshal(bytes, &artifacts)

	return artifacts, err
}

// --------------- test suites --------------------------

func (c ProxyAPIClient) GetTestSuite(id, namespace string) (test testkube.TestSuite, err error) {
	uri := c.getURI("/test-suites/%s", id)
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("namespace", namespace)

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return test, fmt.Errorf("api/get-test returned error: %w", err)
	}

	return c.getTestSuiteFromResponse(resp)
}

func (c ProxyAPIClient) DeleteTestSuite(name string, namespace string) error {
	if name == "" {
		return fmt.Errorf("testsuite name '%s' is not valid", name)
	}
	uri := c.getURI("/test-suites/%s", name)
	return c.makeDeleteRequest(uri, namespace, true)
}

func (c ProxyAPIClient) DeleteTestSuites(namespace string) error {
	uri := c.getURI("/test-suites")
	return c.makeDeleteRequest(uri, namespace, true)
}

func (c ProxyAPIClient) ListTestSuites(namespace string, tags []string) (testSuites testkube.TestSuites, err error) {
	uri := c.getURI("/test-suites")
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("namespace", namespace)

	if len(tags) > 0 {
		req.Param("tags", strings.Join(tags, ","))
	}

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return testSuites, fmt.Errorf("api/list-test-suites returned error: %w", err)
	}

	return c.getTestSuitesFromResponse(resp)
}

// CreateTestSuite creates new TestSuite Custom Resource
func (c ProxyAPIClient) CreateTestSuite(options UpsertTestSuiteOptions) (testSuite testkube.TestSuite, err error) {
	uri := c.getURI("/test-suites")

	request := testkube.TestSuiteUpsertRequest(options)

	body, err := json.Marshal(request)
	if err != nil {
		return testSuite, err
	}

	req := c.GetProxy("POST").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return testSuite, fmt.Errorf("api/create-test-suite returned error: %w", err)
	}

	return c.getTestSuiteFromResponse(resp)
}

// UpdateTestSuite creates new TestSuite Custom Resource
func (c ProxyAPIClient) UpdateTestSuite(options UpsertTestSuiteOptions) (testSuite testkube.TestSuite, err error) {
	uri := c.getURI("/test-suites/%s", options.Name)

	request := testkube.TestSuiteUpsertRequest(options)

	body, err := json.Marshal(request)
	if err != nil {
		return testSuite, err
	}

	req := c.GetProxy("PATCH").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return testSuite, fmt.Errorf("api/udpate-test-suite returned error: %w", err)
	}

	return c.getTestSuiteFromResponse(resp)
}

func (c ProxyAPIClient) getTestSuiteFromResponse(resp rest.Result) (testSuite testkube.TestSuite, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return testSuite, err
	}

	err = json.Unmarshal(bytes, &testSuite)

	return testSuite, err
}

// ExecuteTestSuite starts new external test suite execution, reads data and returns ID
// Execution is started asynchronously client can check later for results
func (c ProxyAPIClient) ExecuteTestSuite(id, namespace, executionName string, executionParams map[string]string) (execution testkube.TestSuiteExecution, err error) {
	uri := c.getURI("/test-suites/%s/executions", id)

	request := testkube.ExecutionRequest{
		Name:      executionName,
		Namespace: namespace,
		Params:    executionParams,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return execution, err
	}

	req := c.GetProxy("POST").Suffix(uri).Body(body)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/execute-test-suite returned error: %w", err)
	}

	return c.getTestExecutionFromResponse(resp)
}

func (c ProxyAPIClient) GetTestSuiteExecution(executionID string) (execution testkube.TestSuiteExecution, err error) {
	uri := c.getURI("/test-suite-executions/%s", executionID)
	req := c.GetProxy("GET").Suffix(uri)
	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return execution, fmt.Errorf("api/get-test-suite-execution returned error: %w", err)
	}

	return c.getTestExecutionFromResponse(resp)
}

// WatchTestSuiteExecution watches for changes in channels of test suite executions steps
func (c ProxyAPIClient) WatchTestSuiteExecution(executionID string) (executionCh chan testkube.TestSuiteExecution, err error) {
	executionCh = make(chan testkube.TestSuiteExecution)

	go func() {
		execution, err := c.GetTestSuiteExecution(executionID)
		if err != nil {
			close(executionCh)
			return
		}
		executionCh <- execution
		for range time.NewTicker(time.Second).C {
			execution, err = c.GetTestSuiteExecution(executionID)
			if err != nil {
				close(executionCh)
				return
			}

			if execution.IsCompleted() {
				close(executionCh)
				return
			}

			executionCh <- execution
		}
	}()
	return
}

// ListExecutions list all executions for given test suite
func (c ProxyAPIClient) ListTestSuiteExecutions(testID string, limit int, tags []string) (executions testkube.TestSuiteExecutionsResult, err error) {
	uri := c.getURI("/test-suite-executions")
	req := c.GetProxy("GET").
		Suffix(uri).
		Param("pageSize", fmt.Sprintf("%d", limit))

	if len(tags) > 0 {
		req.Param("tags", strings.Join(tags, ","))
	}

	if testID != "" {
		req.Param("id", testID)
	}

	resp := req.Do(context.Background())

	if err := c.responseError(resp); err != nil {
		return executions, fmt.Errorf("api/list-test-suite-executions returned error: %w", err)
	}

	return c.getTestExecutionsFromResponse(resp)
}

func (c ProxyAPIClient) getTestSuitesFromResponse(resp rest.Result) (testSuites testkube.TestSuites, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return testSuites, err
	}

	err = json.Unmarshal(bytes, &testSuites)

	return testSuites, err
}

func (c ProxyAPIClient) getTestExecutionFromResponse(resp rest.Result) (execution testkube.TestSuiteExecution, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return execution, err
	}

	err = json.Unmarshal(bytes, &execution)

	return execution, err
}

func (c ProxyAPIClient) getTestExecutionsFromResponse(resp rest.Result) (executions testkube.TestSuiteExecutionsResult, err error) {
	bytes, err := resp.Raw()
	if err != nil {
		return executions, err
	}

	err = json.Unmarshal(bytes, &executions)

	return executions, err
}
