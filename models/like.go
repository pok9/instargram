package models

import (
	"gorm.io/gorm"
)

//date +%s
type Like struct {
	gorm.Model
	UserID     uint     `json:"userID"`
	User       User     `gorm:"foreignKey:UserID"`
	PostMainID uint     `json:"postMainID"`
	PostMain   PostMain `gorm:"foreignKey:PostMainID"`
}
