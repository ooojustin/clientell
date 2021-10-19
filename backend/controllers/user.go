package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"rc.justin.ooo/models"
)

type UserController struct{}

func (u UserController) Login(c *gin.Context) {

	ulf := &models.UserLoginForm{}
	c.BindJSON(&ulf)

	user, err := models.UserFromEmail(ulf.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	bhash := []byte(user.Password)
	bpw := []byte(ulf.Password)
	if bcrypt.CompareHashAndPassword(bhash, bpw) == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
	}

}

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
	c.BindJSON(&ucf)

	// create object in database
	user := models.UserFromUCF(ucf)
	models.CreateNewUser(user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})

}