package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

// Server This interface for HTTP Server
type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
}

type server struct {
	config configmodel.Server
	fiber  *fiber.App
}

// NewServer Create New HTTP Server with Fiber
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
		err := server.fiber.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	serverAddress := fmt.Sprintf(":%d", server.config.Port)
	return server.fiber.Listen(serverAddress)
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
