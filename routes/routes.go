package routes

import (
	"instargram/config"
	"instargram/controllers"
	"instargram/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")
	authenticate := middleware.Authenticate().MiddlewareFunc()

	authGroup := v1.Group("/auth")
	autController := controllers.Auth{DB: db}
	authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
	authGroup.POST("/sign-up", autController.SignUp)
	// authGroup.Use(authenticate)
	// {
	// authGroup.PATCH("/profile", authenticate, autController.UpdateProfile)
	// }

	userController := controllers.User{DB: db}
	userGroup := v1.Group("/user")
	userGroup.Use(authenticate)
	{
		userGroup.GET("/profile", userController.GetProfile)
		userGroup.PUT("/birdate", userController.UpdateUserBirdate)
		userGroup.PUT("/avatar", userController.UpdateUserAvatar)
		userGroup.POST("/follow", userController.FollowUser)
	}
}
