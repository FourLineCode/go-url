package userhandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) authorize(c *fiber.Ctx) error {
	auth := utils.GetAuthInfo(c)

	if !auth.Authorized {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewError("Access denied"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authorized": auth.Authorized,
		"user":       auth.User,
	})
}
