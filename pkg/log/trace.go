// Simple wrapper for zap logger to provide trace logging until Zap will manage additional log levels
// https://github.com/uber-go/zap/issues/680
package log

import (
	"os"

	"go.uber.org/zap"
)

const tracePrefix = "TRACE: "

func Tracew(logger *zap.SugaredLogger, msg string, keysAndValues ...interface{}) {
	if isTraceEnabled() {
		logger.Debugw(tracePrefix+msg, keysAndValues...)
	}
}

func Tracef(logger *zap.SugaredLogger, msg string, keysAndValues ...interface{}) {
	if isTraceEnabled() {
		logger.Debugf(tracePrefix+msg, keysAndValues...)
	}
}

func Trace(logger *zap.SugaredLogger, msg string, keysAndValues ...interface{}) {
	if isTraceEnabled() {
		logger.Debug(tracePrefix + msg)
	}
}

func Traceln(logger *zap.SugaredLogger, msg string, keysAndValues ...interface{}) {
	if isTraceEnabled() {
		logger.Debug(tracePrefix + msg)
	}
}

func isTraceEnabled() bool {
	return os.Getenv("TRACE") == "true"
}
