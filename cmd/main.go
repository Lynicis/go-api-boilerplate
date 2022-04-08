package main

import (
	"os"

	"go.uber.org/zap"

	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/server"
)

func main() {
	log := logger.CreateLogger()
	defer func(log logger.Logger) {
		err := log.Sync()
		if err != nil {
			panic(err)
		}
	}(log)

	getEnvironment := os.Getenv("APP_ENV")
	configPath, err := config.GetConfigPath(getEnvironment)
	if err != nil {
		log.Errorf("Failed to read environment variable", zap.Error(err))
	}

	readConfig, err := config.ReadConfig(configPath)
	if err != nil {
		log.Errorf("Failed to read config file", zap.Error(err))
	}

	configInstance := config.Init(readConfig, getEnvironment)

	serverConfig := configInstance.GetServerConfig()
	serverInstance := server.NewServer(serverConfig, log)

	fiberInstance := serverInstance.GetFiberInstance()
	server.RegisterRoutes(fiberInstance)

	serverInstance.Start()
}
