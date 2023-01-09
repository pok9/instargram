package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	Birthdate time.Time `json:"birthdate"`
	Photo     string    `json:"photo"`
}
