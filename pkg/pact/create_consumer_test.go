//go:build unit

package pact

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestPact_CreateConsumer(t *testing.T) {
	actualConsumer := CreateConsumer(
		"localhost.com",
		"TestProvider",
		"TestConsumer",
	)

	expectedPact := &dsl.Pact{}

	assert.IsType(t, expectedPact, actualConsumer)
}
