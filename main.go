package main

import (
	"log"
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

	var handlers []server.Handler
	srv := server.NewServer(cfg, handlers)
	app := srv.GetFiberInstance()

	log := logger.NewLogger()
	app.Use(logger.Middleware(log))

	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
