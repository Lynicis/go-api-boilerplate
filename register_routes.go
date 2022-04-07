package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/internal/health_handler"
)

func RegisterRoutes(fiber *fiber.App) {
	fiber.Get("/health", health_handler.GetStatus)
}
