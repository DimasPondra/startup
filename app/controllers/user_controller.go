package controllers

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (h *userController) Register(c *gin.Context) {
	var request structs.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.ResponseAPI("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.Register(request)
	if err != nil {
		response := helpers.ResponseAPI("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := structs.UserResponse(user, "tokentokentoken")

	response := helpers.ResponseAPI("Account successfully registered.", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}