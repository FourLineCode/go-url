package utils

import "github.com/gofiber/fiber/v2"

func AuthorizeRoute(c *fiber.Ctx) bool {
	auth := GetAuthInfo(c)

	return auth.Authorized
}

func UnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(NewError("You are not authorized to make this request"))
}
