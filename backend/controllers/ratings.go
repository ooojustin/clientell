package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rc.justin.ooo/models"
)

type RatingController struct{}

func (r RatingController) Update(c *gin.Context) {

	var rating models.Rating
	c.BindJSON(&rating)

	// find the authenticated users rating for this person
	user, _ := c.Get("user")
	var userRating models.Rating
	err := models.GetRating(c.Param("id"), fmt.Sprint(user.(*models.User).ID), &userRating)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "You have no rating for this user.",
		})
		return
	}

	// copy rating variables to db record instance and save
	userRating.Comment = rating.Comment
	userRating.Stars = rating.Stars
	models.DB.Save(&userRating)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userRating,
	})

}

func (r RatingController) Retrieve(c *gin.Context) {

	// get authenticated users rating for specified person
	user, _ := c.Get("user")
	var userRating models.Rating
	err := models.GetRating(c.Param("id"), fmt.Sprint(user.(*models.User).ID), &userRating)

	// handle error if record does not exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "You have no rating for this user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userRating,
	})

}

func (r RatingController) Delete(c *gin.Context) {

	// find the authenticated users rating for this person
	user, _ := c.Get("user")
	var userRating models.Rating
	err := models.GetRating(c.Param("id"), fmt.Sprint(user.(*models.User).ID), &userRating)

	// handle exception if record isnt found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "You have no rating for this user.",
		})
		return
	}

	// delete record from database
	if err = models.DB.Delete(&userRating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// trigger update avg star count for person
	models.UpdateAverageStars(c.Param("id"), models.DB)

	c.JSON(http.StatusOK, gin.H{"success": true})

}

func (r RatingController) Create(c *gin.Context) {

	var err error
	id := c.Param("id")
	user, _ := c.Get("user")
	userId := fmt.Sprint(user.(*models.User).ID)

	// check number of times this user has left a rating in the past 24 hours
	ytd := time.Now().Add(time.Hour * -24) // time 1 day ago
	var ratings []models.Rating
	if err = models.DB.Table("ratings").Where("created_at > ? AND owner_id = ?", ytd, userId).Find(&ratings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// prevent user from creating over 5 ratings a day
	if len(ratings) >= 5 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "You have created too many ratings recently.",
		})
		return
	}

	// retrieve person to rate from database
	var person *models.Person
	person, err = models.PersonFromID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// check if the authenticated user has already rated this person
	// if that's the case, prevent them from creating another rating
	var userRating models.Rating
	err = models.GetRating(id, userId, &userRating)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "You have already rated this user.",
		})
		return
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
		return
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
