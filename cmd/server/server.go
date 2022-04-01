package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/pkg/config"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
}

type server struct {
	config config.Config
	fiber  *fiber.App
}

func NewServer(configInstance config.Config) Server {
	fiberInstance := fiber.New()

	return &server{
		config: configInstance,
		fiber:  fiberInstance,
	}
}

func (server *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt)

	go func() {
		<-shutdownChannel
		err := server.fiber.Shutdown()
		if err != nil {
			log.Fatalf("Error while shutting down the server: %v", err)
		}
	}()

	serverConfig := fmt.Sprintf(":%d", server.config.GetServerConfig().Port)
	return server.fiber.Listen(serverConfig)
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
