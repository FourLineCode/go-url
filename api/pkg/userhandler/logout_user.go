package userhandler

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) logoutUser(c *fiber.Ctx) error {
	c.ClearCookie("auth")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
