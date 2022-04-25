package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/path"
)

type Config interface {
	GetServerConfig() configmodel.Server
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

func GetConfigPath(environment string) (string, error) {
	var configPath string

	projectPath := path.GetRootDirectory()
	baseConfigPath := fmt.Sprintf("%s/config/", projectPath)

	switch environment {
	case "local":
		configPath = fmt.Sprintf("%s/development.yaml", baseConfigPath)
	case "prod":
		configPath = fmt.Sprintf("%s/production.yaml", baseConfigPath)
	case "test":
		configPath = fmt.Sprintf("%s/pkg/config/testdata/config.yaml", projectPath)
	default:
		return configPath,
			fmt.Errorf("you must define valid app environment in environment variables")
	}

	return configPath, nil
}

func ReadConfig(configPath string) (configmodel.Fields, error) {
	var err error

	unmarshalledConfig, err := ioutil.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return configmodel.Fields{}, err
	}

	var configFields configmodel.Fields
	err = yaml.Unmarshal(unmarshalledConfig, &configFields)
	if err != nil {
		return configmodel.Fields{}, err
	}

	return configFields, nil
}

func (config *config) GetServerConfig() configmodel.Server {
	return config.fields.Server
}

func (config *config) GetAppEnvironment() string {
	return config.environment
}
