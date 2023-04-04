package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubeshop/testkube/pkg/envs"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/executor/env"
	outputPkg "github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/executor/runner"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewRunner(params envs.Params) *K6Runner {
	outputPkg.PrintLogf("%s Preparing test runner", ui.IconTruck)

	return &K6Runner{
		Params: params,
	}
}

type K6Runner struct {
	Params envs.Params
}

var _ runner.Runner = &K6Runner{}

const K6Cloud = "cloud"
const K6Run = "run"

func (r *K6Runner) Run(ctx context.Context, execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	outputPkg.PrintLogf("%s Preparing for test run", ui.IconTruck)

	// check that the datadir exists
	_, err = os.Stat(r.Params.DataDir)
	if errors.Is(err, os.ErrNotExist) {
		outputPkg.PrintLogf("%s Datadir %s does not exist", ui.IconCross, r.Params.DataDir)
		return result, err
	}

	var args []string

	k6TestType := strings.Split(execution.TestType, "/")
	if len(k6TestType) != 2 {
		outputPkg.PrintLogf("%s Invalid test type %s", ui.IconCross, execution.TestType)
		return *result.Err(errors.Errorf("invalid test type %s", execution.TestType)), nil
	}

	k6Subtype := k6TestType[1]
	if k6Subtype == K6Cloud {
		args = append(args, K6Cloud)
	} else {
		args = append(args, K6Run)
	}

	envManager := env.NewManagerWithVars(execution.Variables)
	envManager.GetReferenceVars(envManager.Variables)
	for _, variable := range envManager.Variables {
		if variable.Name != "K6_CLOUD_TOKEN" {
			// pass to k6 using -e option
			envvar := fmt.Sprintf("%s=%s", variable.Name, variable.Value)
			args = append(args, "-e", envvar)
		}
	}

	// convert executor env variables to k6 env variables
	// Deprecated: use Basic Variable instead
	for key, value := range execution.Envs {
		if key != "K6_CLOUD_TOKEN" {
			// pass to k6 using -e option
			envvar := fmt.Sprintf("%s=%s", key, value)
			args = append(args, "-e", envvar)
		}
	}

	// pass additional executor arguments/flags to k6
	args = append(args, execution.Args...)

	var directory string

	// in case of a test file execution we will pass the
	// file path as final parameter to k6
	if execution.Content.Type_ == string(testkube.TestContentTypeString) ||
		execution.Content.Type_ == string(testkube.TestContentTypeFileURI) {
		directory = r.Params.DataDir
		args = append(args, "test-content")
	}

	// in case of Git directory we will run k6 here and
	// use the last argument as test file
	if execution.Content.Type_ == string(testkube.TestContentTypeGitFile) ||
		execution.Content.Type_ == string(testkube.TestContentTypeGitDir) ||
		execution.Content.Type_ == string(testkube.TestContentTypeGit) {
		directory = filepath.Join(r.Params.DataDir, "repo")
		path := ""
		workingDir := ""
		if execution.Content != nil && execution.Content.Repository != nil {
			path = execution.Content.Repository.Path
			workingDir = execution.Content.Repository.WorkingDir
		}

		fileInfo, err := os.Stat(filepath.Join(directory, path))
		if err != nil {
			outputPkg.PrintLogf("%s k6 test directory %v not found", ui.IconCross, err)
			return *result.Err(errors.Errorf("k6 test directory %v not found", err)), nil
		}

		if fileInfo.IsDir() {
			args[len(args)-1] = filepath.Join(path, args[len(args)-1])
		} else {
			args = append(args, path)
		}

		// sanity checking for test script
		scriptFile := filepath.Join(directory, workingDir, args[len(args)-1])
		fileInfo, err = os.Stat(scriptFile)
		if errors.Is(err, os.ErrNotExist) || fileInfo.IsDir() {
			outputPkg.PrintLogf("%s k6 test script %s not found", ui.IconCross, scriptFile)
			return *result.Err(errors.Errorf("k6 test script %s not found", scriptFile)), nil
		}
	}

	outputPkg.PrintEvent("Running", directory, "k6", args)
	runPath := directory
	if execution.Content.Repository != nil && execution.Content.Repository.WorkingDir != "" {
		runPath = filepath.Join(directory, execution.Content.Repository.WorkingDir)
	}

	output, err := executor.Run(runPath, "k6", envManager, args...)
	output = envManager.ObfuscateSecrets(output)
	return finalExecutionResult(string(output), err), nil
}

