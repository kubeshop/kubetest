package tests

import (
	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewUpdateTestsCmd() *cobra.Command {

	var (
		testName                 string
		testContentType          string
		file                     string
		executorType             string
		uri                      string
		gitUri                   string
		gitBranch                string
		gitCommit                string
		gitPath                  string
		gitUsername              string
		gitToken                 string
		sourceName               string
		labels                   map[string]string
		variables                map[string]string
		secretVariables          map[string]string
		schedule                 string
		executorArgs             []string
		argsMode                 string
		executionName            string
		variablesFile            string
		envs                     map[string]string
		secretEnvs               map[string]string
		httpProxy, httpsProxy    string
		gitUsernameSecret        map[string]string
		gitTokenSecret           map[string]string
		secretVariableReferences map[string]string
		copyFiles                []string
		image                    string
		command                  []string
		imagePullSecretNames     []string
		timeout                  int64
		gitWorkingDir            string
		gitCertificateSecret     string
		gitAuthType              string
		artifactStorageClassName string
		artifactVolumeMountPath  string
		artifactDirs             []string
		jobTemplate              string
		cronJobTemplate          string
		preRunScript             string
		postRunScript            string
		scraperTemplate          string
		negativeTest             bool
		mountConfigMaps          map[string]string
		variableConfigMaps       []string
		mountSecrets             map[string]string
		variableSecrets          []string
		description              string
	)

	cmd := &cobra.Command{
		Use:   "test",
		Short: "Update test",
		Long:  `Update Test Custom Resource`,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if testName == "" {
				ui.Failf("pass valid test name (in '--name' flag)")
			}

			client, namespace, err := common.GetClient(cmd)
			ui.ExitOnError("getting client", err)

			test, _ := client.GetTest(testName)
			if testName != test.Name {
				ui.Failf("Test with name '%s' not exists in namespace %s", testName, namespace)
			}

			options, err := NewUpdateTestOptionsFromFlags(cmd)
			ui.ExitOnError("getting test options", err)

			test, err = client.UpdateTest(options)
			ui.ExitOnError("updating test "+testName+" in namespace "+namespace, err)

			ui.Success("Test updated", namespace, "/", testName)
		},
	}

	cmd.Flags().StringVarP(&testName, "name", "n", "", "unique test name - mandatory")
	cmd.Flags().StringVarP(&file, "file", "f", "", "test file - will try to read content from stdin if not specified")
	cmd.Flags().StringVarP(&testContentType, "test-content-type", "", "", "content type of test one of string|file-uri|git")

	cmd.Flags().StringVarP(&executorType, "type", "t", "", "test type (defaults to postman-collection)")

	cmd.Flags().StringVarP(&uri, "uri", "", "", "URI of resource - will be loaded by http GET")
	cmd.Flags().StringVarP(&gitUri, "git-uri", "", "", "Git repository uri")
	cmd.Flags().StringVarP(&gitBranch, "git-branch", "", "", "if uri is git repository we can set additional branch parameter")
	cmd.Flags().StringVarP(&gitCommit, "git-commit", "", "", "if uri is git repository we can use commit id (sha) parameter")
	cmd.Flags().StringVarP(&gitPath, "git-path", "", "", "if repository is big we need to define additional path to directory/file to checkout partially")
	cmd.Flags().StringVarP(&gitUsername, "git-username", "", "", "if git repository is private we can use username as an auth parameter")
	cmd.Flags().StringVarP(&gitToken, "git-token", "", "", "if git repository is private we can use token as an auth parameter")
	cmd.Flags().StringVarP(&sourceName, "source", "", "", "source name - will be used together with content parameters")
	cmd.Flags().StringToStringVarP(&labels, "label", "l", nil, "label key value pair: --label key1=value1")
	cmd.Flags().StringToStringVarP(&variables, "variable", "v", nil, "variable key value pair: -v key1=value1")
	cmd.Flags().StringToStringVarP(&secretVariables, "secret-variable", "s", nil, "secret variable key value pair: -s key1=value1")
	cmd.Flags().StringVarP(&schedule, "schedule", "", "", "test schedule in a cron job form: * * * * *")
	cmd.Flags().StringArrayVarP(&command, "command", "", []string{}, "command passed to image in executor")
	cmd.Flags().StringArrayVarP(&executorArgs, "executor-args", "", []string{}, "executor binary additional arguments")
	cmd.Flags().StringVarP(&argsMode, "args-mode", "", "append", "usage mode for arguments. one of append|override")
	cmd.Flags().StringVarP(&executionName, "execution-name", "", "", "execution name, if empty will be autogenerated")
	cmd.Flags().StringVarP(&variablesFile, "variables-file", "", "", "variables file path, e.g. postman env file - will be passed to executor if supported")
	cmd.Flags().StringToStringVarP(&envs, "env", "", map[string]string{}, "envs in a form of name1=val1 passed to executor")
	cmd.Flags().StringToStringVarP(&secretEnvs, "secret-env", "", map[string]string{}, "secret envs in a form of secret_key1=secret_name1 passed to executor")
	cmd.Flags().StringVar(&httpProxy, "http-proxy", "", "http proxy for executor containers")
	cmd.Flags().StringVar(&httpsProxy, "https-proxy", "", "https proxy for executor containers")
	cmd.Flags().StringToStringVarP(&gitUsernameSecret, "git-username-secret", "", map[string]string{}, "git username secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringToStringVarP(&gitTokenSecret, "git-token-secret", "", map[string]string{}, "git token secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringToStringVarP(&secretVariableReferences, "secret-variable-reference", "", nil, "secret variable references in a form name1=secret_name1=secret_key1")
	cmd.Flags().StringArrayVarP(&copyFiles, "copy-files", "", []string{}, "file path mappings from host to pod of form source:destination")
	cmd.Flags().StringVarP(&image, "image", "i", "", "image for container executor")
	cmd.Flags().StringArrayVar(&imagePullSecretNames, "image-pull-secrets", []string{}, "secret name used to pull the image in container executor")
	cmd.Flags().Int64Var(&timeout, "timeout", 0, "duration in seconds for test to timeout. 0 disables timeout.")
	cmd.Flags().StringVarP(&gitWorkingDir, "git-working-dir", "", "", "if repository contains multiple directories with tests (like monorepo) and one starting directory we can set working directory parameter")
	cmd.Flags().StringVarP(&gitCertificateSecret, "git-certificate-secret", "", "", "if git repository is private we can use certificate as an auth parameter stored in a kubernetes secret name")
	cmd.Flags().StringVarP(&gitAuthType, "git-auth-type", "", "basic", "auth type for git requests one of basic|header")
	cmd.Flags().StringVar(&artifactStorageClassName, "artifact-storage-class-name", "", "artifact storage class name for container executor")
	cmd.Flags().StringVar(&artifactVolumeMountPath, "artifact-volume-mount-path", "", "artifact volume mount path for container executor")
	cmd.Flags().StringArrayVarP(&artifactDirs, "artifact-dir", "", []string{}, "artifact dirs for scraping")
	cmd.Flags().StringVar(&jobTemplate, "job-template", "", "job template file path for extensions to job template")
	cmd.Flags().StringVar(&cronJobTemplate, "cronjob-template", "", "cron job template file path for extensions to cron job template")
	cmd.Flags().StringVarP(&preRunScript, "prerun-script", "", "", "path to script to be run before test execution")
	cmd.Flags().StringVarP(&postRunScript, "postrun-script", "", "", "path to script to be run after test execution")
	cmd.Flags().StringVar(&scraperTemplate, "scraper-template", "", "scraper template file path for extensions to scraper template")
	cmd.Flags().BoolVar(&negativeTest, "negative-test", false, "negative test, if enabled, makes failure an expected and correct test result. If the test fails the result will be set to success, and vice versa")
	cmd.Flags().StringToStringVarP(&mountConfigMaps, "mount-configmap", "", map[string]string{}, "config map value pair for mounting it to executor pod: --mount-configmap configmap_name=configmap_mountpath")
	cmd.Flags().StringArrayVar(&variableConfigMaps, "variable-configmap", []string{}, "config map name used to map all keys to basis variables")
	cmd.Flags().StringToStringVarP(&mountSecrets, "mount-secret", "", map[string]string{}, "secret value pair for mounting it to executor pod: --mount-secret secret_name=secret_mountpath")
	cmd.Flags().StringArrayVar(&variableSecrets, "variable-secret", []string{}, "secret name used to map all keys to secret variables")
	cmd.Flags().StringVarP(&description, "description", "", "", "test description")
	cmd.Flags().MarkDeprecated("env", "env is deprecated use variable instead")
	cmd.Flags().MarkDeprecated("secret-env", "secret-env is deprecated use secret-variable instead")

	return cmd
}
