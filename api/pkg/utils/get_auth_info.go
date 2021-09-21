package utils

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type AuthInfo struct {
	Authorized bool
	User       models.User
}

func GetAuthInfo(c *fiber.Ctx) AuthInfo {
	authorized := c.Locals("authorized").(bool)
	user := c.Locals("user").(models.User)

	return AuthInfo{Authorized: authorized, User: user}
}
