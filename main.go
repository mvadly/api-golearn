package main

import (
	"api-golearn/config"
	"api-golearn/v1/auth"
	"api-golearn/v1/middleware"
	"api-golearn/v1/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env not loaded")
	}

	db := config.ConnectDB()

	corsHeader := middleware.CORSMiddleware()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(string("API Golearn is running on port 8000")))
	})

	router.Use(corsHeader)
	auth.RouteAuth(router, db)
	user.RouteUser(router, db)

	router.Run(":8000")
}
