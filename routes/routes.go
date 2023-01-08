package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	articlesGroup := r.Group("/api/v1/articles")

	articlesGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello 1",
		})
	})

	articlesGroup.GET("/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello 2",
		})
	})
}
