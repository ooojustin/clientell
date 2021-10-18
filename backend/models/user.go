package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name" gorm:"size:128"`
	LastName     string `json:"last_name" gorm:"size:128"`
	Email        string `json:"email" gorm:"size:128,unique"`
	PasswordHash []byte `json:"-"`
}
