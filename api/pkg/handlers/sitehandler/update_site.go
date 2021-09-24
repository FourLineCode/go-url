package sitehandler

import (
	"errors"
	"strings"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateSiteInput struct {
	Url string `json:"url" validate:"required,url"`
}

func (h *SiteHandler) updateSite(c *fiber.Ctx) error {
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

	input := new(CreateSiteInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid site url"))
	}
	input.Url = strings.ToLower(input.Url)

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid site url"))
	}

	site.Url = input.Url
	saveResult := h.DB.Save(&site)
	if saveResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Something went wrong"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"site":    site,
	})
}
