package logger

import (
	"go.uber.org/zap"
)

// Logger zap logger methods
type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Errorf(format string, args ...interface{})
	Error(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Sync() error
}

type logger struct {
	*zap.SugaredLogger
}

// CreateLogger create new logger instance
func CreateLogger() Logger {
	zapInstance, _ := zap.NewProduction()
	sugarLogger := zapInstance.Sugar()

	return &logger{
		sugarLogger,
	}
}
