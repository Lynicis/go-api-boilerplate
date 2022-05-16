package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerHTTPMiddlewares(httpServer *fiber.App) {
	httpServer.Use(recover.New())
	httpServer.Use(logger.New())
	httpServer.Use(compress.New())
}
