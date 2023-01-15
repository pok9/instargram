package models

//date +%s
//illegal cycle in declaration of Comment คือ มันเรียกใช้ Comment ซ้ำกัน ทำให้เกิดการเรียกใช้งานซ้ำกัน
type Comment struct {
	Model
	Comment            string   `json:"comment"`
	UserID             uint     `json:"userID"`
	User               User     `gorm:"foreignKey:UserID"`
	PostMainID         uint     `json:"postMainID"`
	PostMain           PostMain `gorm:"foreignKey:PostMainID"`
	CommentRepliedToID uint     `json:"commentRepliedToID"`
	Replied            *Comment `gorm:"foreignKey:CommentRepliedToID"`
}
