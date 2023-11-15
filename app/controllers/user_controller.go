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

		res := helpers.ResponseAPI("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := h.userService.Register(request)
	if err != nil {
		res := helpers.ResponseAPI("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := structs.UserResponse(user, "tokentokentoken")
	res := helpers.ResponseAPI("Account successfully registered.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userController) Login(c *gin.Context) {
	var request structs.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := h.userService.Login(request)
	if err != nil {
		errorMessage := gin.H{"errors": "email or password invalid."}

		res := helpers.ResponseAPI("Login failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	formatter := structs.UserResponse(user, "tokentokentoken")
	res := helpers.ResponseAPI("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}