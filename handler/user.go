package handler

import (
	"backend-golang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// TANGKAP INPUT USER
	// MAP INPUT DARI USER KE STUCT REGISTER USER INPUT
	// STRUCT DIATAS KITA PASSING SEBAGAI PARAMETER SERVICE
	
	var input user.RegisterUserInput
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)

}