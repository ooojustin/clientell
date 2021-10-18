package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/config"
	"rc.justin.ooo/models"
)

type UserController struct{}

func (u UserController) Setup(router *gin.Engine) {
	userGroup := router.Group("user")
	{
		userGroup.POST("/create", u.Create)
	}
}

func (u UserController) Create(c *gin.Context) {

	// bind json posted data to form for creating accounts
	ucf := &models.UserCreateForm{}
	err := c.BindJSON(&ucf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create db user from form
	user := new(models.User)
	dbuser := user.FromUCF(ucf)

	// create object in database
	config.DB.Create(dbuser)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}
