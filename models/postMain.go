package models

type PostMain struct {
	Model
	Caption string `json:"caption"`
	//fk
	UserID string `json:"userID"`

	//seft pk to fk other table
	Post    []Post    `gorm:"foreignkey:PostMainID"`
	Like    []Like    `gorm:"foreignkey:PostMainID"`
	Comment []Comment `gorm:"foreignkey:PostMainID"`
}
