package main

import (
	"log"
	"net/http"

	"api-golearn/config"
	"api-golearn/v1/auth"
	"api-golearn/v1/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env not loaded")
	}
	db := config.ConnectDB()

	authRepositories := auth.NewRepoAuth(db)
	authService := auth.NewAuthService(*authRepositories)
	authHandler := auth.NewAuthHandler(authService)

	corsHeader := middleware.CORSMiddleware()
	tokenValidator := middleware.TokenValidator()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(string("API Golearn is running on port 8000")))
	})

	router.Use(corsHeader)
	router.POST("v1/login", authHandler.GetLogin)
	v1 := router.Group("/v1", tokenValidator)
	{

		v1.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You've been permitted"})
			return
		})
	}

	router.Run(":8000")
}
