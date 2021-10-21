package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rc.justin.ooo/controllers"
	"rc.justin.ooo/middlewares"
	"rc.justin.ooo/models"
)

var router *gin.Engine

func SetupDatabase() {
	// initialize database connection and migrate models
	models.InitDatabase()
	models.DB.AutoMigrate(&models.User{})
	models.DB.AutoMigrate(&models.Person{})
	models.DB.AutoMigrate(&models.Rating{})
}

func SetupRouter() {

	// create router and cors configuration
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Token"},
		AllowMethods:    []string{"GET", "POST", "PATCH", "PUT", "HEAD"},
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

	router.GET("/places", controllers.SearchPlaces)

	user := new(controllers.UserController)
	router.GET("/user/:id", user.Retrieve)
	router.GET("/user", user.TokenRetrieve)
	router.PATCH("/user", user.Update)
	router.POST("/logout", user.Logout)

	person := new(controllers.PersonController)
	router.POST("/person/create", person.Create)
	router.POST("/person/search", person.Search)

}

func main() {

	// setup database & router
	SetupDatabase()
	SetupRouter()

	// run router
	router.Run()

}
