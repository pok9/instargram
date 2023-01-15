package models

//date +%s
type Post struct {
	Model
	Path       string   `json:"path"`
	MediaType  string   `json:"mediaType"`
	UserID     uint     `json:"userID"`
	User       User     `gorm:"foreignKey:UserID"`
	PostMainID uint     `json:"postMainID"`
	PostMain   PostMain `gorm:"foreignKey:PostMainID"`
}
