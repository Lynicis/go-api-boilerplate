package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-rest-api-boilerplate/internal/healthhandler"
)

// RegisterRoutes register all http server endpoints
func RegisterRoutes(fiber *fiber.App) {
	fiber.Get("/health", healthhandler.GetStatus)
}
