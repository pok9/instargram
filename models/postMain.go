package models

type PostMain struct {
	Model
	Caption string `json:"caption"`
	UserID  uint   `json:"userID"`
	User    User   `gorm:"foreignKey:UserID"`
}
