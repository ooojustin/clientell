package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name" gorm:"size:128"`
	LastName     string `json:"last_name" gorm:"size:128"`
	Email        string `json:"email" gorm:"size:128,unique"`
	PasswordHash []byte `json:"-"`
}

type UserSignupForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
