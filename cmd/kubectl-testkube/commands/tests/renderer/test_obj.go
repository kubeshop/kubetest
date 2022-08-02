package renderer

import (
	"fmt"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/renderer"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
)

func TestRenderer(ui *ui.UI, obj interface{}) error {
	test, ok := obj.(testkube.Test)
	if !ok {
		return fmt.Errorf("can't use '%T' as testkube.Test in RenderObj for test", obj)
	}

	ui.Warn("Name:     ", test.Name)
	ui.Warn("Namespace:", test.Namespace)
	ui.Warn("Created:  ", test.Created.String())
	if len(test.Labels) > 0 {
		ui.NL()
		ui.Warn("Labels:   ", testkube.MapToString(test.Labels))
	}
	if test.Schedule != "" {
		ui.NL()
		ui.Warn("Schedule: ", test.Schedule)
	}

	if test.ExecutionRequest != nil && len(test.ExecutionRequest.Variables) > 0 {
		renderer.RenderVariables(test.ExecutionRequest.Variables)
	}

	if test.Content != nil {
		ui.NL()
		ui.Info("Content")
		ui.Warn("Type", test.Content.Type_)
		if test.Content.Uri != "" {
			ui.Warn("Uri: ", test.Content.Uri)
		}

		if test.Content.Repository != nil {
			ui.Warn("Repository: ")
			ui.Warn("  Uri:      ", test.Content.Repository.Uri)
			ui.Warn("  Branch:   ", test.Content.Repository.Branch)
			ui.Warn("  Commit:   ", test.Content.Repository.Commit)
			ui.Warn("  Path:     ", test.Content.Repository.Path)
			ui.Warn("  Username: ", test.Content.Repository.Username)
			ui.Warn("  Token:    ", test.Content.Repository.Token)
		}

		if test.Content.Data != "" {
			ui.Warn("Data: ", "\n", test.Content.Data)
		}
	}

	if test.ExecutionRequest != nil && len(test.ExecutionRequest.Args) > 0 {
		ui.Warn("Executor Args:", test.ExecutionRequest.Args...)
	}

	return nil

}
