package main

import (
	"log"
	"turkic-mythology-gateway/cmd/server"
	"turkic-mythology-gateway/pkg/config"
)

func main() {
	configPath := "config/config.yaml"
	configFields, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	configInstance := config.Init(configFields)
	serverConfig := configInstance.GetServerConfig()

	srv := server.NewGatewayServer(serverConfig.Port)
	err = srv.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
