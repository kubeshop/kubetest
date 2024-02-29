// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//	https://github.com/kubeshop/testkube/blob/main/licenses/TCL.txt

package main

import (
	"fmt"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"

	"github.com/kballard/go-shellquote"

	"github.com/kubeshop/testkube/cmd/tcl/testworkflow-init/data"
	"github.com/kubeshop/testkube/cmd/tcl/testworkflow-init/run"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing ref!")
		os.Exit(155)
	}
	data.Step.Ref = os.Args[1]

	now := time.Now()

	// Load shared state
	data.LoadState()

	// Initialize space for parsing args
	config := map[string]string{}
	computed := []string(nil)
	conditions := []data.Rule(nil)
	resulting := []data.Rule(nil)
	timeouts := []data.Timeout(nil)
	args := []string(nil)

	// Read arguments into the base data
	for i := 2; i < len(os.Args); i += 2 {
		if i+1 == len(os.Args) {
			break
		}
		switch os.Args[i] {
		case "--":
			args = os.Args[i+1:]
			i = len(os.Args)
		case "-i", "--init":
			data.Step.InitStatus = os.Args[i+1]
		case "-c", "--cond":
			v := strings.SplitN(os.Args[i+1], "=", 2)
			refs := strings.Split(v[0], ",")
			if len(v) == 2 {
				conditions = append(conditions, data.Rule{Expr: v[1], Refs: refs})
			} else {
				conditions = append(conditions, data.Rule{Expr: "true", Refs: refs})
			}
		case "-r", "--result":
			v := strings.SplitN(os.Args[i+1], "=", 2)
			refs := strings.Split(v[0], ",")
			if len(v) == 2 {
				resulting = append(resulting, data.Rule{Expr: v[1], Refs: refs})
			} else {
				resulting = append(resulting, data.Rule{Expr: "true", Refs: refs})
			}
		case "-t", "--timeout":
			v := strings.SplitN(os.Args[i+1], "=", 2)
			if len(v) == 2 {
				timeouts = append(timeouts, data.Timeout{Ref: v[0], Duration: v[1]})
			} else {
				timeouts = append(timeouts, data.Timeout{Ref: v[0], Duration: ""})
			}
		case "-e", "--env":
			computed = append(computed, strings.Split(os.Args[i+1], ",")...)
		default:
			config[strings.TrimLeft(os.Args[i], "-")] = os.Args[i+1]
		}
	}

	// Compute environment variables
	for _, name := range computed {
		initial := os.Getenv(name)
		value, err := data.Template(initial)
		if err != nil {
			fmt.Printf(`resolving "%s" environment variable: %s: %s\n`, name, initial, err.Error())
			os.Exit(155)
		}
		_ = os.Setenv(name, value)
	}

	// Compute conditional steps
	for _, c := range conditions {
		expr, err := data.Expression(c.Expr)
		if err != nil {
			fmt.Printf("broken condition for refs: %s: %s: %s\n", strings.Join(c.Refs, ", "), c.Expr, err.Error())
			os.Exit(155)
		}
		v, _ := expr.BoolValue()
		if !v {
			for _, r := range c.Refs {
				data.State.GetStep(r).Skip(now)
			}
		}
	}

	// Start all acknowledged steps
	for _, f := range resulting {
		for _, r := range f.Refs {
			if r != "" {
				data.State.GetStep(r).Start(now)
			}
		}
	}
	for _, t := range timeouts {
		if t.Ref != "" {
			data.State.GetStep(t.Ref).Start(now)
		}
	}
	data.State.GetStep(data.Step.Ref).Start(now)

	// Register timeouts
	for _, t := range timeouts {
		err := data.State.GetStep(t.Ref).SetTimeoutDuration(now, t.Duration)
		if err != nil {
			fmt.Printf("broken timeout for ref: %s: %s: %s\n", t.Ref, t.Duration, err.Error())
			os.Exit(155)
		}
	}

	// Save the resulting conditions
	data.Config.Resulting = resulting

	// Don't call further if the step is already skipped
	if data.State.GetStep(data.Step.Ref).Status == data.StepStatusSkipped {
		if data.Config.Debug {
			fmt.Printf("Skipped.\n")
		}
		data.Finish()
	}

	// Load the rest of the configuration
	for k, v := range config {
		value, err := data.Template(v)
		if err != nil {
			fmt.Printf(`resolving "%s" param: %s: %s\n`, k, v, err.Error())
			os.Exit(155)
		}
		data.LoadConfig(map[string]string{k: value})
	}

	// Compute templates in the cmd/args
	original := slices.Clone(args)
	var err error
	for i := range args {
		args[i], err = data.Template(args[i])
		if err != nil {
			fmt.Printf(`resolving command: %s: %s\n`, shellquote.Join(original...), err.Error())
			os.Exit(155)
		}
	}

	// Fail when there is nothing to run
	if len(args) == 0 {
		fmt.Println("missing command to run")
		os.Exit(189)
	}

	// Handle aborting
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopSignal
		fmt.Println("The task was aborted.")
		data.Step.Status = data.StepStatusAborted
		data.Step.ExitCode = 137
		data.Finish()
	}()

	// Handle timeouts
	for _, t := range timeouts {
		go func(ref string) {
			time.Sleep(data.State.GetStep(ref).TimeoutAt.Sub(time.Now()))
			fmt.Printf("Timed out.\n")
			data.State.GetStep(ref).SetStatus(data.StepStatusTimeout)
			data.Step.Status = data.StepStatusTimeout
			data.Step.ExitCode = 124
			data.Finish()
		}(t.Ref)
	}

	// Start the task
	data.Step.Executed = true
	run.Run(args[0], args[1:])

	os.Exit(0)
}
