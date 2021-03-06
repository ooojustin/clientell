package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"clientellapp.com/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RatingController struct{}

func (r RatingController) List(c *gin.Context) {

	u, _ := c.Get("user")
	user := u.(*models.User)
	userId := fmt.Sprint(user.ID)

	// query for ratings created by the authenticated user
	var ratings []models.Rating
	if err := models.DB.Table("ratings").Where("owner_id = ?", userId).Find(&ratings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// loop through ratings to modify them
	for idx, rating := range ratings {

		// query database for the person who this rating was for
		var person models.Person
		if err := models.DB.Table("people").Where("id = ?", fmt.Sprint(rating.PersonID)).First(&person).Error; err != nil {
			// failed to locate person for this rating
			continue
		}

		// copy values from person object to simplified searh result struct
		var psr models.PersonSearchResult
		copier.Copy(&psr, &person)
		ratings[idx].PersonData = &psr

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ratings,
	})

}

func (r RatingController) ReviewRating(c *gin.Context) {

	user, _ := c.Get("user")
	if !user.(*models.User).Staff {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "Staff only."})
		return
	}

	var rating models.Rating
	if err := models.DB.Table("ratings").Where("needs_review = 1").Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	action := c.Param("action")
	if action == "approve" {
		// rating has been approved and no longer needs to be reviewed
		models.DB.Table("ratings").Where("id = ?", rating.ID).Update("needs_review", "0")
		go models.UpdateAverageStars(fmt.Sprint(rating.PersonID), models.DB)
	} else if action == "deny" {
		// mark rating as hidden
		params := map[string]interface{}{
			"needs_review": false,
			"hidden":       true,
		}
		models.DB.Table("ratings").Where("id = ?", rating.ID).Updates(params)
	} else {
		// unhandled action
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid action."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})

}

// Returns 5 ratings which need to be reviewed.
func (r RatingController) ReviewList(c *gin.Context) {

	user, _ := c.Get("user")
	if !user.(*models.User).Staff {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "Staff only."})
		return
	}

	var count int64
	models.DB.Table("ratings").Where("needs_review = 1").Where("deleted_at IS NULL").Count(&count)

	var ratings []models.Rating
	models.DB.Table("ratings").Where("needs_review = 1").Limit(5).Find(&ratings)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ratings,
		"count":   count,
	})

}

func (r RatingController) Update(c *gin.Context) {

	// parse user request body into rating
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

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

	// update existing rating from user json
	// also update sentiment analysis via goroutine if needed
	updateSentiment := rating.Comment != userRating.Comment
	models.DB.Model(&userRating).Updates(rating)
	if updateSentiment {
		go userRating.UpdateSentiment()
	}

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
	go models.UpdateAverageStars(c.Param("id"), models.DB)

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

	// run rating sentiment analysis via goroutine
	go rating.UpdateSentiment()

	// return serialized rating
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rating,
	})

}
