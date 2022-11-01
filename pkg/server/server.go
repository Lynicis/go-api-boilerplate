package server

import (
	"fmt"
	"go-rest-api-boilerplate/pkg/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
	Shutdown() error
}

type server struct {
	serverPort string
	fiber      *fiber.App
}

func NewServer(config config.Config) Server {
	fiberConfig := fiber.Config{
		DisableStartupMessage: true,
	}

	fiberInstance := fiber.New(fiberConfig)
	serverPort := config.GetServerPort()

	return &server{
		serverPort: serverPort,
		fiber:      fiberInstance,
	}
}

func (server *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		_ = server.fiber.Shutdown()
	}()

	serverAddress := fmt.Sprintf(":%d", server.serverPort)
	return server.fiber.Listen(serverAddress)
}

func (server *server) Shutdown() error {
	return server.fiber.Shutdown()
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
