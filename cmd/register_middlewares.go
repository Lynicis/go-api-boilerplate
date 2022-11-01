package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerMiddlewares(server *fiber.App) {
	server.Use(recover.New())
	server.Use(logger.New())
	server.Use(compress.New())
}
