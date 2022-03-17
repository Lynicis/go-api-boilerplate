package config

import (
	"testing"

	"turkic-mythology-gateway/pkg/config/model"

	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	t.Run("should create new config instance and return config", func(t *testing.T) {
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
