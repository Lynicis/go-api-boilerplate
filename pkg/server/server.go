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
	serverConfig := server.config.GetServerConfig()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		err := server.fiber.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	return server.fiber.Listen(fmt.Sprintf(":%d", serverConfig.Port))
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
