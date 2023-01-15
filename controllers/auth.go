package controllers

import (
	"instargram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

type SignUpReq struct {
	//email omitempty
	Email       string `json:"email,omitempty" gorm:"unique" binding:"omitempty,email"` //omitempty คือ ถ้าเป็นค่าว่างก็ไม่ต้อง validate
	PhoneNumber string `json:"phoneNumber,omitempty" gorm:"unique" binding:"omitempty,numeric,len=10" `
	//Username ต้องมีอย่าางน้อย 1 ตัวอักษร และไม่มีช่องว่าง ถ้ามีตัวเลขต้องมีตัวอักษร
	//alphanum คือ ตัวเลข และตัวอักษร
	Username string `json:"username" binding:"required,min=1,alphanum" gorm:"unique"`
	// Username    string `json:"username" binding:"required" gorm:"unique"`
	//Password อย่างน้อย 6 ตัสอักษร และมีตัวเลขอย่างน้อย 1 ตัว
	Password string `json:"password" binding:"required,min=6,number"`
	FullName string `json:"fullName" binding:"required"`
}

type SignUpRes struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Username    string `json:"username"`
	FullName    string `json:"fullName"`
}

func (a *Auth) SignUp(ctx *gin.Context) {
	var signUpReq SignUpReq
	if err := ctx.ShouldBind(&signUpReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	copier.Copy(&user, &signUpReq)

	user.Password = user.GenerateEncryptPassword()

	if err := a.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	serializedUser := SignUpRes{}
	copier.Copy(&serializedUser, &user)

	ctx.JSON(http.StatusCreated, gin.H{"user": serializedUser})
}
