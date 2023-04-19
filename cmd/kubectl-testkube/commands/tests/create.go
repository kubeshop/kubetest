package tests

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/pkg/api/v1/client"
	apiv1 "github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/crd"
	"github.com/kubeshop/testkube/pkg/ui"
)

// CreateCommonFlags are common flags for creating all test types
type CreateCommonFlags struct {
	ExecutorType             string
	Labels                   map[string]string
	Variables                map[string]string
	SecretVariables          map[string]string
	Schedule                 string
	ExecutorArgs             []string
	ArgsMode                 string
	ExecutionName            string
	VariablesFile            string
	Envs                     map[string]string
	SecretEnvs               map[string]string
	HttpProxy, HttpsProxy    string
	SecretVariableReferences map[string]string
	CopyFiles                []string
	Image                    string
	Command                  []string
	ImagePullSecretNames     []string
	Timeout                  int64
	ArtifactStorageClassName string
	ArtifactVolumeMountPath  string
	ArtifactDirs             []string
	JobTemplate              string
	PreRunScript             string
	ScraperTemplate          string
	NegativeTest             bool
	MountConfigMaps          map[string]string
	VariableConfigMaps       []string
	MountSecrets             map[string]string
	VariableSecrets          []string
	UploadTimeout            string
}

// NewCreateTestsCmd is a command tp create new Test Custom Resource
func NewCreateTestsCmd() *cobra.Command {

	var (
		testName             string
		testContentType      string
		file                 string
		uri                  string
		gitUri               string
		gitBranch            string
		gitCommit            string
		gitPath              string
		gitWorkingDir        string
		gitUsername          string
		gitToken             string
		gitUsernameSecret    map[string]string
		gitTokenSecret       map[string]string
		gitCertificateSecret string
		gitAuthType          string
		sourceName           string
		flags                CreateCommonFlags
	)

	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"tests", "t"},
		Short:   "Create new Test",
		Long:    `Create new Test Custom Resource`,
		Run: func(cmd *cobra.Command, args []string) {
			crdOnly, err := strconv.ParseBool(cmd.Flag("crd-only").Value.String())
			ui.ExitOnError("parsing flag value", err)

			if testName == "" {
				ui.Failf("pass valid test name (in '--name' flag)")
			}

			namespace := cmd.Flag("namespace").Value.String()
			var client client.Client
			if !crdOnly {
				client, namespace = common.GetClient(cmd)
				test, _ := client.GetTest(testName)

				if testName == test.Name {
					ui.Failf("Test with name '%s' already exists in namespace %s", testName, namespace)
				}
			}

			if cmd.Flag("git-uri") != nil {
				err = common.ValidateUpsertOptions(cmd, sourceName)
				ui.ExitOnError("validating passed flags", err)
			}

			err = validateArtifactRequest(flags.ArtifactStorageClassName, flags.ArtifactVolumeMountPath, flags.ArtifactDirs)
			ui.ExitOnError("validating artifact flags", err)

			options, err := NewUpsertTestOptionsFromFlags(cmd)
			ui.ExitOnError("getting test options", err)

			if !crdOnly {
				executors, err := client.ListExecutors("")
				ui.ExitOnError("getting available executors", err)

				contentType := ""
				if options.Content != nil {
					contentType = options.Content.Type_
				}

				err = validateExecutorTypeAndContent(options.Type_, contentType, executors)
				ui.ExitOnError("validating executor type", err)

				if len(flags.CopyFiles) > 0 {
					var timeout time.Duration
					if flags.UploadTimeout != "" {
						timeout, err = time.ParseDuration(flags.UploadTimeout)
						if err != nil {
							ui.ExitOnError("invalid upload timeout duration", err)
						}
					}
					err := uploadFiles(client, testName, apiv1.Test, flags.CopyFiles, timeout)
					ui.ExitOnError("could not upload files", err)
				}

				_, err = client.CreateTest(options)
				ui.ExitOnError("creating test "+testName+" in namespace "+namespace, err)

				ui.Success("Test created", namespace, "/", testName)
			} else {
				(*testkube.TestUpsertRequest)(&options).QuoteTestTextFields()
				data, err := crd.ExecuteTemplate(crd.TemplateTest, options)
				ui.ExitOnError("executing crd template", err)

				ui.Info(data)
			}
		},
	}

	cmd.Flags().StringVarP(&testName, "name", "n", "", "unique test name - mandatory")
	cmd.Flags().StringVarP(&testContentType, "test-content-type", "", "", "content type of test one of string|file-uri|git")

	// create options
	cmd.Flags().StringVarP(&file, "file", "f", "", "test file - will be read from stdin if not specified")
	cmd.Flags().StringVarP(&uri, "uri", "", "", "URI of resource - will be loaded by http GET")
	cmd.Flags().StringVarP(&gitUri, "git-uri", "", "", "Git repository uri")
	cmd.Flags().StringVarP(&gitBranch, "git-branch", "", "", "if uri is git repository we can set additional branch parameter")
	cmd.Flags().StringVarP(&gitCommit, "git-commit", "", "", "if uri is git repository we can use commit id (sha) parameter")
	cmd.Flags().StringVarP(&gitPath, "git-path", "", "", "if repository is big we need to define additional path to directory/file to checkout partially")
	cmd.Flags().StringVarP(&gitWorkingDir, "git-working-dir", "", "", "if repository contains multiple directories with tests (like monorepo) and one starting directory we can set working directory parameter")
	cmd.Flags().StringVarP(&gitUsername, "git-username", "", "", "if git repository is private we can use username as an auth parameter")
	cmd.Flags().StringVarP(&gitToken, "git-token", "", "", "if git repository is private we can use token as an auth parameter")
	cmd.Flags().StringToStringVarP(&gitUsernameSecret, "git-username-secret", "", map[string]string{}, "git username secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringToStringVarP(&gitTokenSecret, "git-token-secret", "", map[string]string{}, "git token secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringVarP(&gitCertificateSecret, "git-certificate-secret", "", "", "if git repository is private we can use certificate as an auth parameter stored in a kubernetes secret name")
	cmd.Flags().StringVarP(&gitAuthType, "git-auth-type", "", "basic", "auth type for git requests one of basic|header")
	cmd.Flags().StringVarP(&sourceName, "source", "", "", "source name - will be used together with content parameters")
	cmd.Flags().MarkDeprecated("env", "env is deprecated use variable instead")
	cmd.Flags().MarkDeprecated("secret-env", "secret-env is deprecated use secret-variable instead")

	AddCreateFlags(cmd, &flags)

	return cmd
}

