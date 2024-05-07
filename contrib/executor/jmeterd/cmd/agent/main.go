package main

import (
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/contrib/executor/jmeterd/pkg/runner"
	"github.com/kubeshop/testkube/pkg/envs"
	"github.com/kubeshop/testkube/pkg/executor/agent"
	"github.com/kubeshop/testkube/pkg/executor/output"
)

func main() {
	ctx := context.Background()
	agent.PreRun(ctx)
	defer agent.PostRun(ctx)
	params, err := envs.LoadTestkubeVariables()
	if err != nil {
		output.PrintError(os.Stderr, errors.Errorf("could not initialize Distributed JMeter Executor environment variables: %v", err))
		os.Exit(1)
	}
	r, err := runner.NewRunner(ctx, params)
	if err != nil {
		output.PrintError(os.Stderr, errors.Wrap(err, "error instantiating Distributed JMeter Executor"))
		os.Exit(1)
	}
	agent.Run(ctx, r, os.Args)
}
