package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLogger(t *testing.T) {
	loggerInstance := CreateLogger()

	expected := &logger{}

	assert.IsType(t, expected, loggerInstance)
}
