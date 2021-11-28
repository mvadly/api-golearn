package user

import (
	"api-golearn/v1/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUser(router *gin.Engine, db *gorm.DB) {

	userRepositories := NewRepoUser(db)
	userService := NewUserService(*userRepositories)
	userHandler := NewUserHandler(userService)
	tokenValidator := middleware.TokenValidator()

	v1 := router.Group("/v1/", tokenValidator)
	{

		// v1.POST("/users", userHandler.GetUsers)
		v1.POST("/create_user", userHandler.CreateUser)
	}
}
