package models

import "gorm.io/datatypes"

type Person struct {
	ID        uint           `gorm:"primarykey"`
	FirstName string         `json:"firstName" gorm:"size:128"`
	LastName  string         `json:"lastName" gorm:"size:128"`
	Address   datatypes.JSON `json:"address"`
}
