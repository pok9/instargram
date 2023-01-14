package models

import (
	"gorm.io/gorm"
)

type PostMain struct {
	gorm.Model
	Caption string `json:"caption"`
	UserID  uint   `json:"userID"`
	User    User   `gorm:"foreignKey:UserID"`
}
