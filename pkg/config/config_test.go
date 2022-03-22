//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"turkic-mythology/pkg/config/model"
)

func Test_Config(t *testing.T) {
	t.Run("should create new config instance and return config instance", func(t *testing.T) {
		testConfigFields := getTestConfigFields()
		configInstance := Init(testConfigFields)

		expected := &config{fields: testConfigFields}

		assert.Equal(t, expected, configInstance)
	})

	t.Run("should read yaml files and return marshalled config", func(t *testing.T) {
		testPath := "./testdata/config.yaml"

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
