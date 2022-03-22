//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"turkic-mythology/pkg/config/model"
)

func Test_Config(t *testing.T) {
	t.Run("should create new config instance and return config instance", func(t *testing.T) {
		testAppEnvironment := "test"
		testConfigFields := getTestConfigFields()
		configInstance := Init(testConfigFields, testAppEnvironment)

		expected := &config{
			environment: testAppEnvironment,
			fields:      testConfigFields,
		}

		assert.Equal(t, expected, configInstance)
	})

	t.Run("should return app environment", func(t *testing.T) {
		testConfigInstance := createTestConfigInstance()
		getAppEnvironment := testConfigInstance.GetAppEnvironment()

		assert.Equal(t, "test", getAppEnvironment)
	})

	t.Run("should handle config path with app environment variable and return config path", func(t *testing.T) {
		testConfigInstance := createTestConfigInstance()
		getConfigPath := testConfigInstance.GetConfigPath()

		expectedPath := "pkg/config/testdata/development.yaml"

		assert.Equal(t, expectedPath, getConfigPath)
	})

	t.Run("should read yaml files and return marshalled config", func(t *testing.T) {
		testPath := "./testdata/development.yaml"

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
			Port: ":1234",
		},
	}
}

func createTestConfigInstance() Config {
	testAppEnvironment := "test"
	testConfigFields := getTestConfigFields()
	configInstance := Init(testConfigFields, testAppEnvironment)

	return configInstance
}
