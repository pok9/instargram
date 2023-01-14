package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Password    string    `json:"password"`
	Username    string    `json:"username" gorm:"uniqueIndex"` // unique
	FullName    string    `json:"fullName"`
	Birthdate   time.Time `json:"birthdate"`
	Avatar      string    `json:"avatar"`
}

func (u *User) GenerateEncryptPassword() string {
	fmt.Printf("u.Password => %#v", u.Password)
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	return string(hash)
}
