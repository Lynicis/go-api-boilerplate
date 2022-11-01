package health

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.
		Status(http.StatusOK).
		JSON(
			fiber.Map{
				"status": "OK",
			},
		)
}
