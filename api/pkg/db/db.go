package db

import (
	"log"

	"github.com/FourLineCode/url-shortener/api/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Initialize() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
