package testsuites

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	apiClient "github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/crd"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewCreateTestSuitesCmd() *cobra.Command {

	var (
		name            string
		file            string
		labels          map[string]string
		variables       map[string]string
		secretVariables map[string]string
		schedule        string
	)

	cmd := &cobra.Command{
		Use:     "testsuite",
		Aliases: []string{"testsuites", "ts"},
		Short:   "Create new TestSuite",
		Long:    `Create new TestSuite Custom Resource`,
		Run: func(cmd *cobra.Command, args []string) {
			var content []byte
			crdOnly, err := strconv.ParseBool(cmd.Flag("crd-only").Value.String())
			ui.ExitOnError("parsing flag value", err)

			if file != "" {
				// read test content
				content, err = os.ReadFile(file)
				ui.ExitOnError("reading file"+file, err)
			} else if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
				content, err = io.ReadAll(os.Stdin)
				ui.ExitOnError("reading stdin", err)
			}

			if name == "" {
				ui.Failf("pass valid test suite name (in '--name' flag)")
			}

			var options testkube.TestSuiteUpsertRequest

			err = json.Unmarshal(content, &options)
			ui.ExitOnError("Invalid file content", err)

			if name != "" {
				options.Name = name
			}

			namespace := cmd.Flag("namespace").Value.String()
			var client apiClient.Client
			if !crdOnly {
				client, namespace = common.GetClient(cmd)
				test, _ := client.GetTestSuite(options.Name)
				if options.Name == test.Name {
					ui.Failf("TestSuite with name '%s' already exists in namespace %s", options.Name, namespace)
				}
			}

			options.Namespace = namespace
			options.Labels = labels

			variables, err := common.CreateVariables(cmd)
			ui.ExitOnError("Invalid variables", err)
			options.Variables = variables

			options.Schedule = cmd.Flag("schedule").Value.String()
			err = validateSchedule(options.Schedule)
			ui.ExitOnError("validating schedule", err)

			if !crdOnly {
				_, err = client.CreateTestSuite((apiClient.UpsertTestSuiteOptions(options)))
				ui.ExitOnError("creating test suite "+options.Name+" in namespace "+options.Namespace, err)

				ui.Success("Test suite created", options.Name)
			} else {
				if options.Description != "" {
					options.Description = fmt.Sprintf("%q", options.Description)
				}

				data, err := crd.ExecuteTemplate(crd.TemplateTestSuite, options)
				ui.ExitOnError("executing crd template", err)

				ui.Info(data)
			}
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "JSON test suite file - will be read from stdin if not specified, look at testkube.TestUpsertRequest")
	cmd.Flags().StringVar(&name, "name", "", "Set/Override test suite name")
	cmd.Flags().StringToStringVarP(&labels, "label", "l", nil, "label key value pair: --label key1=value1")
	cmd.Flags().StringToStringVarP(&variables, "variable", "v", nil, "param key value pair: --variable key1=value1")
	cmd.Flags().StringToStringVarP(&secretVariables, "secret-variable", "s", nil, "secret variable key value pair: --secret-variable key1=value1")
	cmd.Flags().StringVarP(&schedule, "schedule", "", "", "test suite schedule in a cronjob form: * * * * *")

	return cmd
}

func validateSchedule(schedule string) error {
	if schedule == "" {
		return nil
	}

	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	if _, err := specParser.Parse(schedule); err != nil {
		return err
	}

	return nil
}
