package sitehandler

import (
	"strings"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateSiteInput struct {
	Url string `json:"url" validate:"required,url"`
}

func (h *SiteHandler) createSite(c *fiber.Ctx) error {
	authorized := utils.AuthorizeRoute(c)
	if !authorized {
		return utils.UnauthorizedResponse(c)
	}
	auth := utils.GetAuthInfo(c)

	input := new(CreateSiteInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid site url"))
	}

	input.Url = strings.ToLower(input.Url)

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid site url"))
	}

	newSite := models.NewSite(input.Url, auth.User)

	h.DB.Create(&newSite)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"site":    newSite,
	})
}
