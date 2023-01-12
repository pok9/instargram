package models

import (
	"gorm.io/gorm"
)

type PostMain struct {
	gorm.Model
	Caption string `json:"caption"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}
