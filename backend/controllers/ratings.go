package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rc.justin.ooo/models"
)

type RatingController struct{}

func (r RatingController) Create(c *gin.Context) {

	id := c.Param("id")
	person, err := models.PersonFromID(id)

	// make sure we got person to rate from database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	// check if the authenticated user has already rated this person
	// if that's the case, prevent them from creating another rating
	user, _ := c.Get("user")
	var userRating models.Rating
	err = models.DB.Table("ratings").Where("person_id = ? AND owner_id = ?", id, user.(*models.User).ID).First(&userRating).Error
	noUserRating := errors.Is(err, gorm.ErrRecordNotFound)
	if !noUserRating {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "You have already rated this user.",
		})
	}

	// create rating and establish parameters
	var rating models.Rating
	c.BindJSON(&rating)
	rating.Owner = user.(*models.User)
	rating.Person = person

	// save rating to database and make sure there's no error
	err = models.DB.Create(&rating).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	// return serialized rating
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
