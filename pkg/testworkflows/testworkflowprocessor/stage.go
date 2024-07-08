package testworkflowprocessor

import (
	"github.com/kubeshop/testkube/pkg/expressions"
	"github.com/kubeshop/testkube/pkg/imageinspector"
)

//go:generate mockgen -destination=./mock_stage.go -package=testworkflowprocessor "github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor" Stage
type Stage interface {
	StageMetadata
	StageLifecycle
	Len() int
	HasPause() bool
	Signature() Signature
	Resolve(m ...expressions.Machine) error
	ContainerStages() []ContainerStage
	GetImages() map[string]struct{}
	ApplyImages(images map[string]*imageinspector.Info, imageNameResolutions map[string]string) error
	Flatten() []Stage
}
