package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/project_path"
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

	projectPath := project_path.GetRootDirectory()
	baseConfigPath := path.Join(projectPath, "config")

	switch environment {
	case "local":
		configPath = path.Join(baseConfigPath, "development.yaml")
	case "prod":
		configPath = path.Join(baseConfigPath, "production.yaml")
	case "test":
		configPath = path.Join(projectPath, "pkg/config/testdata/config.yaml")
	default:
		return configPath,
			fmt.Errorf("you must define valid app environment in environment variables")
	}

	return configPath, nil
}

func ReadConfig(configPath string) (configmodel.Fields, error) {
	var err error

	unmarshalledConfig, err := ioutil.ReadFile(
		path.Clean(configPath),
	)
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
