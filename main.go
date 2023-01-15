package main

import (
	"fmt"
	"go/build"
	"instargram/config"
	"instargram/migrations"
	"instargram/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("build.Default.GOPATH => ", build.Default.GOPATH) //C:\Users\Administrator\go
	// go env -w GO111MODULE=off
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	defer config.CloseDB()
	migrations.Migrate()

	r := gin.Default()

	// http://127/0.0.1:8080/uploads/articles/8/photo123.jpg
	r.Static("/uploads", "./uploads")
	// uploadDirs := [...]string{"articles", "users"}
	uploadDirs := [...]string{"users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)

	r.Run(":" + os.Getenv("PORT"))
}
