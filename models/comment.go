package models

//date +%s
//illegal cycle in declaration of Comment คือ มันเรียกใช้ Comment ซ้ำกัน ทำให้เกิดการเรียกใช้งานซ้ำกัน
type Comment struct {
	Model
	Comment            string   `json:"comment"`
	UserID             string   `json:"userID"`
	User               User     `gorm:"foreignKey:UserID"`
	PostMainID         string   `json:"postMainID"`
	PostMain           PostMain `gorm:"foreignKey:PostMainID"`
	CommentRepliedToID string   `json:"commentRepliedToID"`
	Replied            *Comment `gorm:"foreignKey:CommentRepliedToID"`
}
