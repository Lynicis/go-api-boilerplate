package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"turkic-mythology/pkg/config/model"
)

type Config interface {
	GetServerConfig() configmodel.Server
	GetConfigPath() string
	GetAppEnvironment() string
}

type config struct {
	environment string
	fields      configmodel.Fields
}

func Init(configFields configmodel.Fields, appEnvironment string) Config {
	return &config{
		environment: appEnvironment,
		fields:      configFields,
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

func (c *config) GetConfigPath() string {
	var configPath string

	getPath, _ := os.Getwd()
	projectPath := filepath.Base(fmt.Sprintf("../../%s", getPath))

	baseConfigPath := filepath.Join(projectPath, "/config")
	configPathForPath := "pkg/config/testdata/config.yaml"

	switch c.environment {
	case "development":
		configPath = filepath.Join(baseConfigPath, "development.yaml")
	case "production":
		configPath = filepath.Join(baseConfigPath, "production.yaml")
	case "test":
		configPath = configPathForPath
	}

	return configPath
}

func (c *config) GetAppEnvironment() string {
	return c.environment
}
