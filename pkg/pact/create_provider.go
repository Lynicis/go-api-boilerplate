package pact

import (
	"os"
	"path"

	"github.com/pact-foundation/pact-go/types"

	"go-rest-api-boilerplate/pkg/project_path"
)

func CreateProvider(
	provider string,
	stateHandlers types.StateHandlers,
) *types.VerifyRequest {
	return &types.VerifyRequest{
		ProviderBaseURL:            provider,
		BrokerURL:                  os.Getenv("PACT_BROKER_URL"),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		StateHandlers:              stateHandlers,
		PactLogLevel:               "INFO",
		PactLogDir:                 path.Join(project_path.GetRootDirectory(), ".pact/logs"),
		PublishVerificationResults: true,
	}
}
