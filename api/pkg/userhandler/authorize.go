package userhandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) authorize(c *fiber.Ctx) error {
	token := c.Cookies("auth")

	valid, claims, err := utils.ValidateJWTSignedToken(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Error validating auth token"))
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewError("Access denied"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authorized": true,
		"user":       claims,
	})
}
