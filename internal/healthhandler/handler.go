package healthhandler

import (
	"github.com/gofiber/fiber/v2"

	healthmodel "go-rest-api-boilerplate/internal/healthhandler/model"
)

// GetStatus handler of health check endpoint
func GetStatus(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(
		healthmodel.HealthEndpoint{
			Status: "OK",
		},
	)
}