// AddCreateFlags adds flags to the create command that can be used by the create from file
func AddCreateFlags(cmd *cobra.Command, flags *CreateCommonFlags) {

	cmd.Flags().StringVarP(&flags.ExecutorType, "type", "t", "", "test type")

	cmd.Flags().StringToStringVarP(&flags.Labels, "label", "l", nil, "label key value pair: --label key1=value1")
	cmd.Flags().StringToStringVarP(&flags.Variables, "variable", "v", nil, "variable key value pair: --variable key1=value1")
	cmd.Flags().StringToStringVarP(&flags.SecretVariables, "secret-variable", "s", nil, "secret variable key value pair: --secret-variable key1=value1")
	cmd.Flags().StringVarP(&flags.Schedule, "schedule", "", "", "test schedule in a cronjob form: * * * * *")
	cmd.Flags().StringArrayVar(&flags.Command, "command", []string{}, "command passed to image in executor")
	cmd.Flags().StringArrayVarP(&flags.ExecutorArgs, "executor-args", "", []string{}, "executor binary additional arguments")
	cmd.Flags().StringVarP(&flags.ArgsMode, "args-mode", "append", "", "usage mode for arguments. one of append|override")
	cmd.Flags().StringVarP(&flags.ExecutionName, "execution-name", "", "", "execution name, if empty will be autogenerated")
	cmd.Flags().StringVarP(&flags.VariablesFile, "variables-file", "", "", "variables file path, e.g. postman env file - will be passed to executor if supported")
	cmd.Flags().StringToStringVarP(&flags.Envs, "env", "", map[string]string{}, "envs in a form of name1=val1 passed to executor")
	cmd.Flags().StringToStringVarP(&flags.SecretEnvs, "secret-env", "", map[string]string{}, "secret envs in a form of secret_key1=secret_name1 passed to executor")
	cmd.Flags().StringVar(&flags.HttpProxy, "http-proxy", "", "http proxy for executor containers")
	cmd.Flags().StringVar(&flags.HttpsProxy, "https-proxy", "", "https proxy for executor containers")
	cmd.Flags().StringToStringVarP(&flags.SecretVariableReferences, "secret-variable-reference", "", nil, "secret variable references in a form name1=secret_name1=secret_key1")
	cmd.Flags().StringArrayVarP(&flags.CopyFiles, "copy-files", "", []string{}, "file path mappings from host to pod of form source:destination")
	cmd.Flags().StringVar(&flags.Image, "image", "", "image for container executor")
	cmd.Flags().StringArrayVar(&flags.ImagePullSecretNames, "image-pull-secrets", []string{}, "secret name used to pull the image in container executor")
	cmd.Flags().Int64Var(&flags.Timeout, "timeout", 0, "duration in seconds for test to timeout. 0 disables timeout.")
	cmd.Flags().StringVar(&flags.ArtifactStorageClassName, "artifact-storage-class-name", "", "artifact storage class name for container executor")
	cmd.Flags().StringVar(&flags.ArtifactVolumeMountPath, "artifact-volume-mount-path", "", "artifact volume mount path for container executor")
	cmd.Flags().StringArrayVarP(&flags.ArtifactDirs, "artifact-dir", "", []string{}, "artifact dirs for container executor")
	cmd.Flags().StringVar(&flags.JobTemplate, "job-template", "", "job template file path for extensions to job template")
	cmd.Flags().StringVarP(&flags.PreRunScript, "prerun-script", "", "", "path to script to be run before test execution")
	cmd.Flags().StringVar(&flags.ScraperTemplate, "scraper-template", "", "scraper template file path for extensions to scraper template")
	cmd.Flags().BoolVar(&flags.NegativeTest, "negative-test", false, "negative test, if enabled, makes failure an expected and correct test result. If the test fails the result will be set to success, and vice versa")
	cmd.Flags().StringToStringVarP(&flags.MountConfigMaps, "mount-configmap", "", map[string]string{}, "config map value pair for mounting it to executor pod: --mount-configmap configmap_name=configmap_mountpath")
	cmd.Flags().StringArrayVar(&flags.VariableConfigMaps, "variable-configmap", []string{}, "config map name used to map all keys to basis variables")
	cmd.Flags().StringToStringVarP(&flags.MountSecrets, "mount-secret", "", map[string]string{}, "secret value pair for mounting it to executor pod: --mount-secret secret_name=secret_mountpath")
	cmd.Flags().StringArrayVar(&flags.VariableSecrets, "variable-secret", []string{}, "secret name used to map all keys to secret variables")
	cmd.Flags().StringVar(&flags.UploadTimeout, "upload-timeout", "", "timeout to use when uploading files, example: 30s")
}

