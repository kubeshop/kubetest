package render

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v2"

	"github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/kubeshop/testkube/pkg/utils"
)

type OutputType string

const (
	OutputGoTemplate OutputType = "go"
	OutputJSON       OutputType = "json"
	OutputYAML       OutputType = "yaml"
	OutputPretty     OutputType = "pretty"
)

type CliObjRenderer func(client client.Client, ui *ui.UI, obj interface{}) error

func RenderJSON(obj interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(obj)
}

func RenderYaml(obj interface{}, w io.Writer) error {
	return yaml.NewEncoder(w).Encode(obj)
}

func RenderGoTemplate(item interface{}, w io.Writer, tpl string) error {
	tmpl, err := utils.NewTemplate("result").Parse(tpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, item)
}

func RenderGoTemplateList(list []interface{}, w io.Writer, tpl string) error {
	tmpl, err := utils.NewTemplate("result").Parse(tpl)
	if err != nil {
		return err
	}

	for _, item := range list {
		err := tmpl.Execute(w, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func RenderPrettyList(obj ui.TableData, w io.Writer) error {
	ui.NL()
	ui.Table(obj, w)
	ui.NL()
	return nil
}

func RenderExecutionResult(client client.Client, execution *testkube.Execution, logsOnly bool, showLogs bool) error {
	result := execution.ExecutionResult
	if result == nil {
		ui.Errf("got execution without `Result`")
		return nil
	}

	ui.NL()
	switch true {
	case result.IsQueued():
		ui.Warn("Test queued for execution")

	case result.IsRunning():
		ui.Warn("Test execution started")

	case result.IsPassed():
		if showLogs {
			ui.Info(result.Output)
		}

		if !logsOnly {
			duration := execution.EndTime.Sub(execution.StartTime)
			ui.Success("Test execution completed with success in " + duration.String())

			info, err := client.GetServerInfo()
			ui.ExitOnError("getting server info", err)

			PrintExecutionURIs(execution, info.DashboardUri)
		}

	case result.IsAborted():
		ui.Warn("Test execution aborted")

	case result.IsTimeout():
		ui.Warn("Test execution timeout")

	case result.IsFailed():
		if logsOnly {
			ui.Info(result.ErrorMessage)
		} else {
			ui.UseStderr()
			ui.Warn("Test execution failed:\n")
			ui.Errf(result.ErrorMessage)

			info, err := client.GetServerInfo()
			ui.ExitOnError("getting server info", err)

			PrintExecutionURIs(execution, info.DashboardUri)
		}

		if showLogs {
			ui.Info(result.Output)
		}
		return errors.New(result.ErrorMessage)

	default:
		if logsOnly {
			ui.Info(result.ErrorMessage)
		} else {
			ui.UseStderr()
			ui.Warn("Test execution status unknown:\n")
			ui.Errf(result.ErrorMessage)
		}

		if showLogs {
			ui.Info(result.Output)
		}
		return errors.New(result.ErrorMessage)
	}

	return nil
}

func PrintExecutionURIs(execution *testkube.Execution, dashboardURI string) {
	ui.NL()
	ui.Link("Test URI:", fmt.Sprintf("%s/tests/%s", dashboardURI, execution.TestName))
	ui.Link("Test Execution URI:", fmt.Sprintf("%s/tests/%s/executions/%s", dashboardURI,
		execution.TestName, execution.Id))
	ui.NL()
}

func PrintTestSuiteExecutionURIs(execution *testkube.TestSuiteExecution, dashboardURI string) {
	ui.NL()
	testSuiteName := ""
	if execution.TestSuite != nil {
		testSuiteName = execution.TestSuite.Name
	}

	ui.Link("Test Suite URI:", fmt.Sprintf("%s/test-suites/%s", dashboardURI, testSuiteName))
	ui.Link("Test Suite Execution URI:", fmt.Sprintf("%s/test-suites/%s/executions/%s", dashboardURI,
		testSuiteName, execution.Id))
	ui.NL()
}
