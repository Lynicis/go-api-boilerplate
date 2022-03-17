package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"turkic-mythology-gateway/pkg/config/model"
)

type Config interface {
	GetServerConfig() configmodel.Server
}

type config struct {
	fields configmodel.Fields
}

func Init(configFields configmodel.Fields) Config {
	return &config{
		fields: configFields,
	}
}

func ReadConfig(configPath string) (configmodel.Fields, error) {
	unmarshalledConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return configmodel.Fields{}, err
	}

	var configFields configmodel.Fields
	err = yaml.Unmarshal(unmarshalledConfig, &configFields)

	return configFields, err
}

func (c *config) GetServerConfig() configmodel.Server {
	return c.fields.Server
}
