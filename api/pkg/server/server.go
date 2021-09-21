package server

import (
	"github.com/FourLineCode/url-shortener/api/pkg/db"
	"github.com/FourLineCode/url-shortener/api/pkg/handlers/userhandler"
	"github.com/FourLineCode/url-shortener/api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func New() *fiber.App {
	dbConn := db.Initialize()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))
	app.Use(middleware.AuthMiddleware)

	registerRoutes(app, dbConn)

	return app
}

func registerRoutes(app *fiber.App, db *gorm.DB) {
	app.Mount("/user", userhandler.NewHandler(db))
}
