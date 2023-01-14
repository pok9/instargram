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
	authGroup.POST("/sign-up-with-email", autController.SignUpWithEmail)
	authGroup.PATCH("/:email", autController.Update)
	authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
	authGroup.POST("/sign-up", autController.SignUp)
	authGroup.Use(authenticate)
	{
		authGroup.GET("/profile", authenticate, autController.GetProfile)
		authGroup.PATCH("/profile", authenticate, autController.UpdateProfile)
	}

	userGroup := v1.Group("/user")
	userController := controllers.User{DB: db}
	userGroup.Use(authenticate)
	{
		userGroup.PUT("/birdate", userController.UpdateUserBirdate)
	}
}
