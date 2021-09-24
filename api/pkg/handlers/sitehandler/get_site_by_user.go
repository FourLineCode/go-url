package sitehandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *SiteHandler) getSiteByUser(c *fiber.Ctx) error {
	authorized := utils.AuthorizeRoute(c)
	if !authorized {
		return utils.UnauthorizedResponse(c)
	}
	auth := utils.GetAuthInfo(c)

	var sites []models.Site

	h.DB.Where("user_id = ?", auth.User.ID).First(&sites)

	return c.Status(fiber.StatusOK).JSON(sites)
}
