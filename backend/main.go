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

var router *gin.Engine

func SetupDatabase() {
	// initialize database connection and migrate models
	db.InitDatabase()
	db.DB.AutoMigrate(&models.User{})
}

func SetupRouter() {

	// create router and cors configuration
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Token"},
	}))

	// add routes & enable authentication middleware
	UnauthorizedRoutes()
	router.Use(middlewares.AuthMiddleware())
	AuthorizedRoutes()

}

func UnauthorizedRoutes() {
	// initialize requests that don't require a token

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "rc.justin.ooo")
	})

	user := new(controllers.UserController)
	router.POST("/login", user.Login)
	router.POST("/create_account", user.Create)

}

func AuthorizedRoutes() {
	// initialize requests that require a token

	user := new(controllers.UserController)
	router.GET("/user/:id", user.Retrieve)
	router.GET("/user", user.TokenRetrieve)

}

func main() {

	// setup database & router
	SetupDatabase()
	SetupRouter()

	// run router
	router.Run()

}
