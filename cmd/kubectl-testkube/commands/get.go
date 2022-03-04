package commands

import (
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/artifacts"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/executors"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/tests"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/testsuites"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/webhooks"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get <resourceName>",
		Aliases: []string{"g"},
		Short:   "Get resources",
		Long:    `Get available resources, get single item or list`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// version validation
			// if client version is less than server version show warning
			client, _ := common.GetClient(cmd)

			err := ValidateVersions(client)
			if err != nil {
				ui.Warn(err.Error())
			}
		},
	}

	cmd.AddCommand(tests.NewGetTestsCmd())
	cmd.AddCommand(testsuites.NewGetTestSuiteCmd())
	cmd.AddCommand(webhooks.NewGetWebhookCmd())
	cmd.AddCommand(executors.NewGetExecutorCmd())
	cmd.AddCommand(tests.NewGetExecutionCmd())
	cmd.AddCommand(testsuites.NewTestSuiteExecutionsCmd())
	cmd.AddCommand(artifacts.NewListArtifactsCmd())

	return cmd
}
