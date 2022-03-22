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
	gatewayServer := server.NewGatewayServer(configInstance)
	gatewayFiberInstance := gatewayServer.GetFiberInstance()

	gatewayFiberInstance.Get("/health", healtcheck.GetStatus)

	err = gatewayServer.Start()
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
