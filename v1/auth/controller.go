package auth

import (
	"api-golearn/v1/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) GetLogin(c *gin.Context) {
	var postLogin RequestLogin
	var data = ResponseLogin{}
	err := c.BindJSON(&postLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code":  http.StatusBadRequest,
			"error":        err.Error(),
			"data_request": postLogin,
		})

		return
	}

	login, err := h.authService.Login(postLogin.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})

		return
	}

	if len(login) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "Username / Password not found",
			"data":        login,
		})

		return
	}

	passwordHash := string(login[0].Password)
	passwordPlain := string(postLogin.Password)
	match := util.ComparePasswords(passwordHash, []byte(passwordPlain))

	if match == false {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "Username / Password not match",
		})

		return
	}

	data.ID = uint32(login[0].ID)
	data.Username = login[0].Username
	data.Email = login[0].Email

	token, err := util.GenerateToken(map[string]string{
		"id":       string(rune(data.ID)),
		"username": data.Username,
		"email":    data.Email,
	})

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status_code": http.StatusExpectationFailed,
			"error":       err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Data found!",
		"data":        data,
		"token":       token,
	})
}

func (h *authHandler) ValidateToken(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Authorization is invalid"})
		return
	}

	lenBearer := len(BEARER_SCHEMA)
	tokenString := authHeader[lenBearer:]
	token, err := util.ValidateTokenString(tokenString)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		c.JSON(http.StatusOK, gin.H{"jwt": claims})

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

}
