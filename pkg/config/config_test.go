//go:build unit

package config

import (
	"fmt"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-rest-api-boilerplate/pkg/path"
)

func TestInit(t *testing.T) {
	testAppEnvironment := "test"
	testConfigFields := getTestConfigFields()
	configInstance := Init(testConfigFields, testAppEnvironment)

	expected := &config{}

	assert.IsType(t, expected, configInstance)
}

func TestConfig_GetAppEnvironment(t *testing.T) {
	testConfigInstance := createTestConfigInstance()
	getAppEnvironment := testConfigInstance.GetAppEnvironment()

	assert.Equal(t, "test", getAppEnvironment)
}

func TestConfig_GetServerConfig(t *testing.T) {
	testPath := getTestPath("config.yaml")
	configFields, err := ReadConfig(testPath)

	assert.Nil(t, err)

	testConfigInstance := Init(configFields, "test")
	testServerConfig := testConfigInstance.GetServerConfig()

	expectedServerConfig := configmodel.Server{}

	assert.NotNil(t, testServerConfig)
	assert.IsType(t, testServerConfig, expectedServerConfig)
}

func TestConfig_GetRPCConfig(t *testing.T) {
	testPath := getTestPath("config.yaml")
	configFields, err := ReadConfig(testPath)

	assert.Nil(t, err)

	testConfigInstance := Init(configFields, "test")
	testRPCConfig := testConfigInstance.GetRPCConfig()

	expectedRPCConfig := configmodel.RPCServer{}

	assert.NotNil(t, testRPCConfig)
	assert.IsType(t, testRPCConfig, expectedRPCConfig)
}

func TestGetConfigPath(t *testing.T) {
	t.Run("should return local config path", func(t *testing.T) {
		getLocalConfigPath, err := GetConfigPath("local")
		assert.Nil(t, err)
		assert.Regexp(t, ".yaml$", getLocalConfigPath)
	})

	t.Run("should return prod config path", func(t *testing.T) {
		getProductionConfigPath, err := GetConfigPath("prod")
		assert.Nil(t, err)
		assert.Regexp(t, "yaml$", getProductionConfigPath)
	})

	t.Run("should return test config path", func(t *testing.T) {
		getTestConfigPath, err := GetConfigPath("test")
		assert.Nil(t, err)
		assert.Regexp(t, ".yaml$", getTestConfigPath)
	})

	t.Run("should return error with wrong environment variable", func(t *testing.T) {
		_, err := GetConfigPath("invalid")
		assert.NotNil(t, err)
	})
}

func TestReadConfig(t *testing.T) {
	t.Run("should read config file and return config model", func(t *testing.T) {
		testPath := getTestPath("config.yaml")
		readConfig, err := ReadConfig(testPath)

		expectedConfig := configmodel.Fields{}

		assert.Nil(t, err)
		assert.IsType(t, readConfig, expectedConfig)
	})

	t.Run("should if config file is not exist return error", func(t *testing.T) {
		testPath := getTestPath("invalid-config.yaml")
		_, err := ReadConfig(testPath)

		expected := &fs.PathError{}

		assert.NotNil(t, err)
		assert.IsType(t, expected, err)
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

func getTestPath(filename string) string {
	var projectBasePath = path.GetProjectBasePath()
	return fmt.Sprintf("%s/pkg/config/testdata/%s", projectBasePath, filename)
}
