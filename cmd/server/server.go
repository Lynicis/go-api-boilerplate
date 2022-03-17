package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

type Server interface {
	GetFiberInstance() *fiber.App
	Start() error
}

type server struct {
	port  string
	fiber *fiber.App
}

func NewGatewayServer(port string) Server {
	fiberInstance := fiber.New()

	return &server{
		port:  port,
		fiber: fiberInstance,
	}
}

func (s *server) GetFiberInstance() *fiber.App {
	return s.fiber
}

func (s *server) Start() error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt)

	go func() {
		<-shutdownChannel
		err := s.fiber.Shutdown()
		if err != nil {
			log.Fatalf("Error while shutting down the server: %v", err)
		}
	}()

	return s.fiber.Listen(s.port)
}
