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

type LoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

func (h *UserHandler) loginUser(c *fiber.Ctx) error {
	input := new(LoginUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid credentials"))
	}
	input.Email = strings.ToLower(input.Email)

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("Invalid credentials"))
	}

	var user models.User
	result := h.DB.Where("email = ?", input.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError("User not found with given email"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewError("Invalid credentials"))
	}

	claims := utils.LoginClaims{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	token, err := utils.NewJWTSignedToken(claims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewError("Error signing authorized token"))
	}

	c.Cookie(&fiber.Cookie{
		Name:  "auth",
		Value: token,
	})

	return c.Status(fiber.StatusOK).JSON(user)
}
