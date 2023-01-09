package routes

import (
	"instargram/config"
	"instargram/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	usersGroup := r.Group("/api/v1/users")
	userController := controllers.Users{DB: db}
	{
		usersGroup.POST("/sign-up-with-email", userController.SignUpWithEmail)
	}

}
