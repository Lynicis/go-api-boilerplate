package main

import (
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

	configPath := "config/config.yaml"
	configFields, err := config.ReadConfig(configPath)
	if err != nil {
		logger.Fatal("Failed to reading config file", zap.Error(err))
	}

	configInstance := config.Init(configFields)
	newServer := server.NewServer(configInstance)
	fiberInstance := newServer.GetFiberInstance()

	fiberInstance.Get("/health", healtcheck.GetStatus)

	err = newServer.Start()
	if err != nil {
		logger.Fatal("Failed to start newServer", zap.Error(err))
	}
}
