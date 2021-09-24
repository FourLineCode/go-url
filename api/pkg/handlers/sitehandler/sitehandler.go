package sitehandler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SiteHandler struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewHandler(db *gorm.DB) *fiber.App {
	h := SiteHandler{App: fiber.New(), DB: db}

	h.registerRoutes()

	return h.App
}

func (h *SiteHandler) registerRoutes() {
	h.App.Get("/all", h.getSites)
	h.App.Post("/create", h.createSite)
	h.App.Get("/usersites", h.getSiteByUser)
	h.App.Get("/url/:key", h.getSiteUrl)
}
