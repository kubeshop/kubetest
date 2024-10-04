package stage

import (
	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/expressions"
	"github.com/kubeshop/testkube/pkg/imageinspector"
)

type groupStage struct {
	stageMetadata
	stageLifecycle
	containerDefaults Container
	children          []Stage
	virtual           bool
	pure              *bool
}

type GroupStage interface {
	Stage
	SetContainerDefaults(c Container)
	ContainerDefaults() Container
	Children() []Stage
	RecursiveChildren() []Stage
	Add(stages ...Stage) GroupStage
	Pure() *bool
	SetPure(pure *bool) GroupStage
}

func NewGroupStage(ref string, virtual bool) GroupStage {
	return &groupStage{
		stageMetadata: stageMetadata{ref: ref},
		virtual:       virtual,
	}
}

func (s *groupStage) SetContainerDefaults(c Container) {
	s.containerDefaults = c
}

func (s *groupStage) ContainerDefaults() Container {
	return s.containerDefaults
}

func (s *groupStage) Len() int {
	count := 0
	for _, ch := range s.Children() {
		count += ch.Len()
	}
	return count
}

func (s *groupStage) HasPause() bool {
	return s.paused || (len(s.Children()) > 0 && s.Children()[0].HasPause())
}

func (s *groupStage) signature(full bool) Signature {
	sig := []Signature(nil)
	for _, ch := range s.Children() {
		var si Signature
		if full {
			si = ch.FullSignature()
		} else {
			si = ch.Signature()
		}
		_, ok := ch.(GroupStage)
		// Include children directly, if the stage is virtual
		if !full && ok && si.Name() == "" && !si.Optional() && !si.Negative() {
			sig = append(sig, si.Children()...)
		} else {
			sig = append(sig, si)
		}
	}

	return &signature{
		RefValue:      s.ref,
		NameValue:     s.name,
		CategoryValue: s.category,
		OptionalValue: s.optional,
		NegativeValue: s.negative,
		ChildrenValue: sig,
	}
}

func (s *groupStage) Signature() Signature {
	return s.signature(false)
}

func (s *groupStage) FullSignature() Signature {
	return s.signature(true)
}

func (s *groupStage) ContainerStages() []ContainerStage {
	c := []ContainerStage(nil)
	for _, ch := range s.children {
		c = append(c, ch.ContainerStages()...)
	}
	return c
}

func (s *groupStage) Children() []Stage {
	return s.children
}

func (s *groupStage) RecursiveChildren() []Stage {
	res := make([]Stage, 0)
	for _, ch := range s.children {
		if v, ok := ch.(GroupStage); ok {
			res = append(res, v.RecursiveChildren()...)
		} else {
			res = append(res, ch)
		}
	}
	return res
}

func (s *groupStage) GetImages(isGroupNeeded bool) map[string]bool {
	v := make(map[string]bool)
	for _, ch := range s.children {
		for name, needsMetadata := range ch.GetImages(isGroupNeeded) {
			v[name] = v[name] || needsMetadata
		}
	}
	return v
}

func (s *groupStage) Flatten() []Stage {
	// Flatten children
	next := []Stage(nil)
	for _, ch := range s.children {
		next = append(next, ch.Flatten()...)
	}
	s.children = next

	// Delete empty stage
	if len(s.children) == 0 {
		return nil
	}

	// Flatten when it is completely virtual stage
	if s.virtual {
		return s.children
	}

	// Merge stage into single one below if possible
	first := s.children[0]
	if len(s.children) == 1 && (s.name == "" || first.Name() == "") && (s.timeout == "" || first.Timeout() == "") && (!s.paused || !first.Paused()) {
		if first.Name() == "" {
			first.SetName(s.name)
		}
		if first.Condition() == "" {
			// Virtualize with the default condition
			first.AppendConditions(s.condition, "passed")
		} else {
			first.AppendConditions(s.condition)
		}
		if first.Timeout() == "" {
			first.SetTimeout(s.timeout)
		}
		if s.negative {
			first.SetNegative(!first.Negative())
		}
		if s.optional {
			first.SetOptional(true)
		}
		if s.paused {
			first.SetPaused(true)
		}
		if s.pure != nil {
			if firstContainer, ok := first.(ContainerStage); ok && *s.pure {
				firstContainer.SetPure(true)
			} else if firstGroup, ok := first.(GroupStage); ok && firstGroup.Pure() == nil {
				firstGroup.SetPure(s.pure)
			}
		}
		return []Stage{first}
	}

	return []Stage{s}
}

func (s *groupStage) Add(stages ...Stage) GroupStage {
	for _, ch := range stages {
		if ch != nil {
			s.children = append(s.children, ch.Flatten()...)
		}
	}
	return s
}

func (s *groupStage) ApplyImages(images map[string]*imageinspector.Info, imageNameResolutions map[string]string) error {
	for i := range s.children {
		err := s.children[i].ApplyImages(images, imageNameResolutions)
		if err != nil {
			return errors.Wrap(err, "applying image data")
		}
	}
	return nil
}

func (s *groupStage) Resolve(m ...expressions.Machine) error {
	for i := range s.children {
		err := s.children[i].Resolve(m...)
		if err != nil {
			return errors.Wrap(err, "group stage container")
		}
	}
	return expressions.Simplify(&s.stageMetadata, m...)
}

func (s *groupStage) Pure() *bool {
	return s.pure
}

func (s *groupStage) SetPure(pure *bool) GroupStage {
	s.pure = pure
	return s
}
