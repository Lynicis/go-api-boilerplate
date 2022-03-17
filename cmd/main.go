package main

import (
	"log"
	"turkic-mythology-gateway/cmd/server"
	"turkic-mythology-gateway/internal/healtcheck"
	"turkic-mythology-gateway/pkg/config"
)

func main() {
	configPath := "config/config.yaml"
	configFields, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	configInstance := config.Init(configFields)
	gatewayServer := server.NewGatewayServer(configInstance)
	gatewayFiberInstance := gatewayServer.GetFiberInstance()

	gatewayFiberInstance.Get("/health", healtcheck.GetStatus)

	err = gatewayServer.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
