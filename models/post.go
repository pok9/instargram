package models

//date +%s
type Post struct {
	Model
	Path      string `json:"path"`
	MediaType string `json:"mediaType"`

	//มี fk 2 ตัวคือ user,postMain
	UserID     uint `json:"userID"`
	PostMainID uint `json:"postMainID"`
}
