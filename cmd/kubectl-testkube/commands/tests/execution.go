package tests

import (
	"fmt"

	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewTestExecutionCmd() *cobra.Command {
	var (
		name                     string
		watchEnabled             bool
		params                   map[string]string
		downloadArtifactsEnabled bool
		downloadDir              string
	)

	cmd := &cobra.Command{
		Use:     "execution",
		Aliases: []string{"run"},
		Short:   "Starts new test",
		Long:    `Starts new test based on Test Custom Resource name, returns results to console`,
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()

			if len(args) == 0 {
				ui.ExitOnError("Invalid arguments", fmt.Errorf("please pass execution ID"))
			}

			client, _ := GetClient(cmd)

			executionID := args[0]
			execution, err := client.GetTestExecution(executionID)
			ui.ExitOnError("getting recent execution data id:"+execution.Id, err)

			printTestExecutionDetails(execution)

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
