//go:build unit

package config

import (
	"fmt"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
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
	testPath := getTestPath()
	configFields, err := ReadConfig(testPath)

	assert.Nil(t, err)

	testConfigInstance := Init(configFields, "test")
	testServerConfig := testConfigInstance.GetServerConfig()

	expectedServerConfig := configmodel.Server{}

	assert.NotNil(t, testServerConfig)
	assert.IsType(t, testServerConfig, expectedServerConfig)
}

func TestConfig_GetRPCConfig(t *testing.T) {
	testPath := getTestPath()
	configFields, err := ReadConfig(testPath)

	assert.Nil(t, err)

	testConfigInstance := Init(configFields, "test")
	testRPCConfig := testConfigInstance.GetRPCConfig()

	expectedRPCConfig := configmodel.RPCServer{}

	assert.NotNil(t, testRPCConfig)
	assert.IsType(t, testRPCConfig, expectedRPCConfig)
}

func TestGetConfigPath(t *testing.T) {
	getLocalConfigPath, err := GetConfigPath("local")
	assert.Nil(t, err)
	assert.Regexp(t, ".yaml$", getLocalConfigPath)

	getProductionConfigPath, err := GetConfigPath("prod")
	assert.Nil(t, err)
	assert.Regexp(t, "yaml$", getProductionConfigPath)

	getTestConfigPath, err := GetConfigPath("test")
	assert.Nil(t, err)
	assert.Regexp(t, ".yaml$", getTestConfigPath)
}

func TestReadConfig(t *testing.T) {
	testPath := getTestPath()
	readConfig, err := ReadConfig(testPath)

	assert.Nil(t, err)

	expectedConfig := configmodel.Fields{}

	assert.NotNil(t, readConfig)
	assert.IsType(t, readConfig, expectedConfig)

	testConfigFields := getTestConfigFields()
	assert.Equal(t, testConfigFields, readConfig)
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

func getTestPath() string {
	var projectBasePath = path.GetProjectBasePath()
	return fmt.Sprintf("%s/pkg/config/testdata/config.yaml", projectBasePath)
}
