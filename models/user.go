package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"` // unique
	Password    string    `json:"password"`
	Username    string    `json:"username" gorm:"uniqueIndex"` // unique
	FullName    string    `json:"fullName"`
	Birthdate   time.Time `json:"birthdate"`
	Avatar      string    `json:"avatar"`

	PostMain        []PostMain `gorm:"foreignkey:UserID"`
	Post            []Post     `gorm:"foreignkey:UserID"`
	Like            []Like     `gorm:"foreignkey:UserID"`
	FollowingUserID []Follower `gorm:"foreignkey:FollowingUserID"`
	FollowedUserID  []Follower `gorm:"foreignkey:FollowedUserID"`
	Comment         []Comment  `gorm:"foreignkey:UserID"`
}

func (u *User) GenerateEncryptPassword() string {
	fmt.Printf("u.Password => %#v", u.Password)
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	return string(hash)
}
