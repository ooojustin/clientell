package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	// connect to database and panic if failed
	var err error
	dsn := "alpha:e%2V672vKCrz6z@tcp(rc.justin.ooo:3306)/alpha?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

}
