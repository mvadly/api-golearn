package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	var reqPage Pagination
	err := c.BindJSON(&reqPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}

	// pagination := util.GeneratePagination(c)
	// fmt.Println(pagination)
	data, err := h.userService.GetUsers(reqPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": data})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var requestCreateUser RequestCreateUser
	err := c.BindJSON(&requestCreateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}

	create, err := h.userService.CreateUser(requestCreateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     fmt.Sprintf("user (%s) berhasil dibuat", create.Name),
	})

}
