//go:build unit

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLogger(t *testing.T) {
	loggerInstance := CreateLogger()

	expectedLogger := &logger{}

	assert.IsType(t, expectedLogger, loggerInstance)
}
