package models

//date +%s
type Post struct {
	Model
	Path       string   `json:"path"`
	MediaType  string   `json:"mediaType"`
	UserID     string   `json:"userID"`
	User       User     `gorm:"foreignKey:UserID"`
	PostMainID string   `json:"postMainID"`
	PostMain   PostMain `gorm:"foreignKey:PostMainID"`
}
