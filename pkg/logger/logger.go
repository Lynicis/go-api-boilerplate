package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Errorf(format string, args ...interface{})
	Error(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

func CreateLogger() Logger {
	zapInstance, _ := zap.NewProduction()
	sugarLogger := zapInstance.Sugar()

	return &logger{
		sugarLogger,
	}
}
