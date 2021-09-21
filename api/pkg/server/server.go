package server

import (
	"github.com/FourLineCode/url-shortener/api/pkg/db"
	"github.com/FourLineCode/url-shortener/api/pkg/handlers/userhandler"
	"github.com/FourLineCode/url-shortener/api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New() *fiber.App {
	dbConn := db.Initialize()

	app := fiber.New()

	app.Use(middleware.AuthMiddleware)

	registerRoutes(app, dbConn)

	return app
}

func registerRoutes(app *fiber.App, db *gorm.DB) {
	app.Mount("/user", userhandler.NewHandler(db))
}
