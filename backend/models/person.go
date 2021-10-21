package models

import "gorm.io/datatypes"

type Person struct {
	ID           uint           `gorm:"primarykey"`
	FirstName    string         `json:"firstName" gorm:"size:128"`
	LastName     string         `json:"lastName" gorm:"size:128"`
	Address      datatypes.JSON `json:"address"`
	AverageStars float64        `json:"avgStars"`
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
