package server

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!\n")
	})

	return app
}
