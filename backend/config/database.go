package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rc.justin.ooo/models"
)

var DB *gorm.DB

func InitDatabase() {

	// connect to database and panic if failed
	// using sqlite local db file for now
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	// run model migrations
	DB.AutoMigrate(&models.User{})

}
