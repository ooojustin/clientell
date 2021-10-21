package models

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Rating struct {
	ID       uint    `gorm:"primarykey"`
	PersonID uint    `json:"-"`
	Person   *Person `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
	Stars    int     `json:"stars"`
	Comment  string  `json:"comment" gorm:"size:1000"`
}

func (r *Rating) AfterCreate(tx *gorm.DB) (err error) {

	// convert person uid to string
	personID := fmt.Sprint(r.Person.ID)

	// get all ratings for person this rating is created for
	var ratings []Rating
	tx.Table("ratings").Where("person_id = ?", personID).Find(&ratings)

	// find average of users ratings
	// (this can be done with a sql query too, but for loop for now)
	var sum int
	for _, rating := range ratings {
		sum += rating.Stars
	}
	avg := float64(sum) / float64(len(ratings))
	avg = math.Round(avg*100) / 100

	// update average stars variable for person
	person, _ := PersonFromID(personID)
	person.AverageStars = avg
	tx.Save(person)

	return nil // no error

}
