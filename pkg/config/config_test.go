//go:build unit

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	var err error

	err = os.Setenv(ServerPort, "8080")
	require.NoError(t, err)

	err = os.Setenv(PactBrokerURL, "https://test-broker.com")
	require.NoError(t, err)

	err = os.Setenv(PactBrokerToken, "test-broker-token")
	require.NoError(t, err)

	var cfg Config
	cfg, err = ReadConfig()

	assert.NoError(t, err)
	assert.IsType(t, &config{}, cfg)
}

func TestConfig_GetPactBrokerToken(t *testing.T) {
	cfg := &config{
		pactBrokerToken: "pact-broker-token",
	}
	token := cfg.GetPactBrokerToken()

	assert.Equal(t, token, cfg.pactBrokerToken)
}

func TestConfig_GetServerPort(t *testing.T) {
	cfg := &config{
		serverPort: "8080",
	}
	serverPort := cfg.GetServerPort()

	assert.Equal(t, serverPort, cfg.serverPort)
}

func TestConfig_GetPactBrokerURl(t *testing.T) {
	cfg := &config{
		pactBrokerURl: "https://test-broker.com",
	}
	pactBrokerURl := cfg.GetPactBrokerURl()

	assert.Equal(t, pactBrokerURl, cfg.pactBrokerURl)
}
