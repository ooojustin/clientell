package models

import (
	"fmt"
	"math"

	"clientellapp.com/utils"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	PersonID    uint    `json:"-"`
	Person      *Person `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
	OwnerID     uint    `json:"ownerID"`
	Owner       *User   `json:"owner,omitempty" gorm:"foreignKey:OwnerID;references:ID"`
	Stars       int     `json:"stars"`
	Comment     string  `json:"comment" gorm:"size:256"`
	Tags        string  `json:"tags" gorm:"size:256"` // string containing tags separated by comma
	JobType     string  `json:"jobType" gorm:"size:32"`
	Sentiment   string  `json:"sentiment" gorm:"size:32"`
	NeedsReview bool    `json:"-"`
}

// Analyze sentiment of a rating's comment and stores it automatically.
func (r Rating) UpdateSentiment() {
	if sentiment, err := utils.AnalyzeSentiment(r.Comment); err == nil {
		data := map[string]interface{}{
			"sentiment":    sentiment.Sentiment,
			"needs_review": sentiment.ConfidenceScores.Negative > 0.8,
		}
		DB.Table("ratings").Where("id = ?", r.ID).Updates(data)
	}
}

func GetRating(personID string, ownerID string, rating *Rating) error {
	err := DB.Table("ratings").Where("person_id = ? AND owner_id = ?", personID, ownerID).First(rating).Error
	return err
}

func UpdateAverageStars(id string, db *gorm.DB) {

	// get all ratings for person this rating is created for
	var ratings []Rating
	db.Table("ratings").Where("person_id = ?", id).Find(&ratings)

	// find average of users ratings
	// (this can be done with a sql query too, but for loop for now)
	var sum int
	for _, rating := range ratings {
		sum += rating.Stars
	}
	avg := float64(sum) / float64(len(ratings))
	avg = math.Round(avg*100) / 100

	// patch NaN occurrance from 0 remaining ratings
	if math.IsNaN(avg) {
		avg = 0
	}

	// update average stars variable for person
	person, _ := PersonFromID(id)
	person.AverageStars = avg
	db.Save(person)

}

func (r *Rating) AfterCreate(tx *gorm.DB) (err error) {
	personID := fmt.Sprint(r.Person.ID)
	UpdateAverageStars(personID, tx)
	return nil
}

func (r *Rating) AfterUpdate(tx *gorm.DB) (err error) {
	personID := fmt.Sprint(r.PersonID)
	UpdateAverageStars(personID, tx)
	return nil
}
