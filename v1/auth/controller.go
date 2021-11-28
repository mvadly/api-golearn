package auth

import (
	"api-golearn/v1/util"
	"net/http"
	"strings"

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

	token, err := util.GenerateToken(util.ResponseTokenCreated(data))

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

func (h *authHandler) GetMyProfile(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.Replace(auth, "Bearer ", "", 1)
	// fmt.Println(token)
	data, err := util.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})

		return
	}

	id := data["id"]
	profile, err := h.authService.Profile(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})

		return
	}

	if len(profile) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "Data not found",
			"req":         id,
		})

		return
	}

	var profileData = ResponseProfile{
		ID:       uint32(profile[0].ID),
		Username: profile[0].Username,
		Email:    profile[0].Email,
		Name:     profile[0].Name,
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"profile":     profileData,
	})

}
