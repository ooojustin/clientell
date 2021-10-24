package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	// connect to database and panic if failed
	// using remote mysql db in release and local mysqlite db in debug
	var err error
	if gin.Mode() == gin.ReleaseMode {
		dsn := "alpha:e%2V672vKCrz6z@tcp(clientellapp.com:3306)/alpha?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database.")
	}

}
