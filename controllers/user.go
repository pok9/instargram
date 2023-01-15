package controllers

import (
	"fmt"
	"instargram/models"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

type userRes struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
}

// /auth/profile => jwt => sub (UserID) => User => User
func (u *User) GetProfile(ctx *gin.Context) {
	fmt.Println("user2 => ", ctx.Keys["sub"])
	// user
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	fmt.Println("user => ", user)

	var serializedUser userRes
	copier.Copy(&serializedUser, user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})
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

type UpdateUserAvatar struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type updateUserAvatarRes struct {
	Avatar string `json:"avatar,omitempty"`
}

func (u *User) UpdateUserAvatar(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	// Get file
	file, err := ctx.FormFile("avatar")
	if err != nil || file == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Avatar != "" {
		// Delete old file
		//http://127.0.0.1:5000/upload/articles/<ID>/image.png
		// 1. /upload/articles/<ID>/image.png
		user.Avatar = strings.Replace(user.Avatar, os.Getenv("HOST"), "", 1)
		// 2.<WD>/upload/articles/<ID>/image.png
		pwd, _ := os.Getwd()
		fmt.Println("os.Getwd() => ", pwd)
		// 3.Remove <WD>/upload/articles/<ID>image.png
		fmt.Println("user.Avatar => ", user.Avatar)
		os.Remove(pwd + user.Avatar)
	}

	// Create Path
	// ID => 8, uploads/users/8
	path := "uploads/users/" + string(user.ID)
	os.MkdirAll(path, 0755)

	// Upload File

	//uploads/users/8/image.png
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Attach File to user
	user.Avatar = os.Getenv("HOST") + "/" + filename

	if err := u.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedUser := updateUserAvatarRes{}
	copier.CopyWithOption(&serializedUser, &user, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})
}

//follow user
type followUserReq struct {
	FollowedUserID string `json:"followedUserID" binding:"required"`
}

type followUserRes struct {
	FollowingUserID string `json:"followingUserID,omitempty"`
	FollowedUserID  string `json:"followedUserID,omitempty"`
}

func (u *User) FollowUser(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	var followUserReq followUserReq
	if err := ctx.ShouldBindJSON(&followUserReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	//user.ID != followUserReq.FollowedUserID
	if user.ID == followUserReq.FollowedUserID {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "you can't follow yourself"})
		return
	}

	follower := models.Follower{
		FollowingUserID: user.ID,
		FollowedUserID:  followUserReq.FollowedUserID,
	}
	if err := u.DB.Create(&follower).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedFollower := followUserRes{}
	copier.CopyWithOption(&serializedFollower, &follower, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	ctx.JSON(http.StatusOK, gin.H{"follower": serializedFollower})
}

//unfollow user
type unfollowUserReq struct {
	FollowedUserID string `json:"followedUserID" binding:"required"`
}

type unfollowUserRes struct {
	FollowingUserID string `json:"followingUserID,omitempty"`
	FollowedUserID  string `json:"followedUserID,omitempty"`
}

func (u *User) UnfollowUser(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	var unfollowUserReq unfollowUserReq
	if err := ctx.ShouldBindJSON(&unfollowUserReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	//user.ID != followUserReq.FollowedUserID
	if user.ID == unfollowUserReq.FollowedUserID {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "you can't unfollow yourself"})
		return
	}

	follower := models.Follower{
		FollowingUserID: user.ID,
		FollowedUserID:  unfollowUserReq.FollowedUserID,
	}
	if err := u.DB.Delete(&follower).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedFollower := unfollowUserRes{}
	copier.CopyWithOption(&serializedFollower, &follower, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	ctx.JSON(http.StatusOK, gin.H{"follower": serializedFollower})
}
