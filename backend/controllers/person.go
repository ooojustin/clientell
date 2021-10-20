package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/models"
)

type PersonController struct{}

func (p PersonController) Search(c *gin.Context) {

	var psf models.PersonSearchForm
	c.BindJSON(&psf)

	var people []models.Person
	err := models.DB.Raw("SELECT * FROM people WHERE (first_name LIKE ? AND last_name LIKE ?)", psf.FirstName, psf.LastName).Scan(&people).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    people,
	})

}

func (p PersonController) Create(c *gin.Context) {

	var person models.Person
	c.BindJSON(&person)

	if err := models.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    person,
	})

}
