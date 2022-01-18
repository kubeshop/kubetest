package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	apiClient "github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewCreateTestsCmd() *cobra.Command {

	var (
		name string
		file string
		tags []string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new test",
		Long:  `Create new Test Custom Resource, `,
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()

			var content []byte
			var err error

			if file != "" {
				// read test content
				content, err = ioutil.ReadFile(file)
				ui.ExitOnError("reading file"+file, err)
			} else if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
				content, err = ioutil.ReadAll(os.Stdin)
				ui.ExitOnError("reading stdin", err)
			}

			var options testkube.TestUpsertRequest

			err = json.Unmarshal(content, &options)
			ui.ExitOnError("Invalid file content", err)

			if name != "" {
				options.Name = name
			}

			client, _ := common.GetClient(cmd)

			test, _ := client.GetTest(options.Name, options.Namespace)
			if options.Name == test.Name {
				ui.Failf("Test with name '%s' already exists in namespace %s", options.Name, options.Namespace)
			}

			options.Tags = tags

			test, err = client.CreateTest((apiClient.UpsertTestOptions(options)))
			ui.ExitOnError("creating test "+options.Name+" in namespace "+options.Namespace, err)
			ui.Success("Test created", options.Name)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "JSON test file - will be read from stdin if not specified, look at testkube.TestUpsertRequest")
	cmd.Flags().StringVar(&name, "name", "", "Set/Override test name")
	cmd.Flags().StringSliceVar(&tags, "tags", nil, "comma separated list of tags: --tags tag1,tag2,tag3")

	return cmd
}
