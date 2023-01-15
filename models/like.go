package models

//date +%s
type Like struct {
	Model
	UserID     uint     `json:"userID"`
	User       User     `gorm:"foreignKey:UserID"`
	PostMainID uint     `json:"postMainID"`
	PostMain   PostMain `gorm:"foreignKey:PostMainID"`
}
