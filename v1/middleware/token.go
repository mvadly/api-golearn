package middleware

import (
	"api-golearn/v1/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Authorization is invalid"})
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}

		lenBearer := len(BEARER_SCHEMA)
		tokenString := authHeader[lenBearer:]
		token, err := util.ValidateTokenString(tokenString)

		if token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
