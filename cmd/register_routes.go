package main

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"

	example_rpc "go-rest-api-boilerplate/internal/example-rpc"
	"go-rest-api-boilerplate/internal/health"
)

func registerHTTPRoutes(httpServer *fiber.App) {
	httpServer.Get("/health", health.GetStatus)
}

func registerRPCHandlers(rpcServer *grpc.Server) {
	example_rpc.RegisterHandler(rpcServer)
}