// finalExecutionResult processes the output of the test run
func finalExecutionResult(output string, err error) (result testkube.ExecutionResult) {
	succeeded := isSuccessful(output)
	switch {
	case err == nil && succeeded:
		outputPkg.PrintLogf("%s Test run successful", ui.IconCheckMark)
		result.Status = testkube.ExecutionStatusPassed
	case err == nil && !succeeded:
		outputPkg.PrintLogf("%s Test run failed: some checks have failed", ui.IconCross)
		result.Status = testkube.ExecutionStatusFailed
		result.ErrorMessage = "some checks have failed"
	case err != nil && strings.Contains(err.Error(), "exit status 99"):
		// tests have run, but some checks + thresholds have failed
		outputPkg.PrintLogf("%s Test run failed: some thresholds have failed: %s", ui.IconCross, err.Error())
		result.Status = testkube.ExecutionStatusFailed
		result.ErrorMessage = "some thresholds have failed"
	default:
		// k6 was unable to run at all
		outputPkg.PrintLogf("%s Test run failed: %s", ui.IconCross, err.Error())
		result.Status = testkube.ExecutionStatusFailed
		result.ErrorMessage = err.Error()
		return result
	}

	// always set these, no matter if error or success
	result.Output = output
	result.OutputType = "text/plain"

	result.Steps = []testkube.ExecutionStepResult{}
	for _, name := range parseScenarioNames(output) {
		result.Steps = append(result.Steps, testkube.ExecutionStepResult{
			// use the scenario name with description here
			Name:     name,
			Duration: parseScenarioDuration(output, splitScenarioName(name)),

			// currently there is no way to extract individual scenario status
			Status: string(testkube.PASSED_ExecutionStatus),
		})
	}

	return result
}

// isSuccessful checks the output of the k6 test to make sure nothing fails
func isSuccessful(summary string) bool {
	return areChecksSuccessful(summary) && !containsErrors(summary)
}

// areChecksSuccessful verifies the summary at the end of the execution to see
// if any of the checks failed
func areChecksSuccessful(summary string) bool {
	lines := splitSummaryBody(summary)
	for _, line := range lines {
		if !strings.Contains(line, "checks") {
			continue
		}
		if strings.Contains(line, "100.00%") {
			return true
		}
		return false
	}

	return true
}

// containsErrors checks for error level messages.
// As discussed in this GitHub issue: https://github.com/grafana/k6/issues/1680,
// k6 summary does not include tests failing because an error was encountered.
// To make sure no errors happened, we check the output for error level messages
func containsErrors(summary string) bool {
	return strings.Contains(summary, "level=error")
}

func parseScenarioNames(summary string) []string {
	lines := splitSummaryBody(summary)
	var names []string

	for _, line := range lines {
		if strings.Contains(line, "* ") {
			name := strings.TrimLeft(strings.TrimSpace(line), "* ")
			names = append(names, name)
		}
	}

	return names
}

func parseScenarioDuration(summary string, name string) string {
	lines := splitSummaryBody(summary)

	var duration string
	for _, line := range lines {
		if strings.Contains(line, name) && strings.Contains(line, "[ 100% ]") {
			index := strings.Index(line, "]") + 1
			line = strings.TrimSpace(line[index:])
			line = strings.ReplaceAll(line, "  ", " ")

			// take next line and trim leading spaces
			metrics := strings.Split(line, " ")
			duration = metrics[2]
			break
		}
	}

	return duration
}

func splitScenarioName(name string) string {
	return strings.Split(name, ":")[0]
}

func splitSummaryBody(summary string) []string {
	return strings.Split(summary, "\n")
}

// GetType returns runner type
func (r *K6Runner) GetType() runner.Type {
	return runner.TypeMain
}
