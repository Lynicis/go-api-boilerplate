package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
	Shutdown() error
}

type server struct {
	config configmodel.HTTPServer
	fiber  *fiber.App
}

func NewHTTPServer(serverConfig configmodel.HTTPServer) Server {
	fiberConfig := fiber.Config{
		DisableStartupMessage: true,
	}

	fiberInstance := fiber.New(fiberConfig)

	return &server{
		config: serverConfig,
		fiber:  fiberInstance,
	}
}

func (server *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		_ = server.fiber.Shutdown()
	}()

	serverAddress := fmt.Sprintf(":%d", server.config.Port)
	return server.fiber.Listen(serverAddress)
}

func (server *server) Shutdown() error {
	return server.fiber.Shutdown()
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
