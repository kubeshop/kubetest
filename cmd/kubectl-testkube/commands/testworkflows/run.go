package testworkflows

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common/render"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/testworkflows/renderer"
	apiclientv1 "github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
)

const (
	LogTimestampLength = 30 // time.RFC3339Nano without 00:00 timezone
)

func NewRunTestWorkflowCmd() *cobra.Command {
	var (
		executionName string
		config        map[string]string
		watchEnabled  bool
	)

	cmd := &cobra.Command{
		Use:     "testworkflow [name]",
		Aliases: []string{"testworkflows", "tw"},
		Args:    cobra.ExactArgs(1),
		Short:   "Starts test workflow execution",

		Run: func(cmd *cobra.Command, args []string) {
			if common.IsBothEnabledAndDisabledSet(cmd) {
				ui.Failf("both --enable-webhooks and --disable-webhooks flags are set, please use only one")
			}

			outputFlag := cmd.Flag("output")
			outputType := render.OutputPretty
			if outputFlag != nil {
				outputType = render.OutputType(outputFlag.Value.String())
			}

			outputPretty := outputType == render.OutputPretty
			namespace := cmd.Flag("namespace").Value.String()
			client, _, err := common.GetClient(cmd)
			ui.ExitOnError("getting client", err)

			var disableWebhooks bool
			if cmd.Flag("enable-webhooks").Changed {
				disableWebhooks = false
			}
			if cmd.Flag("disable-webhooks").Changed {
				disableWebhooks = true
			}

			name := args[0]
			execution, err := client.ExecuteTestWorkflow(name, testkube.TestWorkflowExecutionRequest{
				Name:            executionName,
				Config:          config,
				DisableWebhooks: disableWebhooks,
			})
			ui.ExitOnError("execute test workflow "+name+" from namespace "+namespace, err)
			err = renderer.PrintTestWorkflowExecution(cmd, os.Stdout, execution)
			ui.ExitOnError("render test workflow execution", err)

			var exitCode = 0
			if outputPretty {
				ui.NL()
				if watchEnabled {
					exitCode = uiWatch(execution, client)
					ui.NL()
				} else {
					uiShellWatchExecution(execution.Id)
				}

				uiShellGetExecution(execution.Id)
			}

			os.Exit(exitCode)
		},
	}

	cmd.Flags().StringVarP(&executionName, "name", "n", "", "execution name, if empty will be autogenerated")
	cmd.Flags().StringToStringVarP(&config, "config", "", map[string]string{}, "configuration variables in a form of name1=val1 passed to executor")
	cmd.Flags().BoolVarP(&watchEnabled, "watch", "f", false, "watch for changes after start")
	cmd.Flags().Bool("disable-webhooks", false, "disable webhooks for this execution")
	cmd.Flags().Bool("enable-webhooks", false, "enable webhooks for this execution")

	return cmd
}

func uiWatch(execution testkube.TestWorkflowExecution, client apiclientv1.Client) int {
	result, err := watchTestWorkflowLogs(execution.Id, execution.Signature, client)
	ui.ExitOnError("reading test workflow execution logs", err)

	// Apply the result in the execution
	execution.Result = result
	if result.IsFinished() {
		execution.StatusAt = result.FinishedAt
	}

	// Display message depending on the result
	switch {
	case result.Initialization.ErrorMessage != "":
		ui.Warn("test workflow execution failed:\n")
		ui.Errf(result.Initialization.ErrorMessage)
		return 1
	case result.IsFailed():
		ui.Warn("test workflow execution failed")
		return 1
	case result.IsAborted():
		ui.Warn("test workflow execution aborted")
		return 1
	case result.IsPassed():
		ui.Success("test workflow execution completed with success in " + result.FinishedAt.Sub(result.QueuedAt).String())
	}
	return 0
}

func uiShellGetExecution(id string) {
	ui.ShellCommand(
		"Use following command to get test workflow execution details",
		"kubectl testkube get twe "+id,
	)
}

