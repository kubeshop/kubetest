package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/kubeshop/testkube/pkg/event/bus"
	"github.com/kubeshop/testkube/pkg/log"
	"github.com/kubeshop/testkube/pkg/logs"
	"github.com/kubeshop/testkube/pkg/logs/adapter"
	"github.com/kubeshop/testkube/pkg/logs/config"
	"github.com/kubeshop/testkube/pkg/logs/state"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/oklog/run"
)

func main() {
	var g run.Group

	log := log.DefaultLogger.With("service", "logs-service-init")

	ctx, cancel := context.WithCancel(context.Background())

	cfg := Must(config.Get())

	// Event bus
	nc := Must(bus.NewNATSConnection(cfg.NatsURI))
	defer func() {
		log.Infof("closing nats connection")
		nc.Close()
	}()

	js := Must(jetstream.New(nc))

	kv := Must(js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: cfg.KVBucketName}))
	state := state.NewState(kv)

	svc := logs.NewLogsService(nc, js, state).
		WithAddress(cfg.HttpAddress)

	// TODO - add adapters here
	svc.AddAdapter(adapter.NewDummyAdapter())

	g.Add(func() error {
		err := interrupt(log, ctx)
		return err
	}, func(error) {
		log.Warnf("interrupt signal received, canceling context")
		cancel()
	})

	g.Add(func() error {
		return svc.Run(ctx)
	}, func(error) {
		err := svc.Shutdown(ctx)
		if err != nil {
			log.Errorw("error shutting down logs service", "error", err)
		}
		log.Warn("logs service shutdown")
	})

	// We need to do a http health check to be backward compatible with Kubernetes below 1.25
	g.Add(func() error {
		return svc.RunHealthCheckHandler(ctx)
	}, func(error) {
		cancel()
	})

	if err := g.Run(); err != nil {
		log.Warnw("logs service run group returned an error", "error", err)
	}

	log.Infof("exiting")
}

func interrupt(logger *zap.SugaredLogger, ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-c:
		return errors.New("signal received" + s.String())
	case <-ctx.Done():
		return context.Canceled
	}
}

// Must helper function to panic on error
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
