package middleware

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("auth-token")

	valid, claims, err := utils.ValidateJWTSignedToken(token)
	if err != nil || !valid {
		c.Locals("authorized", false)
		c.Locals("user", models.User{})
		return c.Next()
	}

	c.Locals("authorized", true)
	c.Locals("user", claims.User)

	return c.Next()
}
