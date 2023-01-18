package models

//date +%s
//illegal cycle in declaration of Comment คือ มันเรียกใช้ Comment ซ้ำกัน ทำให้เกิดการเรียกใช้งานซ้ำกัน
type Comment struct {
	Model
	Comment string `json:"comment"`

	UserID     uint `json:"userID"`
	PostMainID uint `json:"postMainID"`

	CommentID          *uint     `json:"commentID"`
	CommentRepliedToID []Comment `json:"commentRepliedToID" gorm:"foreignkey:CommentID"`
}
