package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rc.justin.ooo/controllers"
	"rc.justin.ooo/db"
	"rc.justin.ooo/middlewares"
	"rc.justin.ooo/models"
)

func main() {

	// initialize database connection
	db.InitDatabase()
	db.DB.AutoMigrate(&models.User{})

	// configure gin router with cors middleware
	router := gin.Default()
	router.Use(cors.Default())

	// index path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "rc.justin.ooo")
	})

	// enable authentication middleware
	router.Use(middlewares.AuthMiddleware())

	// setup user routes
	user := new(controllers.UserController)
	user.Setup(router)

	router.Run()

}
