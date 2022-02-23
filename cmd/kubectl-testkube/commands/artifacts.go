package commands

import (
	"os"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common/validator"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/tests"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

var (
	executionID string
	filename    string
	destination string
	downloadDir string
)

func NewArtifactsCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "artifacts",
		Short: "Artifacts management commands",
		Args:  validator.ExecutionID,
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

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "should I show additional debug messages")

	cmd.AddCommand(NewListArtifactsCmd())
	cmd.AddCommand(NewDownloadSingleArtifactsCmd())
	cmd.AddCommand(NewDownloadAllArtifactsCmd())

	return cmd
}

func NewListArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <executionID>",
		Short: "List artifacts of the given execution ID",
		Args:  validator.ExecutionID,
		Run: func(cmd *cobra.Command, args []string) {
			executionID = args[0]
			cmd.SilenceUsage = true
			client, _ := common.GetClient(cmd)
			artifacts, err := client.GetExecutionArtifacts(executionID)
			ui.ExitOnError("getting artifacts ", err)

			ui.Table(artifacts, os.Stdout)
		},
	}

	cmd.PersistentFlags().StringVarP(&client, "client", "c", "proxy", "Client used for connecting to testkube API one of proxy|direct")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "should I show additional debug messages")

	cmd.PersistentFlags().StringVarP(&executionID, "execution-id", "e", "", "ID of the execution")

	// output renderer flags
	return cmd
}

func NewDownloadSingleArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download-one <executionID> <fileName> <destinationDir>",
		Short: "download artifact",
		Args:  validator.ExecutionIDAndFileNames,
		Run: func(cmd *cobra.Command, args []string) {
			executionID := args[0]
			filename := args[1]
			destination := args[2]

			client, _ := common.GetClient(cmd)
			f, err := client.DownloadFile(executionID, filename, destination)
			ui.ExitOnError("downloading file"+filename, err)

			ui.Info("File %s downloaded.\n", f)
		},
	}

	cmd.PersistentFlags().StringVarP(&client, "client", "c", "proxy", "Client used for connecting to testkube API one of proxy|direct")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "should I show additional debug messages")

	cmd.PersistentFlags().StringVarP(&executionID, "execution-id", "e", "", "ID of the execution")
	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "name of the file")
	cmd.PersistentFlags().StringVarP(&destination, "destination", "d", "", "name of the file")

	// output renderer flags
	return cmd
}

func NewDownloadAllArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download <executionID>",
		Short: "download artifacts",
		Args:  validator.ExecutionID,
		Run: func(cmd *cobra.Command, args []string) {
			executionID := args[0]
			client, _ := common.GetClient(cmd)
			tests.DownloadArtifacts(executionID, downloadDir, client)
		},
	}

	cmd.PersistentFlags().StringVarP(&client, "client", "c", "proxy", "Client used for connecting to testkube API one of proxy|direct")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "should I show additional debug messages")

	cmd.PersistentFlags().StringVarP(&executionID, "execution-id", "e", "", "ID of the execution")
	cmd.Flags().StringVar(&downloadDir, "download-dir", "artifacts", "download dir")

	// output renderer flags
	return cmd
}
