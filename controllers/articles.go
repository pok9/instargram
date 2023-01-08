package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Articles struct{}

func (a *Articles) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello 1",
	})
}

func (a *Articles) FindOne(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello 2",
	})
}
