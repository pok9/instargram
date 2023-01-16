package models

//date +%s
type Like struct {
	Model
	UserID     string   `json:"userID"`
	User       User     `gorm:"foreignKey:UserID"`
	PostMainID string   `json:"postMainID"`
	PostMain   PostMain `gorm:"foreignKey:PostMainID"`
}
