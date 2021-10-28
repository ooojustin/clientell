package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"clientellapp.com/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PersonController struct{}

func (p PersonController) Retrieve(c *gin.Context) {

	id := c.Param("id")
	user, _ := c.Get("user")
	person, err := models.PersonFromID(id)
	userId := user.(*models.User).ID

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// retrieve ratings for this person from the database
	var ratings []models.Rating
	qstr := "SELECT * FROM ratings WHERE person_id = ? AND deleted_at IS NULL AND (needs_review = 0 OR owner_id = ?)"
	err = models.DB.Raw(qstr, id, userId).Find(&ratings).Error

	// check if authenticated user has already rated the person
	var userRating models.Rating
	err = models.DB.Table("ratings").Where("person_id = ? AND owner_id = ?", id, userId).First(&userRating).Error
	noUserRating := errors.Is(err, gorm.ErrRecordNotFound)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"person":  person,
			"ratings": ratings,
			"canRate": noUserRating,
		},
	})

}

func (p PersonController) Search(c *gin.Context) {

	var psf models.PersonSearchForm
	c.BindJSON(&psf)

	var err error
	var people []models.Person
	if placeId, ok := psf.Address["place_id"]; ok {
		// search by address - find rows which have the same place id
		err = models.DB.Table("people").Find(&people, datatypes.JSONQuery("address").Equals(placeId.(string), "place_id")).Error
	} else {
		// search by first and last name
		err = models.DB.Raw("SELECT * FROM people WHERE first_name LIKE ? AND last_name LIKE ?", psf.FirstName+"%", psf.LastName+"%").Scan(&people).Error
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    people,
	})

}

func (p PersonController) Create(c *gin.Context) {

	var err error
	user, _ := c.Get("user")
	userId := fmt.Sprint(user.(*models.User).ID)

	// check number of times this user has created a person in the past 24 hours
	ytd := time.Now().Add(time.Hour * -24) // time 1 day ago
	var people []models.Person
	if err = models.DB.Table("people").Where("created_at > ? AND creator_id = ?", ytd, userId).Find(&people).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// prevent user from creating over 5 people a day
	if len(people) >= 5 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "You have created too many people recently.",
		})
		return
	}

	var person models.Person
	c.BindJSON(&person)
	person.Creator = user.(*models.User)

	if err = models.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    person,
	})

}
