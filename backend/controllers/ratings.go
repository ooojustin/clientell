package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/models"
)

type RatingController struct{}

func (r RatingController) Create(c *gin.Context) {

	id := c.Param("id")
	person, err := models.PersonFromID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	var rating models.Rating
	c.BindJSON(&rating)

	rating.Person = person

	err = models.DB.Create(&rating).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rating,
	})

}

func (r RatingController) List(c *gin.Context) {

	// get list of ratings from database for specified person
	id := c.Param("id")
	var ratings []models.Rating
	err := models.DB.Table("ratings").Where("person_id = ?", id).Find(&ratings).Error

	// return the error if there was an issue querying
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// return list of ratings
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ratings,
	})

}
