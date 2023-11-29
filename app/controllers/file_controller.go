package controllers

import (
	"fmt"
	"net/http"
	"startup/app"
	"startup/app/helpers"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type fileController struct{}

func NewFileController() *fileController {
	return &fileController{}
}

func (h *fileController) Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var request structs.FileUploadRequest
	c.ShouldBind(&request)

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterImageTypeValidation(validate)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to upload files.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	fmt.Println("batas loh ini =============")
	// fmt.Println(errBind.Error())
	fmt.Println(form.File["files[]"])
	fmt.Println(form.Value["directory"][0])
}