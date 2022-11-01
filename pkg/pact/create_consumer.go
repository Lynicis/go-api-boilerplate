package pact

import (
	"path/filepath"
	"time"

	"github.com/pact-foundation/pact-go/dsl"

	"go-rest-api-boilerplate/pkg/project"
)

func CreateConsumer(hostName, providerName, consumerName string) *dsl.Pact {
	rootDirectory := project.GetRootDirectory()
	return &dsl.Pact{
		Host:              hostName,
		Consumer:          providerName,
		Provider:          consumerName,
		ClientTimeout:     10 * time.Second,
		LogLevel:          "INFO",
		LogDir:            filepath.Join(rootDirectory, ".pact/logs"),
		PactFileWriteMode: "merge",
		PactDir:           filepath.Join(rootDirectory, ".pact/pacts"),
	}
}
