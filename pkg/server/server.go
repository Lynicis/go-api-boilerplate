package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

// Server http server with fiber framework domain
type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
	Shutdown()
}

type server struct {
	config configmodel.Server
	fiber  *fiber.App
}

// NewServer create new server instance
func NewServer(serverConfig configmodel.Server) Server {
	fiberInstance := fiber.New()

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

func (server *server) Shutdown() {
	_ = server.fiber.Shutdown()
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
