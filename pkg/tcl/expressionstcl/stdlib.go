// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package expressionstcl

import (
	"encoding/json"
	"fmt"
	math2 "math"
	"strings"

	"github.com/kballard/go-shellquote"
	"gopkg.in/yaml.v3"
)

type StdFunction struct {
	ReturnType Type
	Handler    func(...StaticValue) (Expression, error)
}

type stdMachine struct{}

var StdLibMachine = &stdMachine{}

var stdFunctions = map[string]StdFunction{
	"string": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			str := ""
			for i := range value {
				next, _ := value[i].StringValue()
				str += next
			}
			return NewValue(str), nil
		},
	},
	"list": {
		Handler: func(value ...StaticValue) (Expression, error) {
			v := make([]interface{}, len(value))
			for i := range value {
				v[i] = value[i].Value()
			}
			return NewValue(v), nil
		},
	},
	"join": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) == 0 || len(value) > 2 {
				return nil, fmt.Errorf(`"join" function expects 1-2 arguments, %d provided`, len(value))
			}
			if value[0].IsNone() {
				return value[0], nil
			}
			if !value[0].IsSlice() {
				return nil, fmt.Errorf(`"join" function expects a slice as 1st argument: %v provided`, value[0].Value())
			}
			slice, err := value[0].SliceValue()
			if err != nil {
				return nil, fmt.Errorf(`"join" function error: reading slice: %s`, err.Error())
			}
			v := make([]string, len(slice))
			for i := range slice {
				v[i], _ = toString(slice[i])
			}
			separator := ","
			if len(value) == 2 {
				separator, _ = value[1].StringValue()
			}
			return NewValue(strings.Join(v, separator)), nil
		},
	},
	"split": {
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) == 0 || len(value) > 2 {
				return nil, fmt.Errorf(`"split" function expects 1-2 arguments, %d provided`, len(value))
			}
			str, _ := value[0].StringValue()
			separator := ","
			if len(value) == 2 {
				separator, _ = value[1].StringValue()
			}
			return NewValue(strings.Split(str, separator)), nil
		},
	},
	"int": {
		ReturnType: TypeInt64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"int" function expects 1 argument, %d provided`, len(value))
			}
			v, err := value[0].IntValue()
			if err != nil {
				return nil, err
			}
			return NewValue(v), nil
		},
	},
	"bool": {
		ReturnType: TypeBool,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"bool" function expects 1 argument, %d provided`, len(value))
			}
			v, err := value[0].BoolValue()
			if err != nil {
				return nil, err
			}
			return NewValue(v), nil
		},
	},
	"float": {
		ReturnType: TypeFloat64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"float" function expects 1 argument, %d provided`, len(value))
			}
			v, err := value[0].FloatValue()
			if err != nil {
				return nil, err
			}
			return NewValue(v), nil
		},
	},
	"tojson": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"tojson" function expects 1 argument, %d provided`, len(value))
			}
			b, err := json.Marshal(value[0].Value())
			if err != nil {
				return nil, fmt.Errorf(`"tojson" function had problem marshalling: %s`, err.Error())
			}
			return NewValue(string(b)), nil
		},
	},
	"json": {
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"json" function expects 1 argument, %d provided`, len(value))
			}
			if !value[0].IsString() {
				return nil, fmt.Errorf(`"json" function argument should be a string`)
			}
			var v interface{}
			err := json.Unmarshal([]byte(value[0].Value().(string)), &v)
			if err != nil {
				return nil, fmt.Errorf(`"json" function had problem unmarshalling: %s`, err.Error())
			}
			return NewValue(v), nil
		},
	},
	"toyaml": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"toyaml" function expects 1 argument, %d provided`, len(value))
			}
			b, err := yaml.Marshal(value[0].Value())
			if err != nil {
				return nil, fmt.Errorf(`"toyaml" function had problem marshalling: %s`, err.Error())
			}
			return NewValue(string(b)), nil
		},
	},
	"yaml": {
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"yaml" function expects 1 argument, %d provided`, len(value))
			}
			if !value[0].IsString() {
				return nil, fmt.Errorf(`"yaml" function argument should be a string`)
			}
			var v interface{}
			err := yaml.Unmarshal([]byte(value[0].Value().(string)), &v)
			if err != nil {
				return nil, fmt.Errorf(`"yaml" function had problem unmarshalling: %s`, err.Error())
			}
			return NewValue(v), nil
		},
	},
	"shellquote": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			args := make([]string, len(value))
			for i := range value {
				args[i], _ = value[i].StringValue()
			}
			return NewValue(shellquote.Join(args...)), nil
		},
	},
	"trim": {
		ReturnType: TypeString,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"trim" function expects 1 argument, %d provided`, len(value))
			}
			if !value[0].IsString() {
				return nil, fmt.Errorf(`"trim" function argument should be a string`)
			}
			str, _ := value[0].StringValue()
			return NewValue(strings.TrimSpace(str)), nil
		},
	},
	"len": {
		ReturnType: TypeInt64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"len" function expects 1 argument, %d provided`, len(value))
			}
			if value[0].IsSlice() {
				v, err := value[0].SliceValue()
				return NewValue(int64(len(v))), err
			}
			if value[0].IsString() {
				v, err := value[0].StringValue()
				return NewValue(int64(len(v))), err
			}
			if value[0].IsMap() {
				v, err := value[0].MapValue()
				return NewValue(int64(len(v))), err
			}
			return nil, fmt.Errorf(`"len" function expects string, slice or map, %v provided`, value[0])
		},
	},
	"floor": {
		ReturnType: TypeInt64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"floor" function expects 1 argument, %d provided`, len(value))
			}
			f, err := value[0].FloatValue()
			if err != nil {
				return nil, fmt.Errorf(`"floor" function expects a number, %s provided: %v`, value[0], err)
			}
			return NewValue(int64(math2.Floor(f))), nil
		},
	},
	"ceil": {
		ReturnType: TypeInt64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"ceil" function expects 1 argument, %d provided`, len(value))
			}
			f, err := value[0].FloatValue()
			if err != nil {
				return nil, fmt.Errorf(`"ceil" function expects a number, %s provided: %v`, value[0], err)
			}
			return NewValue(int64(math2.Ceil(f))), nil
		},
	},
	"round": {
		ReturnType: TypeInt64,
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 1 {
				return nil, fmt.Errorf(`"round" function expects 1 argument, %d provided`, len(value))
			}
			f, err := value[0].FloatValue()
			if err != nil {
				return nil, fmt.Errorf(`"round" function expects a number, %s provided: %v`, value[0], err)
			}
			return NewValue(int64(math2.Round(f))), nil
		},
	},
	"chunk": {
		Handler: func(value ...StaticValue) (Expression, error) {
			if len(value) != 2 {
				return nil, fmt.Errorf(`"chunk" function expects 2 arguments, %d provided`, len(value))
			}
			list, err := value[0].SliceValue()
			if err != nil {
				return nil, fmt.Errorf(`"chunk" function expects 1st argument to be a list, %s provided: %v`, value[0], err)
			}
			size, err := value[1].IntValue()
			if err != nil {
				return nil, fmt.Errorf(`"chunk" function expects 2nd argument to be integer, %s provided: %v`, value[1], err)
			}
			if size <= 0 {
				return nil, fmt.Errorf(`"chunk" function expects 2nd argument to be >= 1, %s provided: %v`, value[1], err)
			}
			chunks := make([][]interface{}, 0)
			l := int64(len(list))
			for i := int64(0); i < l; i += size {
				end := i + size
				if end > l {
					end = l
				}
				chunks = append(chunks, list[i:end])
			}
			return NewValue(chunks), nil
		},
	},
}

const (
	stringCastStdFn = "string"
	boolCastStdFn   = "bool"
	intCastStdFn    = "int"
	floatCastStdFn  = "float"
)

func CastToString(v Expression) Expression {
	if v.Static() != nil {
		return NewStringValue(v.Static().Value())
	} else if v.Type() == TypeString {
		return v
	}
	return newCall(stringCastStdFn, []Expression{v})
}

func CastToBool(v Expression) Expression {
	if v.Type() == TypeBool {
		return v
	}
	return newCall(boolCastStdFn, []Expression{v})
}

func CastToInt(v Expression) Expression {
	if v.Type() == TypeInt64 {
		return v
	}
	return newCall(intCastStdFn, []Expression{v})
}

func CastToFloat(v Expression) Expression {
	if v.Type() == TypeFloat64 {
		return v
	}
	return newCall(intCastStdFn, []Expression{v})
}

func IsStdFunction(name string) bool {
	_, ok := stdFunctions[name]
	return ok
}

func GetStdFunctionReturnType(name string) Type {
	return stdFunctions[name].ReturnType
}

func (*stdMachine) Get(name string) (Expression, bool, error) {
	return nil, false, nil
}

func (*stdMachine) Call(name string, args ...StaticValue) (Expression, bool, error) {
	fn, ok := stdFunctions[name]
	if ok {
		exp, err := fn.Handler(args...)
		return exp, true, err
	}
	return nil, false, nil
}
