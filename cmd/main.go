package main

import (
	"os"

	"go.uber.org/zap"

	"turkic-mythology/cmd/server"
	"turkic-mythology/internal/healtcheck"
	"turkic-mythology/pkg/config"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	configPath := "config/development.yaml"
	configFields, err := config.ReadConfig(configPath)
	if err != nil {
		logger.Fatal("Failed to reading config file", zap.Error(err))
	}

	appEnvironment := os.Getenv("APP_ENVIRONMENT")
	configInstance := config.Init(configFields, appEnvironment)
	newServer := server.NewServer(configInstance)
	fiberInstance := newServer.GetFiberInstance()

	fiberInstance.Get("/health", healtcheck.GetStatus)

	err = newServer.Start()
	if err != nil {
		logger.Fatal("Failed to start newServer", zap.Error(err))
	}
}
