package main

import (
	"instargram/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// http://127/0.0.1:8080/uploads/articles/8/photo123.jpg
	r.Static("/uploads", "./uploads")

	routes.Serve(r)

	r.Run(":" + os.Getenv("PORT"))
}
