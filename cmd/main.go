package main

import (
	"os"

	"go.uber.org/zap"

	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/server"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	getEnvironment := os.Getenv("APP_ENV")
	configPath, err := config.GetConfigPath(getEnvironment)
	if err != nil {
		logger.Fatal("Failed to read environment variable", zap.Error(err))
	}

	readConfig, err := config.ReadConfig(configPath)
	if err != nil {
		logger.Fatal("Failed to read config file", zap.Error(err))
	}

	configInstance := config.Init(readConfig, getEnvironment)

	serverConfig := configInstance.GetServerConfig()
	serverInstance := server.NewServer(serverConfig)
	fiberInstance := serverInstance.GetFiberInstance()

	RegisterRoutes(fiberInstance)

	err = serverInstance.Start()
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
