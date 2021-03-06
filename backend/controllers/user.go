package controllers

import (
	"net/http"

	"clientellapp.com/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (u UserController) Update(c *gin.Context) {
	ret, ok := c.Get("user")
	if ok {
		uuf := &models.UserUpdateForm{}
		c.BindJSON(&uuf)
		user := ret.(*models.User)
		user.FirstName = uuf.FirstName
		user.LastName = uuf.LastName
		if err := models.SaveUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false})
	}
}

func (u UserController) Logout(c *gin.Context) {
	ret, ok := c.Get("user")
	if ok {
		// randomly generate user token and save
		user := ret.(*models.User)
		user.Token = uuid.NewV4()
		if err := models.SaveUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false})
	}
}

func (u UserController) TokenRetrieve(c *gin.Context) {
	// get user from request context
	// it'll be set via the AuthMiddleware handler in middlewares/auth.go
	user, ok := c.Get("user")
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false})
	}
}

func (u UserController) Login(c *gin.Context) {

	ulf := &models.UserLoginForm{}
	c.BindJSON(&ulf)

	// retrieve user record from entered email address
	user, err := models.UserFromEmail(ulf.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	// verify that the user entered the correct password
	bhash := []byte(user.Password)
	bpw := []byte(ulf.Password)
	if bcrypt.CompareHashAndPassword(bhash, bpw) == nil {

		// if this is a staff page login, prevent normal users from proceeding
		if ulf.Staff && !user.Staff {
			c.JSON(http.StatusForbidden, gin.H{"success": false})
			return
		}

		// login successful, return use data
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    user,
		})

	} else {
		// user entered invalid password
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
	if err := c.ShouldBindJSON(&ucf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// create object in database
	user := models.UserFromUCF(ucf)
	err := models.CreateNewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})

}
