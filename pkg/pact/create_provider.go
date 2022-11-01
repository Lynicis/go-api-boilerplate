package pact

import (
	"path/filepath"

	"github.com/pact-foundation/pact-go/types"

	"go-rest-api-boilerplate/pkg/project"
)

func CreateProvider(
	hostName string,
	pactBrokerURl string,
	pactBrokerToken string,
	stateHandlers types.StateHandlers,
) *types.VerifyRequest {
	return &types.VerifyRequest{
		ProviderBaseURL:            hostName,
		BrokerURL:                  pactBrokerURl,
		BrokerToken:                pactBrokerToken,
		StateHandlers:              stateHandlers,
		PactLogLevel:               "INFO",
		PactLogDir:                 filepath.Join(project.GetRootDirectory(), ".pact/logs"),
		PublishVerificationResults: true,
	}
}