func validateExecutorTypeAndContent(executorType, contentType string, executors testkube.ExecutorsDetails) error {
	typeValid := false
	executorTypes := []string{}
	contentTypes := []string{}

	for _, ed := range executors {
		executorTypes = append(executorTypes, ed.Executor.Types...)
		for _, et := range ed.Executor.Types {
			if et == executorType {
				typeValid = true
				contentTypes = ed.Executor.ContentTypes
				break
			}
		}
	}

	if !typeValid {
		return fmt.Errorf("invalid executor type '%s' use one of: %v", executorType, executorTypes)
	}

	if len(contentTypes) != 0 {
		contentValid := false
		for _, ct := range contentTypes {
			if ct == contentType {
				contentValid = true
				break
			}
		}

		if !contentValid {
			return fmt.Errorf("invalid content type '%s' use one of: %v", contentType, contentTypes)
		}
	}

	return nil
}

func validateArtifactRequest(artifactStorageClassName, artifactVolumeMountPath string, artifactDirs []string) error {
	if artifactStorageClassName != "" || artifactVolumeMountPath != "" || len(artifactDirs) != 0 {
		if artifactStorageClassName == "" || artifactVolumeMountPath == "" {
			return fmt.Errorf("both artifact storage class name and mount path should be provided")
		}
	}

	return nil
}
