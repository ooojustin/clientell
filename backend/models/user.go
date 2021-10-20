package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `json:"firstName" gorm:"size:128"`
	LastName  string    `json:"lastName" gorm:"size:128"`
	Email     string    `json:"email" gorm:"size:128,unique"`
	Password  string    `json:"-"`
	Token     uuid.UUID `json:"token"`
}

type UserCreateForm struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateForm struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func SaveUser(user *User) error {
	if err := DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateNewUser(user *User) error {
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UserFrom(field string, value string) (*User, error) {
	var user User
	err := DB.Table("users").Where(field+" = ?", value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserFromID(id string) (*User, error) {
	return UserFrom("id", id)
}

func UserFromToken(token string) (*User, error) {
	return UserFrom("token", token)
}

func UserFromEmail(email string) (*User, error) {
	return UserFrom("email", email)
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
		Token:     uuid.NewV4(),
	}

	return user

}
