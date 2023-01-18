package controllers

import (
	"gorm.io/gorm"
)

type PostMain struct {
	DB *gorm.DB
}

type CreatePostMain struct {
	Caption string `json:"caption"`
}
