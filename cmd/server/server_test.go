package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"turkic-mythology-gateway/pkg/config"
)

func Test_Server(t *testing.T) {
	testConfig := setupTestConfig()

	t.Run("should create server and return server", func(t *testing.T) {
		testServer := NewGatewayServer(testConfig)

		expected := &server{
			config: testConfig,
			fiber:  testServer.GetFiberInstance(),
		}

		assert.Equal(t, expected, testServer)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		testServer := NewGatewayServer(testConfig)

		go func() {
			err := testServer.Start()
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
		}()
	})
}

func setupTestConfig() config.Config {
	testPath := "./testdata/config.yaml"
	readConfig, _ := config.ReadConfig(testPath)
	createConfig := config.Init(readConfig)

	return createConfig
}
