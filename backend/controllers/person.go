package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rc.justin.ooo/models"
)

type PersonController struct{}

func (p PersonController) Retrieve(c *gin.Context) {

	id := c.Param("id")
	person, err := models.PersonFromID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	// retrieve ratings for this person from the database
	var ratings []models.Rating
	err = models.DB.Table("ratings").Where("person_id = ?", id).Find(&ratings).Error

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"person":  person,
			"ratings": ratings,
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
		err = models.DB.Raw("SELECT * FROM people WHERE address LIKE ?", "%"+placeId.(string)+"%").Scan(&people).Error
	} else {
		// search by first and last name
		err = models.DB.Raw("SELECT * FROM people WHERE first_name LIKE ? AND last_name LIKE ?", psf.FirstName+"%", psf.LastName+"%").Scan(&people).Error
	}

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
