package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"size:128"`
	LastName  string `json:"last_name" gorm:"size:128"`
	Email     string `json:"email" gorm:"size:128,unique"`
	Password  string `json:"-"`
}

type UserCreateForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u User) FromUCF(ucf *UserCreateForm) *User {

	// hash password user bcrypt
	bpassword := []byte(ucf.Password)
	bhash, _ := bcrypt.GenerateFromPassword(bpassword, bcrypt.DefaultCost)

	// create gorm user model
	user := &User{
		FirstName: ucf.FirstName,
		LastName:  ucf.LastName,
		Email:     ucf.Email,
		Password:  string(bhash),
	}

	return user

}
