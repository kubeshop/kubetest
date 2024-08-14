package testworkflows

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common/render"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/testworkflows/renderer"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewGetTestWorkflowExecutionsCmd() *cobra.Command {
	var (
		limit            int
		selectors        []string
		testWorkflowName string
		logsOnly         bool
	)

	cmd := &cobra.Command{
		Use:     "testworkflowexecution [executionID]",
		Aliases: []string{"testworkflowexecutions", "twe", "tw-execution", "twexecution"},
		Args:    cobra.MaximumNArgs(1),
		Short:   "Gets TestWorkflow execution details",
		Long:    `Gets TestWorkflow execution details by ID, or list if id is not passed`,

		Run: func(cmd *cobra.Command, args []string) {
			outputFlag := cmd.Flag("output")
			outputType := render.OutputPretty
			if outputFlag != nil {
				outputType = render.OutputType(outputFlag.Value.String())
			}

			outputPretty := outputType == render.OutputPretty

			client, _, err := common.GetClient(cmd)
			ui.ExitOnError("getting client", err)

			if len(args) == 0 {
				client, _, err := common.GetClient(cmd)
				ui.ExitOnError("getting client", err)

				executions, err := client.ListTestWorkflowExecutions(testWorkflowName, limit, strings.Join(selectors, ","))
				ui.ExitOnError("getting test workflow executions list", err)
				err = render.List(cmd, testkube.TestWorkflowExecutionSummaries(executions.Results), os.Stdout)
				ui.ExitOnError("rendering list", err)
				return
			}

			executionID := args[0]
			execution, err := client.GetTestWorkflowExecution(executionID)
			ui.ExitOnError("getting recent test workflow execution data id:"+execution.Id, err)
			if !logsOnly {
				err = render.Obj(cmd, execution, os.Stdout, renderer.TestWorkflowExecutionRenderer)
				ui.ExitOnError("rendering obj", err)
			}

			if outputPretty {
				ui.Info("Getting logs for test workflow execution", executionID)

				logs, err := client.GetTestWorkflowExecutionLogs(executionID)
				ui.ExitOnError("getting logs from executor", err)

				sigs := flattenSignatures(execution.Signature)

				var results map[string]testkube.TestWorkflowStepResult
				if execution.Result != nil {
					results = execution.Result.Steps
				}

				printRawLogLines(logs, sigs, results)
				if !logsOnly {
					render.PrintTestWorkflowExecutionURIs(&execution)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&testWorkflowName, "testworkflow", "w", "", "test workflow name")
	cmd.Flags().IntVar(&limit, "limit", 1000, "max number of records to return")
	cmd.Flags().StringSliceVarP(&selectors, "label", "l", nil, "label key value pair: --label key1=value1")
	cmd.Flags().BoolVar(&logsOnly, "logs-only", false, "show only execution logs")

	return cmd
}
