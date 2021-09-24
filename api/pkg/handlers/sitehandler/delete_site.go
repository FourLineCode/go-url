package sitehandler

import (
	"errors"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *SiteHandler) deleteSite(c *fiber.Ctx) error {
	authorized := utils.AuthorizeRoute(c)
	if !authorized {
		return utils.UnauthorizedResponse(c)
	}
	auth := utils.GetAuthInfo(c)

	key := c.Params("key")
	var site models.Site

	result := h.DB.Where("key = ?", key).First(&site)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Url not found with given key"))
	}

	if site.UserID != auth.User.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewError("You are not authorized to update this site"))
	}

	deleteResult := h.DB.Delete(&site)
	if deleteResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Something went wrong"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
