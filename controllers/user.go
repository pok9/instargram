package controllers

import (
	"instargram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

//update birthdate  day month year to date
type updateUserBirdateReq struct {
	Day   int `form:"day" binding:"required"`
	Month int `form:"month" binding:"required"`
	Year  int `form:"year" binding:"required"`
}

type updateUserBirdateRes struct {
	Birthdate time.Time `json:"birthdate,omitempty"`
}

func (u *User) UpdateUserBirdate(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	var updateUserBirdateReq updateUserBirdateReq
	if err := ctx.ShouldBind(&updateUserBirdateReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// var user models.User
	user.Birthdate = time.Date(updateUserBirdateReq.Year, time.Month(updateUserBirdateReq.Month), updateUserBirdateReq.Day, 0, 0, 0, 0, time.UTC)
	if err := u.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedUser := updateUserBirdateRes{}
	copier.CopyWithOption(&serializedUser, &user, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})
}