func uiShellWatchExecution(id string) {
	ui.ShellCommand(
		"Watch test workflow execution until complete",
		"kubectl testkube watch twe "+id,
	)
}

func flattenSignatures(sig []testkube.TestWorkflowSignature) []testkube.TestWorkflowSignature {
	res := make([]testkube.TestWorkflowSignature, 0)
	for _, s := range sig {
		if len(s.Children) == 0 {
			res = append(res, s)
		} else {
			res = append(res, flattenSignatures(s.Children)...)
		}
	}
	return res
}

func printResultDifference(res1 *testkube.TestWorkflowResult, res2 *testkube.TestWorkflowResult, steps []testkube.TestWorkflowSignature) bool {
	if res1 == nil || res2 == nil {
		return false
	}
	changed := false
	for i, s := range steps {
		r1 := res1.Steps[s.Ref]
		r2 := res2.Steps[s.Ref]
		r1Status := testkube.QUEUED_TestWorkflowStepStatus
		r2Status := testkube.QUEUED_TestWorkflowStepStatus
		if r1.Status != nil {
			r1Status = *r1.Status
		}
		if r2.Status != nil {
			r2Status = *r2.Status
		}
		if r1Status == r2Status {
			continue
		}
		name := s.Category
		if s.Name != "" {
			name = s.Name
		}
		took := r2.FinishedAt.Sub(r2.QueuedAt).Round(time.Millisecond)
		changed = true

		switch r2Status {
		case testkube.RUNNING_TestWorkflowStepStatus:
			fmt.Print(ui.LightCyan(fmt.Sprintf("\n• (%d/%d) %s\n", i+1, len(steps), name)))
		case testkube.SKIPPED_TestWorkflowStepStatus:
			fmt.Print(ui.LightGray("• skipped\n"))
		case testkube.PASSED_TestWorkflowStepStatus:
			fmt.Print(ui.Green(fmt.Sprintf("\n• passed in %s\n", took)))
		case testkube.ABORTED_TestWorkflowStepStatus:
			fmt.Print(ui.Red("\n• aborted\n"))
		default:
			if s.Optional {
				fmt.Print(ui.Yellow(fmt.Sprintf("\n• %s in %s (ignored)\n", string(r2Status), took)))
			} else {
				fmt.Print(ui.Red(fmt.Sprintf("\n• %s in %s\n", string(r2Status), took)))
			}
		}
	}

	return changed
}

func watchTestWorkflowLogs(id string, signature []testkube.TestWorkflowSignature, client apiclientv1.Client) (*testkube.TestWorkflowResult, error) {
	ui.Info("Getting logs from test workflow job", id)

	notifications, err := client.GetTestWorkflowExecutionNotifications(id)
	ui.ExitOnError("getting logs from executor", err)

	steps := flattenSignatures(signature)

	var result *testkube.TestWorkflowResult
	var isLineBeginning = true
	for l := range notifications {
		if l.Output != nil {
			continue
		}
		if l.Result != nil {
			isLineBeginning = printResultDifference(result, l.Result, steps)
			result = l.Result
			continue
		}

		// Strip timestamp + space for all new lines in the log
		for len(l.Log) > 0 {
			if isLineBeginning {
				if len(l.Log) >= 29 && l.Log[29] == '+' {
					// Custom timezone (+00:00)
					l.Log = l.Log[len(time.RFC3339Nano)+1:]
				} else {
					// UTC timezone (Z)
					l.Log = l.Log[LogTimestampLength+1:]
				}
				isLineBeginning = false
			}
			newLineIndex := strings.Index(l.Log, "\n")
			if newLineIndex == -1 {
				fmt.Print(l.Log)
				break
			} else {
				fmt.Print(l.Log[0 : newLineIndex+1])
				l.Log = l.Log[newLineIndex+1:]
				isLineBeginning = true
			}
		}
	}

	ui.NL()

	return result, err
}
