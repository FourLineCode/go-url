package utils

import "github.com/gofiber/fiber/v2"

func NewError(msg string) fiber.Map {
	return fiber.Map{"error": msg}
}
