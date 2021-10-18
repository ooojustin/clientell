package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rc.justin.ooo/config"
	"rc.justin.ooo/controllers"
)

func main() {

	// initialize database connection
	config.InitDatabase()

	// configure gin router with cors middleware
	router := gin.Default()
	router.Use(cors.Default())

	// index path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "rc.justin.ooo")
	})

	// setup user routes
	user := new(controllers.UserController)
	user.Setup(router)

	router.Run()

}
