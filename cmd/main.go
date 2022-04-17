package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	loggermiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go-rest-api-boilerplate/internal/healthhandler"
	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/server"
)

func main() {
	var err error

	log := logger.CreateLogger()

	getEnvironment := os.Getenv("APP_ENV")
	configPath, err := config.GetConfigPath(getEnvironment)
	if err != nil {
		log.Fatal(err)
	}

	readConfig, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	configInstance := config.Init(readConfig, getEnvironment)

	serverConfig := configInstance.GetServerConfig()
	serverInstance := server.NewServer(serverConfig)

	fiberInstance := serverInstance.GetFiberInstance()
	registerMiddlewares(fiberInstance)
	registerRoutes(fiberInstance)

	err = serverInstance.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func registerMiddlewares(server *fiber.App) {
	server.Use(recover.New())
	server.Use(loggermiddleware.New())
	server.Use(compress.New())
}

func registerRoutes(server *fiber.App) {
	server.Get("/health", healthhandler.GetStatus)
}
