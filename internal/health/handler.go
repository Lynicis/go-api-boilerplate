package health

import (
	"github.com/gofiber/fiber/v2"

	health_handler_model "go-rest-api-boilerplate/internal/health/model"
)

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.
		Status(200).
		JSON(
			health_handler_model.HealthEndpoint{
				Status: "OK",
			},
		)
}
