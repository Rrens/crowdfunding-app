package handler

import (
	"backend-golang/helper"
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
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}


		response := helper.APIResponse("Register Account fail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	if err != nil{
		response := helper.APIResponse("Register Account fail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)

}