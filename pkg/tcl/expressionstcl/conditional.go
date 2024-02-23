// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package expressionstcl

import (
	"fmt"
	"maps"
)

type conditional struct {
	condition Expression
	truthy    Expression
	falsy     Expression
}

func newConditional(condition, truthy, falsy Expression) Expression {
	return &conditional{condition: condition, truthy: truthy, falsy: falsy}
}

func (s *conditional) String() string {
	return fmt.Sprintf("%s ? %s : %s", s.condition.String(), s.truthy.String(), s.falsy.String())
}

func (s *conditional) SafeString() string {
	return "(" + s.String() + ")"
}

func (s *conditional) Template() string {
	return "{{" + s.String() + "}}"
}

func (s *conditional) SafeSimplify(m ...MachineCore) (v Expression, changed bool, err error) {
	var ch bool
	s.condition, ch, err = s.condition.SafeSimplify(m...)
	changed = changed || ch
	if err != nil {
		return nil, changed, err
	}
	if s.condition.Static() != nil {
		var b bool
		b, err = s.condition.Static().BoolValue()
		if err != nil {
			return nil, true, err
		}
		if b {
			return s.truthy, true, err
		}
		return s.falsy, true, err
	}
	s.truthy, ch, err = s.truthy.SafeSimplify(m...)
	changed = changed || ch
	if err != nil {
		return nil, changed, err
	}
	s.falsy, ch, err = s.falsy.SafeSimplify(m...)
	changed = changed || ch
	if err != nil {
		return nil, changed, err
	}
	return s, changed, nil
}

func (s *conditional) Simplify(m ...MachineCore) (v Expression, err error) {
	return deepSimplify(s, m...)
}

func (s *conditional) Static() StaticValue {
	return nil
}

func (s *conditional) Accessors() map[string]struct{} {
	result := make(map[string]struct{})
	maps.Copy(result, s.condition.Accessors())
	maps.Copy(result, s.truthy.Accessors())
	maps.Copy(result, s.falsy.Accessors())
	return result
}

func (s *conditional) Functions() map[string]struct{} {
	result := make(map[string]struct{})
	maps.Copy(result, s.condition.Functions())
	maps.Copy(result, s.truthy.Functions())
	maps.Copy(result, s.falsy.Functions())
	return result
}
