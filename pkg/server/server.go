package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-rest-api-boilerplate/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterRoutes(ctx *fiber.App)
}

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
	Shutdown() error
}

type server struct {
	serverPort string
	fiber      *fiber.App
	handlers   []Handler
}

func NewServer(config config.Config, handlers []Handler) Server {
	fiberInstance := fiber.New()
	serverPort := config.GetServerPort()

	return &server{
		serverPort: serverPort,
		fiber:      fiberInstance,
		handlers:   handlers,
	}
}

func (server *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		_ = server.fiber.Shutdown()
	}()

	registerRoutes(server.fiber, server.handlers)

	serverAddress := fmt.Sprintf(":%s", server.serverPort)
	return server.fiber.Listen(serverAddress)
}

func (server *server) Shutdown() error {
	return server.fiber.Shutdown()
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}

func registerRoutes(app *fiber.App, handlers []Handler) {
	for _, handler := range handlers {
		handler.RegisterRoutes(app)
	}
}
