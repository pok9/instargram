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

type userRes struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
}

// /auth/profile => jwt => sub (UserID) => User => User
func (a *Auth) GetProfile(ctx *gin.Context) {
	// user
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	var serializedUser userRes
	copier.Copy(&serializedUser, user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})
}

// func setUserImage(ctx *gin.Context, user *models.User) error {
// 	file, _ := ctx.FormFile("avatar")
// 	if file == nil {
// 		return nil
// 	}

// 	if user.Avatar != "" { //ลบรูปภาพเก่าทั้ง
// 		user.Avatar = strings.Replace(user.Avatar, os.Getenv("HOST"), "", 1)
// 		pwd, _ := os.Getwd()
// 		os.Remove(pwd + user.Avatar)
// 	}

// 	path := "uploads/users/" + strconv.Itoa(int(user.ID))
// 	os.MkdirAll(path, os.ModePerm)
// 	filename := path + "/" + file.Filename
// 	if err := ctx.SaveUploadedFile(file, filename); err != nil {
// 		return err
// 	}

// 	db := config.GetDB()
// 	user.Avatar = os.Getenv("HOST") + "/" + filename
// 	db.Save(user)

// 	return nil
// }
