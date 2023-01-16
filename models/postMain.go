package models

type PostMain struct {
	Model
	Caption string `json:"caption"`
	UserID  string `json:"userID"`
	User    User   `gorm:"foreignKey:UserID"`
}
