package main

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-boilerplate/internal/health"
)

func registerRoutes(server *fiber.App) {
	server.Get("/health", health.GetStatus)
}
