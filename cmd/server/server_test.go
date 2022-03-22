package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"turkic-mythology-gateway/pkg/config"
)

func Test_Server(t *testing.T) {
	testPath := "./testdata/config.yaml"
	readConfig, _ := config.ReadConfig(testPath)
	createConfig := config.Init(readConfig)

	t.Run("should create server and return server", func(t *testing.T) {
		testServer := NewGatewayServer(createConfig)

		expected := &server{
			config: createConfig,
			fiber:  testServer.GetFiberInstance(),
		}

		assert.Equal(t, expected, testServer)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		testServer := NewGatewayServer(createConfig)

		go func() {
			err := testServer.Start()
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
		}()
	})
}
