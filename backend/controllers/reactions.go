package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"clientellapp.com/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReactionController struct{}

func (r ReactionController) Delete(c *gin.Context) {

	ratingId := c.Param("id")
	user, _ := c.Get("user")
	userId := fmt.Sprint(user.(*models.User).ID)

	// retrieve the user's reaction to this rating from the database
	var userReaction models.Reaction
	err := models.DB.Table("reactions").Where("rating_id = ? AND owner_id = ?", ratingId, userId).First(&userReaction).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// get the rating from the database
	var rating models.Rating
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Failed to retrieve rating."})
		return
	}

	// determine new reaction count for emoji
	val := 0
	switch userReaction.Type {
	case "thumbs_down":
		val = rating.RThumbsDown - 1
	case "thumbs_up":
		val = rating.RThumbsUp - 1
	case "funny":
		val = rating.RFunny - 1
	case "fire":
		val = rating.RFire - 1
	case "heart":
		val = rating.RHeart - 1
	}

	// update rating's reaction count for the selected type
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).Update("r_"+userReaction.Type, val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// refresh rating variable from database
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Failed to retrieve rating."})
		return
	}

	// delete reaction from database
	if err := models.DB.Delete(&userReaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rating,
	})

}

func (r ReactionController) Create(c *gin.Context) {

	ratingId := c.Param("id")
	user, _ := c.Get("user")
	userId := fmt.Sprint(user.(*models.User).ID)

	// get the rating from the database
	var rating models.Rating
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Failed to retrieve rating."})
		return
	}

	// check if user has already reaction to this rating
	var userReaction models.Reaction
	err := models.DB.Table("reactions").Where("rating_id = ? AND owner_id = ?", ratingId, userId).First(&userReaction).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "User has already reacted to this rating.",
		})
		return
	}

	// initialize reaction and set values
	var reaction models.Reaction
	c.BindJSON(&reaction)
	reaction.OwnerID = user.(*models.User).ID
	reaction.RatingID = rating.ID

	// determine new reaction count for selected type
	val := 0
	switch reaction.Type {
	case "thumbs_down":
		val = rating.RThumbsDown + 1
	case "thumbs_up":
		val = rating.RThumbsUp + 1
	case "funny":
		val = rating.RFunny + 1
	case "fire":
		val = rating.RFire + 1
	case "heart":
		val = rating.RHeart + 1
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid reaction type.",
		})
		return
	}

	// update rating's reaction count for the selected type
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).Update("r_"+reaction.Type, val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create reaction: " + err.Error(),
		})
		return
	}

	// create reaction in database
	if err := models.DB.Save(&reaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create reaction: " + err.Error(),
		})
		return
	}

	// refresh rating variable from database
	if err := models.DB.Table("ratings").Where("id = ?", ratingId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Failed to retrieve rating."})
		return
	}
	rating.Reaction = reaction.Type

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"reaction": reaction,
			"rating":   rating,
		},
	})

}
