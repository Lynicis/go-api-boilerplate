package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/path"
)

// Config Provide all config methods
type Config interface {
	GetServerConfig() configmodel.Server
	GetRPCConfig() configmodel.RPCServer
	GetAppEnvironment() string
}

type config struct {
	environment string
	fields      configmodel.Fields
}

// Init Initialize config package
func Init(configFields configmodel.Fields, appEnvironment string) Config {
	return &config{
		environment: appEnvironment,
		fields:      configFields,
	}
}

// GetConfigPath Handle environment and return config path
func GetConfigPath(environment string) (string, error) {
	var configPath string

	projectPath := path.GetProjectBasePath()
	baseConfigPath := fmt.Sprintf("%s/config/", projectPath)

	switch environment {
	case "local":
		configPath = fmt.Sprintf("%s/development.yaml", baseConfigPath)
	case "prod":
		configPath = fmt.Sprintf("%s/production.yaml", baseConfigPath)
	case "test":
		configPath = fmt.Sprintf("%s/pkg/config/testdata/config.yaml", projectPath)
	default:
		return "", fmt.Errorf("you must define valid app environment in environment variables")
	}

	return configPath, nil
}

// ReadConfig Read config file and return marshalled config
func ReadConfig(configPath string) (configmodel.Fields, error) {
	unmarshalledConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return configmodel.Fields{}, err
	}

	var configFields configmodel.Fields
	_ = yaml.Unmarshal(unmarshalledConfig, &configFields)

	return configFields, nil
}

func (config *config) GetServerConfig() configmodel.Server {
	return config.fields.Server
}

func (config *config) GetRPCConfig() configmodel.RPCServer {
	return config.fields.RPCServer
}

func (config *config) GetAppEnvironment() string {
	return config.environment
}
