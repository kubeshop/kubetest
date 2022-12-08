package executors

import (
	"fmt"
	"os"

	apiClient "github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

// NewUpsertExecutorOptionsFromFlags creates upsert executor options fom command flags
func NewUpsertExecutorOptionsFromFlags(cmd *cobra.Command) (options apiClient.UpsertExecutorOptions, err error) {
	name := cmd.Flag("name").Value.String()
	types, err := cmd.Flags().GetStringArray("types")
	if err != nil {
		return options, err
	}

	executorType := cmd.Flag("executor-type").Value.String()
	uri := cmd.Flag("uri").Value.String()
	image := cmd.Flag("image").Value.String()
	command, err := cmd.Flags().GetStringArray("command")
	if err != nil {
		return options, err
	}

	executorArgs, err := cmd.Flags().GetStringArray("args")
	if err != nil {
		return options, err
	}

	jobTemplate := cmd.Flag("job-template").Value.String()
	jobTemplateContent := ""
	if jobTemplate != "" {
		b, err := os.ReadFile(jobTemplate)
		ui.ExitOnError("reading job template", err)
		jobTemplateContent = string(b)
	}

	imagePullSecretNames, err := cmd.Flags().GetStringArray("image-pull-secrets")
	if err != nil {
		return options, err
	}

	var imageSecrets []testkube.LocalObjectReference
	for _, secretName := range imagePullSecretNames {
		imageSecrets = append(imageSecrets, testkube.LocalObjectReference{Name: secretName})
	}

	labels, err := cmd.Flags().GetStringToString("label")
	if err != nil {
		return options, err
	}

	features, err := cmd.Flags().GetStringArray("feature")
	if err != nil {
		return options, err
	}

	options = apiClient.UpsertExecutorOptions{
		Name:             name,
		Types:            types,
		ExecutorType:     executorType,
		Image:            image,
		ImagePullSecrets: imageSecrets,
		Command:          command,
		Args:             executorArgs,
		Uri:              uri,
		JobTemplate:      jobTemplateContent,
		Features:         features,
		Labels:           labels,
	}

	return options, nil
}

// NewUpsertExecutorOptionsFromFlags creates update executor options fom command flags
func NewUpdateExecutorOptionsFromFlags(cmd *cobra.Command) (options apiClient.UpdateExecutorOptions, err error) {
	var fields = []struct {
		name        string
		destination **string
	}{
		{
			"name",
			&options.Name,
		},
		{
			"executor-type",
			&options.ExecutorType,
		},
		{
			"uri",
			&options.Uri,
		},
		{
			"image",
			&options.Image,
		},
	}

	for _, field := range fields {
		if cmd.Flag(field.name).Changed {
			value := cmd.Flag(field.name).Value.String()
			*field.destination = &value
		}
	}

	var slices = []struct {
		name        string
		destination **[]string
	}{
		{
			"types",
			&options.Types,
		},
		{
			"command",
			&options.Command,
		},
		{
			"args",
			&options.Args,
		},
		{
			"features",
			&options.Features,
		},
	}

	for _, slice := range slices {
		if cmd.Flag(slice.name).Changed {
			value, err := cmd.Flags().GetStringArray(slice.name)
			if err != nil {
				return options, err
			}

			*slice.destination = &value
		}
	}

	if cmd.Flag("job-template").Changed {
		jobTemplate := cmd.Flag("job-template").Value.String()
		b, err := os.ReadFile(jobTemplate)
		if err != nil {
			return options, fmt.Errorf("reading job template %w", err)
		}

		value := string(b)
		options.JobTemplate = &value
	}

	if cmd.Flag("image-pull-secrets").Changed {
		imagePullSecretNames, err := cmd.Flags().GetStringArray("image-pull-secrets")
		if err != nil {
			return options, err
		}

		var imageSecrets []testkube.LocalObjectReference
		for _, secretName := range imagePullSecretNames {
			imageSecrets = append(imageSecrets, testkube.LocalObjectReference{Name: secretName})
		}

		options.ImagePullSecrets = &imageSecrets
	}

	if cmd.Flag("label").Changed {
		labels, err := cmd.Flags().GetStringToString("label")
		if err != nil {
			return options, err
		}

		options.Labels = &labels
	}

	return options, nil
}
