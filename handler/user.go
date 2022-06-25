package handler

import (
	"financial-tracker-be/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	user user.Service
}

func UserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var userBody user.UserRegisterBody

	err := c.ShouldBindJSON(&userBody)
	newUser, err := h.user.UserRegister(userBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Proceed. Check your email for confirmation ",
		"data":    "",
		"token":   newUser.ConfirmationToken,
	})
}

func (h *userHandler) LoginUser(c *gin.Context) {

	var userAuth user.UserLoginRequest

	c.ShouldBindJSON(&userAuth)

	userToken, err := h.user.UserLoginAes(userAuth)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged In",
		"data":    userToken,
	})
}
