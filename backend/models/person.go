package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName    string         `json:"firstName" gorm:"size:128"`
	LastName     string         `json:"lastName" gorm:"size:128"`
	Address      datatypes.JSON `json:"address"`
	AverageStars float64        `json:"avgStars"`
	CreatorID    uint           `json:"creatorID"`
	Creator      *User          `json:"creator,omitempty" gorm:"foreignKey:CreatorID;references:ID"`
}

type PersonSearchForm struct {
	FirstName string                 `json:"firstName"`
	LastName  string                 `json:"lastName"`
	Address   map[string]interface{} `json:"address"`
}

func PersonFrom(field string, value string) (*Person, error) {
	var person Person
	err := DB.Table("people").Where(field+" = ?", value).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func PersonFromID(id string) (*Person, error) {
	return PersonFrom("id", id)
}
