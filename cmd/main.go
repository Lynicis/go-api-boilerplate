package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"go-rest-api-boilerplate/pkg/config"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/project"
	"go-rest-api-boilerplate/pkg/server"
)

func main() {
	var err error

	log := logger.CreateLogger()

	isAtRemote := os.Getenv(config.IsAtRemote)
	if isAtRemote == "" {
		rootDirectory := project.GetRootDirectory()
		dotenvPath := filepath.Join(rootDirectory, ".env")
		err = godotenv.Load(dotenvPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	var cfg config.Config
	cfg, err = config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	httpServer := server.NewServer(cfg)
	fiberInstance := httpServer.GetFiberInstance()

	registerMiddlewares(fiberInstance)
	registerRoutes(fiberInstance)

	err = httpServer.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Application Running...")
}
