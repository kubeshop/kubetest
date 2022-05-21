package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubeshop/testkube/pkg/executor/output"
	phttp "github.com/kubeshop/testkube/pkg/http"
	"github.com/kubeshop/testkube/pkg/problem"
	"golang.org/x/oauth2"
)

type transport struct {
	headers map[string]string
	base    http.RoundTripper
}

// RoundTrip is a method to adjust http request
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}

	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}

	return base.RoundTrip(req)
}

// GetHTTPClient prepares http client
func GetHTTTPClient(token *oauth2.Token) (*http.Client, error) {
	httpClient := phttp.NewClient()
	if token != nil {
		data, err := json.Marshal(token)
		if err != nil {
			return nil, err
		}

		httpClient.Transport = &transport{headers: map[string]string{"Authorization": "Bearer " + base64.StdEncoding.EncodeToString(data)}}
	}

	return httpClient, nil
}

// NewDirectClient returns new direct client
func NewDirectClient[A All](httpClient *http.Client, apiURI string) DirectClient[A] {
	return DirectClient[A]{
		client: httpClient,
		apiURI: apiURI,
	}
}

// DirectClient implements direct client
type DirectClient[A All] struct {
	client *http.Client
	apiURI string
}

// baseExecute is base execute method
func (t DirectClient[A]) baseExec(method, uri, resource string, body []byte, params map[string]string) (resp *http.Response, err error) {
	var buffer io.Reader
	if body != nil {
		buffer = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, uri, buffer)
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	for key, value := range params {
		if value != "" {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err = t.client.Do(req)
	if err != nil {
		return resp, err
	}

	if err = t.responseError(resp); err != nil {
		return resp, fmt.Errorf("api/%s-%s returned error: %w", method, resource, err)
	}

	return resp, nil
}

// Execute is a method to make an api call for a single object
func (t DirectClient[A]) Execute(method, uri string, body []byte, params map[string]string) (result A, err error) {
	resp, err := t.baseExec(method, uri, fmt.Sprintf("%T", result), body, params)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	return t.getFromResponse(resp)
}

// ExecuteMultiple is a method to make an api call for multiple objects
func (t DirectClient[A]) ExecuteMultiple(method, uri string, body []byte, params map[string]string) (result []A, err error) {
	resp, err := t.baseExec(method, uri, fmt.Sprintf("%T", result), body, params)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	return t.getFromResponses(resp)
}

// Delete is a method to make delete api call
func (t DirectClient[A]) Delete(uri, selector string, isContentExpected bool) error {
	resp, err := t.baseExec(http.MethodDelete, uri, uri, nil, map[string]string{"selector": selector})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if isContentExpected && resp.StatusCode != http.StatusNoContent {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("request returned error: %s", respBody)
	}

	return nil
}

// GetURI returns uri for api method
func (t DirectClient[A]) GetURI(pathTemplate string, params ...interface{}) string {
	path := fmt.Sprintf(pathTemplate, params...)
	return fmt.Sprintf("%s/%s%s", t.apiURI, Version, path)
}

// GetLogs returns logs stream from job pods, based on job pods logs
func (t DirectClient[A]) GetLogs(uri string, logs chan output.Output) error {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "text/event-stream")
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}

	go func() {
		defer close(logs)
		defer resp.Body.Close()

		StreamToLogsChannel(resp.Body, logs)
	}()

	return nil
}

// GetFile returns file artifact
func (t DirectClient[A]) GetFile(uri, fileName, destination string) (name string, err error) {
	resp, err := t.client.Get(uri)
	if err != nil {
		return name, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return name, fmt.Errorf("error: %d", resp.StatusCode)
	}

	split := strings.Split(fileName, "/")
	f, err := os.Create(filepath.Join(destination, split[len(split)-1]))
	if err != nil {
		return name, err
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		return name, err
	}

	if err = t.responseError(resp); err != nil {
		return name, fmt.Errorf("api/download-file returned error: %w", err)
	}

	return f.Name(), nil
}

func (t DirectClient[A]) getFromResponse(resp *http.Response) (result A, err error) {
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (t DirectClient[A]) getFromResponses(resp *http.Response) (result []A, err error) {
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

// responseError tries to lookup if response is of Problem type
func (t DirectClient[A]) responseError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		var pr problem.Problem

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("can't get problem from api response: can't read response body %w", err)
		}

		err = json.Unmarshal(bytes, &pr)
		if err != nil {
			return fmt.Errorf("can't get problem from api response: %w, output: %s", err, string(bytes))
		}

		return fmt.Errorf("problem: %+v", pr.Detail)
	}

	return nil
}
