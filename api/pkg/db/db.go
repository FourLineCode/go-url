package db

import (
	"log"

	"github.com/FourLineCode/url-shortener/api/internal/config"
	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Initialize(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DBUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Site{})
}
