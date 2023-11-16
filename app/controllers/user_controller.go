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
	authService services.AuthService
}

func NewUserController(userService services.UserService, authService services.AuthService) *userController {
	return &userController{userService, authService}
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

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.UserResponse(user, token)
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

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.UserResponse(user, token)
	res := helpers.ResponseAPI("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
	var request structs.CheckEmailRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Email checking failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, _ := h.userService.IsEmailAvailable(request)
	data := gin.H{"is_available": isEmailAvailable}
	var message string

	if isEmailAvailable {
		message = "Email is available."
	} else {
		message = "Email address has been registered."
	}

	res := helpers.ResponseAPI(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{"errors": "avatar is required."}
		res := helpers.ResponseAPI("Failed to upload avatar image.", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	filename := helpers.GenerateRandomFileName(file.Filename)
	path := "images/avatars/" + filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helpers.ResponseAPI("Failed to upload avatar image.", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	currentUser := c.MustGet("currentUser").(structs.User)

	_, err = h.userService.SaveAvatar(currentUser.ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helpers.ResponseAPI("Failed to upload avatar image.", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helpers.ResponseAPI("Avatar successfully uploaded.", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userController) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(structs.User)
	
	formatter := structs.UserResponse(currentUser, "")

	res := helpers.ResponseAPI("Successfully fetch user data.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}