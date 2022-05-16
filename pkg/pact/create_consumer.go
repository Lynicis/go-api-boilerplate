package pact

import (
	"path"
	"time"

	"github.com/pact-foundation/pact-go/dsl"

	"go-rest-api-boilerplate/pkg/project_path"
)

func CreateConsumer(provider, consumer string) *dsl.Pact {
	rootDirectory := project_path.GetRootDirectory()
	return &dsl.Pact{
		Host:              "127.0.0.1",
		Consumer:          provider,
		Provider:          consumer,
		ClientTimeout:     10 * time.Second,
		LogLevel:          "INFO",
		LogDir:            path.Join(rootDirectory, ".pact/logs"),
		PactFileWriteMode: "merge",
		PactDir:           path.Join(rootDirectory, ".pact/pacts"),
	}
}
