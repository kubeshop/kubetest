package adapter

import (
	"github.com/kubeshop/testkube/pkg/log"
	"github.com/kubeshop/testkube/pkg/logs/events"
	"go.uber.org/zap"
)

var _ Adapter = &DebugAdapter{}

// NewDebugAdapter creates new DebugAdapter which will write logs to stdout
func NewDebugAdapter() *DebugAdapter {
	return &DebugAdapter{
		l: log.DefaultLogger,
	}
}

type DebugAdapter struct {
	l *zap.SugaredLogger
}

func (s *DebugAdapter) Notify(id string, e events.Log) error {
	s.l.Debugw("got event", "id", id, "event", e)
	return nil
}

func (s *DebugAdapter) Stop(id string) error {
	s.l.Debugw("Stopping", "id", id)
	return nil
}

func (s *DebugAdapter) Name() string {
	return "dummy"
}
