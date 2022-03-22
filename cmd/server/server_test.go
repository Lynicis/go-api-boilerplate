//go:build unit

package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"turkic-mythology/pkg/config"
	"turkic-mythology/pkg/path"
)

func Test_Server(t *testing.T) {
	testConfig := setupTestConfig()

	t.Run("should create server instance and return server instance", func(t *testing.T) {
		testServer := NewServer(testConfig)

		expected := &server{
			config: testConfig,
			fiber:  testServer.GetFiberInstance(),
		}

		assert.Equal(t, expected, testServer)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		testServer := NewServer(testConfig)

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
	basePath := path.GetProjectBasePath()
	testPath := fmt.Sprintf("%s/testdata/development.yaml", basePath)
	testAppEnvironment := "test"
	readConfig, _ := config.ReadConfig(testPath)
	createConfig := config.Init(readConfig, testAppEnvironment)

	return createConfig
}
