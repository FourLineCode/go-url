package userhandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	authorized := utils.AuthorizeRoute(c)
	if !authorized {
		return utils.UnauthorizedResponse(c)
	}

	var users []models.User
	results := h.DB.Find(&users)
	if results.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError(results.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
