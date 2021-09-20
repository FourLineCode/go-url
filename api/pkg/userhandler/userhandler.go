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
	h.App.Get("/", h.getUsers)
	h.App.Get("/:id", h.getUserById)
	h.App.Post("/register", h.registerUser)
	h.App.Post("/login", h.loginUser)
}
