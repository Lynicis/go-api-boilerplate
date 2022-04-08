package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/logger"
)

// Server methods for getter and start http server with fiber
type Server interface {
	GetFiberInstance() *fiber.App
	Start()
}

type server struct {
	config configmodel.Server
	fiber  *fiber.App
	logger logger.Logger
}

// NewServer function is create new server instance
func NewServer(serverConfig configmodel.Server, logger logger.Logger) Server {
	fiberInstance := fiber.New()

	return &server{
		config: serverConfig,
		fiber:  fiberInstance,
		logger: logger,
	}
}

func (server *server) Start() {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		_ = server.fiber.Shutdown()
	}()

	serverAddress := fmt.Sprintf(":%d", server.config.Port)
	err := server.fiber.Listen(serverAddress)
	if err != nil {
		server.logger.Fatalf("something wrong while server starting", zap.Error(err))
	}
}

func (server *server) GetFiberInstance() *fiber.App {
	return server.fiber
}
