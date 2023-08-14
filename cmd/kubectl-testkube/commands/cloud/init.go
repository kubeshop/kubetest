package cloud

import (
	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/config"
	"github.com/kubeshop/testkube/pkg/telemetry"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewInitCmd() *cobra.Command {
	options := common.HelmOptions{
		NoMinio:     true,
		NoMongo:     true,
		NoDashboard: true,
	}

	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Install Helm chart registry in current kubectl context and update dependencies",
		Aliases: []string{"install"},
		Run: func(cmd *cobra.Command, args []string) {
			ui.Info("WELCOME TO")
			ui.Logo()

			cfg, err := config.Load()
			ui.ExitOnError("loading config file", err)
			ui.NL()
			sendAttemptTelemetry(cmd, cfg)

			// create new cloud uris
			options.CloudUris = common.NewCloudUris(options.CloudRootDomain)
			if !options.NoConfirm {
				ui.Warn("This will install Testkube to the latest version. This may take a few minutes.")
				ui.Warn("Please be sure you're on valid kubectl context before continuing!")
				ui.NL()

				currentContext, err := common.GetCurrentKubernetesContext()
				sendErrTelemetry(cmd, cfg, "k8s_context")
				ui.ExitOnError("getting current context", err)
				ui.Alert("Current kubectl context:", currentContext)
				ui.NL()

				ok := ui.Confirm("Do you want to continue?")
				if !ok {
					ui.Errf("Testkube installation cancelled")
					sendErrTelemetry(cmd, cfg, "user_cancel")
					return
				}
			}

			spinner := ui.NewSpinner("Installing Testkube")
			err = common.HelmUpgradeOrInstallTestkubeCloud(options, cfg, false)
			sendErrTelemetry(cmd, cfg, "helm_install")
			ui.ExitOnError("Installing Testkube", err)
			spinner.Success()

			ui.NL()

			ui.H2("Saving testkube cli cloud context")
			var token, refreshToken string
			if !common.IsUserLoggedIn(cfg, options) {
				token, refreshToken, err = common.LoginUser(options.CloudUris.Auth)
				sendErrTelemetry(cmd, cfg, "login")
				ui.ExitOnError("user login", err)
			}
			err = common.PopulateLoginDataToContext(options.CloudOrgId, options.CloudEnvId, token, refreshToken, options, cfg)
			sendErrTelemetry(cmd, cfg, "setting_context")
			ui.ExitOnError("Setting cloud environment context", err)

			ui.Info(" Happy Testing! 🚀")
			ui.NL()
		},
	}

	cmd.Flags().StringVar(&options.Chart, "chart", "kubeshop/testkube", "chart name (usually you don't need to change it)")
	cmd.Flags().StringVar(&options.Name, "name", "testkube", "installation name (usually you don't need to change it)")
	cmd.Flags().StringVar(&options.Namespace, "namespace", "testkube", "namespace where to install")
	cmd.Flags().StringVar(&options.Values, "values", "", "path to Helm values file")

	cmd.Flags().StringVar(&options.CloudAgentToken, "agent-token", "", "Testkube Cloud agent key")
	cmd.Flags().StringVar(&options.CloudOrgId, "org-id", "", "Testkube Cloud organization id")
	cmd.Flags().StringVar(&options.CloudEnvId, "env-id", "", "Testkube Cloud environment id")

	cmd.Flags().StringVar(&options.CloudRootDomain, "cloud-root-domain", "testkube.io", "defaults to testkube.io, usually don't need to be changed [required for cloud mode]")

	cmd.Flags().BoolVar(&options.NoConfirm, "no-confirm", false, "don't ask for confirmation - unatended installation mode")
	cmd.Flags().BoolVar(&options.DryRun, "dry-run", false, "dry run mode - only print commands that would be executed")

	return cmd
}

func sendErrTelemetry(cmd *cobra.Command, clientCfg config.Data, errType string) {
	if clientCfg.TelemetryEnabled {
		ui.Debug("collecting anonymous telemetry data, you can disable it by calling `kubectl testkube disable telemetry`")
		out, err := telemetry.SendCmdErrorEvent(cmd, common.Version, errType)
		if ui.Verbose && err != nil {
			ui.Err(err)
		}
		ui.Debug("telemetry send event response", out)
	}
}

func sendAttemptTelemetry(cmd *cobra.Command, clientCfg config.Data) {
	if clientCfg.TelemetryEnabled {
		ui.Debug("collecting anonymous telemetry data, you can disable it by calling `kubectl testkube disable telemetry`")
		out, err := telemetry.SendCmdAttemptEvent(cmd, common.Version)
		if ui.Verbose && err != nil {
			ui.Err(err)
		}
		ui.Debug("telemetry send event response", out)
	}
}
