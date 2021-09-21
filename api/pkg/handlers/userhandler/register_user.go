package userhandler

import (
	"errors"
	"strings"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"github.com/FourLineCode/url-shortener/api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=2,max=32"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

func (h *UserHandler) registerUser(c *fiber.Ctx) error {
	input := new(RegisterUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid user information"))
	}
	input.Email = strings.ToLower(input.Email)
	input.Username = strings.ToLower(input.Username)

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid user information"))
	}

	var userExists models.User
	result := h.DB.Where("email = ?", input.Email).Or("username = ?", input.Username).First(&userExists)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("User already exists with given username of email"))
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Error hashing password"))
	}

	newUser := models.NewUser(input.Email, input.Username, string(hashedPasswordBytes))

	h.DB.Create(&newUser)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"user":    newUser,
	})
}
