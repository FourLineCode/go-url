package sitehandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *SiteHandler) getSites(c *fiber.Ctx) error {
	authorized := utils.AuthorizeRoute(c)
	if !authorized {
		return utils.UnauthorizedResponse(c)
	}

	var sites []models.Site
	results := h.DB.Joins("User").Find(&sites)
	if results.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError(results.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(sites)
}
