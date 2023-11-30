package controllers

import (
	"net/http"
	"startup/app"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService services.UserService
	authService services.AuthService
	fileService services.FileService
}

func NewUserController(userService services.UserService, authService services.AuthService, fileService services.FileService) *userController {
	return &userController{userService, authService, fileService}
}

func (h *userController) Register(c *gin.Context) {
	var request structs.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterEmailAvailableValidation(validate, h.userService)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Register account failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := h.userService.Register(request)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.UserResponse(user, &token)
	res := helpers.ResponseAPI("Account successfully registered.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userController) Login(c *gin.Context) {
	var request structs.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := h.userService.Login(request)
	if err != nil {
		var errors []string
		errorMessage := gin.H{"errors": append(errors, "Email or password invalid.")}

		res := helpers.ResponseAPI("Login failed.", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.UserResponse(user, &token)
	res := helpers.ResponseAPI("Successfully logged in.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
	var request structs.CheckEmailRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Email checking failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, _ := h.userService.IsEmailAvailable(request)
	data := gin.H{"is_available": isEmailAvailable}
	message := "Email address has been registered."

	if isEmailAvailable {
		message = "Email is available."
	}

	res := helpers.ResponseAPI(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userController) UploadAvatar(c *gin.Context) {
	var request structs.UploadAvatarRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterExistsInFilesValidation(validate, h.fileService)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to change avatar.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	request.User = c.MustGet("currentUser").(structs.User)

	_, err = h.userService.SaveAvatar(request)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Avatar successfully updated.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, res)
}

func (h *userController) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(structs.User)

	formatter := structs.UserResponse(currentUser, nil)

	res := helpers.ResponseAPI("Successfully fetch user data.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}
