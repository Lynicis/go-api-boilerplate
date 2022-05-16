//go:build unit

package config

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/project_path"
)

func Test_Init(t *testing.T) {
	testAppEnvironment := "test"
	testConfigFields := getTestConfigFields()
	configInstance := Init(testConfigFields, testAppEnvironment)

	expected := &config{}

	assert.IsType(t, expected, configInstance)
}

func Test_GetConfigPath(t *testing.T) {
	t.Run("should return local config path", func(t *testing.T) {
		getLocalConfigPath, err := GetConfigPath("local")

		assert.NoError(t, err)
		assert.Regexp(t, ".yaml$", getLocalConfigPath)
	})

	t.Run("should return prod config path", func(t *testing.T) {
		getProductionConfigPath, err := GetConfigPath("prod")

		assert.NoError(t, err)
		assert.Regexp(t, "yaml$", getProductionConfigPath)
	})

	t.Run("should return test config path", func(t *testing.T) {
		getTestConfigPath, err := GetConfigPath("test")

		assert.NoError(t, err)
		assert.Regexp(t, ".yaml$", getTestConfigPath)
	})

	t.Run("should return error with wrong environment variable", func(t *testing.T) {
		_, err := GetConfigPath("invalid")

		assert.Error(t, err)
	})
}

func Test_ReadConfig(t *testing.T) {
	t.Run("should read config file and return config model", func(t *testing.T) {
		testPath := getTestPath("config.yaml")
		readConfig, err := ReadConfig(testPath)

		expectedConfig := configmodel.Fields{}

		assert.NoError(t, err)
		assert.IsType(t, readConfig, expectedConfig)
	})

	t.Run("should if config file is not exist return error", func(t *testing.T) {
		testPath := getTestPath("invalid-config.yaml")
		_, err := ReadConfig(testPath)

		expectedError := &fs.PathError{}

		assert.Error(t, err)
		assert.IsType(t, expectedError, err)
	})

	t.Run("should config file have type error return error", func(t *testing.T) {
		testPath := getTestPath("type-error.yaml")
		_, err := ReadConfig(testPath)

		expectedError := &yaml.TypeError{}

		assert.Error(t, err)
		assert.IsType(t, expectedError, err)
	})
}

func TestConfig_GetAppEnvironment(t *testing.T) {
	testConfigInstance := createTestConfigInstance()

	getAppEnvironment := testConfigInstance.GetAppEnvironment()

	assert.Equal(t, "test", getAppEnvironment)
}

func TestConfig_GetServerConfig(t *testing.T) {
	testPath := getTestPath("config.yaml")
	configFields, err := ReadConfig(testPath)

	testConfigInstance := Init(configFields, "test")
	testServerConfig := testConfigInstance.GetServerConfig()

	expectedServerConfig := configmodel.Server{}

	assert.NoError(t, err)
	assert.NotNil(t, testServerConfig)
	assert.IsType(t, testServerConfig, expectedServerConfig)
}

/*
	Helpers for test
*/
func getTestConfigFields() configmodel.Fields {
	return configmodel.Fields{
		Server: configmodel.Server{
			HTTP: configmodel.HTTPServer{
				Port: 9090,
			},
			RPC: configmodel.RPCServer{
				Port: 9091,
			},
		},
	}
}

func createTestConfigInstance() Config {
	testAppEnvironment := "test"
	testConfigFields := getTestConfigFields()
	return Init(testConfigFields, testAppEnvironment)
}

func getTestPath(filename string) string {
	var projectBasePath = project_path.GetRootDirectory()

	return fmt.Sprintf(
		"%s/pkg/config/testdata/%s",
		projectBasePath,
		filename,
	)
}
