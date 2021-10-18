package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"rc.justin.ooo/db"
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

func CreateNewUser(user *User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UserFromID(id string) (*User, error) {
	var user User
	err := db.DB.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserFromUCF(ucf *UserCreateForm) *User {

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
