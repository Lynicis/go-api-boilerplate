package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"turkic-mythology-gateway/pkg/config"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
}

type server struct {
	config config.Config
	fiber  *fiber.App
}

func NewGatewayServer(configInstance config.Config) Server {
	fiberInstance := fiber.New()

	return &server{
		config: configInstance,
		fiber:  fiberInstance,
	}
}

func (s *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt)

	go func() {
		<-shutdownChannel
		err := s.fiber.Shutdown()
		if err != nil {
			log.Fatalf("Error while shutting down the s: %v", err)
		}
	}()

	GetServerPort := s.config.GetServerConfig().Port
	return s.fiber.Listen(GetServerPort)
}

func (s *server) GetFiberInstance() *fiber.App {
	return s.fiber
}
