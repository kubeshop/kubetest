package tests

import (
	"fmt"
	"os"
	"time"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

const WatchInterval = 2 * time.Second

func NewStartTestCmd() *cobra.Command {
	var (
		name                     string
		watchEnabled             bool
		params                   map[string]string
		downloadArtifactsEnabled bool
		downloadDir              string
	)

	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"run"},
		Short:   "Starts new test",
		Long:    `Starts new test based on Test Custom Resource name, returns results to console`,
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()

			if len(args) == 0 {
				ui.ExitOnError("Invalid arguments", fmt.Errorf("please pass test name to run"))
			}

			testID := args[0]

			client, namespace := GetClient(cmd)
			namespacedName := fmt.Sprintf("%s/%s", namespace, testID)

			execution, err := client.ExecuteTest(testID, namespace, name, params)
			ui.ExitOnError("starting test execution "+namespacedName, err)

			printTestExecutionDetails(execution)

			execution, err = client.GetTestExecution(execution.Id)
			ui.ExitOnError("getting recent execution data id:"+execution.Id, err)

			uiPrintTestStatus(execution)

			uiShellTestCommandBlock(execution.Id)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "execution name, if empty will be autogenerated")
	cmd.Flags().StringToStringVarP(&params, "param", "p", map[string]string{}, "execution envs passed to executor")
	cmd.Flags().BoolVarP(&watchEnabled, "watch", "f", false, "watch for changes after start")
	cmd.Flags().StringVar(&downloadDir, "download-dir", "artifacts", "download dir")
	cmd.Flags().BoolVarP(&downloadArtifactsEnabled, "download-artifacts", "a", false, "downlaod artifacts automatically")

	return cmd
}

func uiPrintTestStatus(execution testkube.TestExecution) {
	switch execution.Status {
	case testkube.TestStatusQueued:
		ui.Warn("Test queued for execution")

	case testkube.TestStatusPending:
		ui.Warn("Test execution started")

	case testkube.TestStatusSuccess:
		duration := execution.EndTime.Sub(execution.StartTime)
		ui.Success("Test execution completed with sucess in " + duration.String())

	case testkube.TestStatusError:
		ui.Errf("Test execution failed")
		os.Exit(1)
	}

	ui.NL()
}

func uiShellTestCommandBlock(id string) {
	ui.ShellCommand(
		"Use following command to get test execution details",
		"kubectl testkube tests execution "+id,
	)

	ui.NL()
}
