package main

import (
	"net/http"

	"clientellapp.com/controllers"
	"clientellapp.com/middlewares"
	"clientellapp.com/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		AllowHeaders:    []string{"Token", "content-type"},
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
		c.String(http.StatusOK, "clientell!")
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
	router.GET("/person/:id", person.Retrieve)

	rating := new(controllers.RatingController)
	router.GET("/person/:id/rating", rating.Retrieve)
	router.GET("/person/:id/ratings", rating.List)
	router.PATCH("/person/:id/editRating", rating.Update)
	router.POST("/person/:id/createRating", rating.Create)
	router.POST("/person/:id/deleteRating", rating.Delete)
	router.GET("/listReviewRatings", rating.ReviewList)
	router.POST("/reviewRating/:id/:action", rating.ReviewRating)

}

func main() {

	// enable release mode
	// gin.SetMode(gin.ReleaseMode)

	// setup database & router
	SetupDatabase()
	SetupRouter()

	// run router
	router.Run()

}
