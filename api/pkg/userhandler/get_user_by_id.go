package userhandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) getUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	results := h.DB.First(&user, id)
	if results.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError(results.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
