package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/models"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	id := c.Param("id")
	user, err := models.UserFromID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u UserController) Create(c *gin.Context) {

	// bind json posted data to form for creating accounts
	ucf := &models.UserCreateForm{}
	err := c.BindJSON(&ucf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"successs": false,
			"error":    err.Error(),
		})
		return
	}

	// create db user from form
	user := models.UserFromUCF(ucf)

	// create object in database
	models.CreateNewUser(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})

}
