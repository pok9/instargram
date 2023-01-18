package models

//date +%s
type Like struct {
	Model

	UserID     uint `json:"userID"`
	PostMainID uint `json:"postMainID"`
}
