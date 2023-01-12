package controllers

import (
	"instargram/models"
	"instargram/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

type signUpWithEmailReq struct {
	Email string `json:"email" binding:"required"`
}
type signUpWithEmailRes struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (a *Auth) SignUpWithEmail(ctx *gin.Context) {
	var signUpWithEmailReq signUpWithEmailReq
	if err := ctx.ShouldBindJSON(&signUpWithEmailReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	//copier จะเอาชื่อ struct มาใส่ เช่น struct:Email = models.User struct:Email
	copier.Copy(&user, &signUpWithEmailReq)

	if err := a.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	serializedUser := signUpWithEmailRes{}
	copier.Copy(&serializedUser, &user)

	ctx.JSON(http.StatusCreated, gin.H{"user": serializedUser})
}

type updateUserFormReq struct {
	Password  string    `form:"password,omitempty"`
	FullName  string    `form:"fullName,omitempty"`
	Birthdate time.Time `form:"birthdate,omitempty"`
	Avatar    string    `form:"avatar,omitempty"`
}
type updateUserFormRes struct {
	Email     string    `json:"email,omitempty" `
	FullName  string    `json:"fullName,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
}

func (a *Auth) Update(ctx *gin.Context) {
	email := ctx.Param("email")

	var form updateUserFormReq
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := a.DB.Where("email = ?", email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//copy with out empty string
	copier.CopyWithOption(&user, &form, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if form.Password != "" {
		user.Password = user.GenerateEncryptPassword()
	}

	utils.PrintStruct(form)
	utils.PrintStruct(user)
	if err := a.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedUser updateUserFormRes
	copier.Copy(&serializedUser, &user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})
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

type updateProfileFormReq struct {
	FullName  string    `form:"fullName"`
	Birthdate time.Time `form:"birthdate"`
	Avatar    string    `form:"avatar"`
}

func (a *Auth) UpdateProfile(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user := sub.(*models.User)

	var form updateProfileFormReq
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	copier.CopyWithOption(&user, &form, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if err := a.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

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
