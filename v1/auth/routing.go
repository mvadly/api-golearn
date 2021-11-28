package auth

import (
	"api-golearn/v1/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAuth(router *gin.Engine, db *gorm.DB) {

	authRepositories := NewRepoAuth(db)
	authService := NewAuthService(*authRepositories)
	authHandler := NewAuthHandler(authService)
	tokenValidator := middleware.TokenValidator()

	router.POST("v1/login", authHandler.GetLogin)
	v1 := router.Group("/v1", tokenValidator)
	{

		v1.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You've been permitted"})
			return
		})

		v1.GET("/profile", authHandler.GetMyProfile)
	}
}
