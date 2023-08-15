package config

import (
	"github.com/cyn166/urubu-do-pix/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Deposit{})
}
