//go:build unit

package config

import (
	"fmt"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-rest-api-boilerplate/pkg/path"
)

func Test_Config(t *testing.T) {
	var projectBasePath = path.GetProjectBasePath()

	t.Run("should create new config instance and return config instance", func(t *testing.T) {
		testAppEnvironment := "test"
		testConfigFields := getTestConfigFields()
		configInstance := Init(testConfigFields, testAppEnvironment)

		expected := &config{}

		assert.IsType(t, expected, configInstance)
	})

	t.Run("should return app environment", func(t *testing.T) {
		testConfigInstance := createTestConfigInstance()
		getAppEnvironment := testConfigInstance.GetAppEnvironment()

		assert.Equal(t, "test", getAppEnvironment)
	})

	t.Run("should handle config path with app environment variable and return config path", func(t *testing.T) {
		getConfigPath, err := GetConfigPath("test")
		assert.Nil(t, err)

		expectedPath := fmt.Sprintf("%s/pkg/config/testdata/config.yaml", projectBasePath)

		assert.Equal(t, expectedPath, getConfigPath)
	})

	t.Run("should read yaml files and return marshalled config", func(t *testing.T) {
		testPath := fmt.Sprintf("%s/pkg/config/testdata/config.yaml", projectBasePath)

		configFields, err := ReadConfig(testPath)

		assert.Nil(t, err)
		assert.NotNil(t, configFields)

		testConfigFields := getTestConfigFields()
		assert.Equal(t, testConfigFields, configFields)
	})
}

func getTestConfigFields() configmodel.Fields {
	return configmodel.Fields{
		Server: configmodel.Server{
			Port: 1234,
		},
		RPCServer: configmodel.RPCServer{
			Port: 5678,
		},
	}
}

func createTestConfigInstance() Config {
	testAppEnvironment := "test"
	testConfigFields := getTestConfigFields()
	return Init(testConfigFields, testAppEnvironment)
}
