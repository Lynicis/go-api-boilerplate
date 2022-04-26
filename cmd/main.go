package main

import (
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	logger_middleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"go-rest-api-boilerplate/internal/health"
	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/http_server"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/rpc_server"
)

func main() {
	var err error

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)

	log := logger.CreateLogger()

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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
	log.Infof("App environment: %s", configInstance.GetAppEnvironment())

	serverConfig := configInstance.GetServerConfig()
	httpServerInstance := http_server.NewHTTPServer(serverConfig.HTTP)
	rpcServerInstance := rpc_server.NewRPCServer(serverConfig.RPC)

	fiberInstance := httpServerInstance.GetFiberInstance()
	grpcInstance := rpcServerInstance.GetRPCServer()

	registerMiddlewares(fiberInstance)
	registerRoutes(fiberInstance)

	health.RegisterHealthCheckService(grpcInstance)

	go func() {
		err = httpServerInstance.Start()
		if err != nil {
			log.Fatal(err)
		}

		waitGroup.Done()
	}()

	go func() {
		err = rpcServerInstance.Start()
		if err != nil {
			log.Fatal(err)
		}

		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func registerMiddlewares(httpServer *fiber.App) {
	httpServer.Use(recover.New())
	httpServer.Use(logger_middleware.New())
	httpServer.Use(compress.New())
}

func registerRoutes(httpServer *fiber.App) {
	httpServer.Get("/health", health.GetStatus)
}
