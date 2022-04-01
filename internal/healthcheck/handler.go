package healthcheck

import "github.com/gofiber/fiber/v2"

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}
