package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Setup(router *gin.Engine) {
	userGroup := router.Group("user")
	{
		userGroup.POST("/create", u.CreateAccount)
	}
}

func (u UserController) CreateAccount(c *gin.Context) {
	// TODO
	c.String(http.StatusOK, "create account")
}
