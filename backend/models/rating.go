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
	OwnerID  uint    `json:"ownerID"`
	Owner    *User   `json:"owner,omitempty" gorm:"foreignKey:OwnerID;references:ID"`
	Stars    int     `json:"stars"`
	Comment  string  `json:"comment" gorm:"size:1000"`
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
