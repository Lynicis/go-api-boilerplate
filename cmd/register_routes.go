package main

import (
	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/internal/healthcheck"
)

func RegisterRoutes(fiberInstance *fiber.App) fiber.Router {
	return fiberInstance.Get("/health", healthcheck.GetStatus)
}
