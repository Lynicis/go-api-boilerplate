package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/pkg/config"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
	Stop() error
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
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		err := server.fiber.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	serverConfig := fmt.Sprintf(":%d", server.config.GetServerConfig().Port)
	return server.fiber.Listen(serverConfig)
}

func (server *server) Stop() error {
	return server.fiber.Shutdown()
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
