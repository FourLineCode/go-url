package sitehandler

import (
	"errors"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *SiteHandler) getSiteUrl(c *fiber.Ctx) error {
	key := c.Params("key")

	var site models.Site

	result := h.DB.Where("key = ?", key).First(&site)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Url not found with given key"))
	}

	return c.Status(fiber.StatusOK).JSON(site)
}
