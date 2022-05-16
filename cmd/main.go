package main

import (
	"os"
	"path"

	"github.com/joho/godotenv"

	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/http_server"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/project_path"
	"go-rest-api-boilerplate/pkg/rpc_server"
)

func main() {
	var err error

	log := logger.CreateLogger()

	rootDirectory := project_path.GetRootDirectory()
	dotenvPath := path.Join(rootDirectory, ".env")

	err = godotenv.Load(dotenvPath)
	if err != nil {
		log.Fatal(err)
	}

	appEnvironment := os.Getenv("APP_ENV")
	configPath, err := config.GetConfigPath(appEnvironment)
	if err != nil {
		log.Fatal(err)
	}

	readConfig, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	configInstance := config.Init(readConfig, appEnvironment)

	serverConfig := configInstance.GetServerConfig()
	httpServerInstance := http_server.NewHTTPServer(serverConfig.HTTP)
	rpcServerInstance := rpc_server.NewRPCServer(serverConfig.RPC)

	fiberInstance := httpServerInstance.GetFiberInstance()
	grpcInstance := rpcServerInstance.GetRPCServer()

	registerHTTPMiddlewares(fiberInstance)
	registerHTTPRoutes(fiberInstance)
	registerRPCHandlers(grpcInstance)
	startServer(
		httpServerInstance,
		rpcServerInstance,
		log,
	)
}
