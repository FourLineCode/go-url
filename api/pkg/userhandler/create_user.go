package userhandler

import (
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	input := new(CreateUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid user information"))
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Error hashing password"))
	}

	newUser := models.NewUser(input.Email, input.Username, string(hashedPasswordBytes))

	h.DB.Create(&newUser)

	return c.Status(fiber.StatusOK).JSON(newUser)
}
