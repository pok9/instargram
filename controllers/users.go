package controllers

import (
	"instargram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

type signUpWithEmailReq struct {
	Email string `json:"email" binding:"required"`
}
type signUpWithEmailRes struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (u *Users) SignUpWithEmail(ctx *gin.Context) {
	var signUpWithEmailReq signUpWithEmailReq
	if err := ctx.ShouldBindJSON(&signUpWithEmailReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	//copier จะเอาชื่อ struct มาใส่ เช่น struct:Email = models.User struct:Email
	copier.Copy(&user, &signUpWithEmailReq)

	if err := u.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	serializedUser := signUpWithEmailRes{}
	copier.Copy(&serializedUser, &user)

	ctx.JSON(http.StatusCreated, gin.H{"user": serializedUser})
}
