package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewHandler(db *gorm.DB) *fiber.App {
	h := UserHandler{App: fiber.New(), DB: db}

	h.registerRoutes()

	return h.App
}

func (h *UserHandler) registerRoutes() {
	h.App.Get("/all", h.getUsers)
	h.App.Get("/info/:id", h.getUserById)
	h.App.Post("/register", h.registerUser)
	h.App.Post("/login", h.loginUser)
	h.App.Get("/authorize", h.authorize)
	h.App.Post("/logout", h.logoutUser)
}
