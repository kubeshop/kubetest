// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//	https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package testworkflowprocessor

import (
	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/tcl/expressionstcl"
)

type containerStage struct {
	stageMetadata  `expr:"include"`
	stageLifecycle `expr:"include"`
	container      Container `expr:"include"`
}

type ContainerStage interface {
	Stage
	Container() Container
}

func NewContainerStage(ref string, container Container) ContainerStage {
	return &containerStage{
		stageMetadata: stageMetadata{ref: ref},
		container:     container.CreateChild(),
	}
}

func (s *containerStage) Len() int {
	return 1
}

func (s *containerStage) Signature() Signature {
	return &signature{
		RefValue:      s.ref,
		NameValue:     s.name,
		OptionalValue: s.optional,
		NegativeValue: s.negative,
		ChildrenValue: nil,
	}
}

func (s *containerStage) ContainerStages() []ContainerStage {
	return []ContainerStage{s}
}

func (s *containerStage) GetImages() map[string]struct{} {
	return map[string]struct{}{s.container.Image(): {}}
}

func (s *containerStage) Flatten() []Stage {
	return []Stage{s}
}

func (s *containerStage) Resolve(m ...expressionstcl.Machine) error {
	err := s.container.Resolve(m...)
	if err != nil {
		return errors.Wrap(err, "stage container")
	}
	return expressionstcl.Simplify(s, m...)
}

func (s *containerStage) Container() Container {
	return s.container
}
