package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/internal/healthhandler"
	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/server"
)

func main() {
	var err error

	log := logger.CreateLogger()
	defer func(log logger.Logger) {
		err = log.Sync()
		if err != nil {
			panic(err)
		}
	}(log)

	getEnvironment := os.Getenv("APP_ENV")
	configPath, err := config.GetConfigPath(getEnvironment)
	if err != nil {
		log.Fatalf("failed to read environment variable: %s", err)
	}

	readConfig, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to read config file %s", err)
	}

	configInstance := config.Init(readConfig, getEnvironment)

	serverConfig := configInstance.GetServerConfig()
	serverInstance := server.NewServer(serverConfig)

	fiberInstance := serverInstance.GetFiberInstance()
	registerRoutes(fiberInstance)

	err = serverInstance.Start()
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}

func registerRoutes(server *fiber.App) {
	server.Get("/health", healthhandler.GetStatus)
}
