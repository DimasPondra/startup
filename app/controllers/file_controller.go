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

type fileController struct {
	fileService services.FileService
}

func NewFileController(fileService services.FileService) *fileController {
	return &fileController{fileService}
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

	files := form.File["files[]"]
	directory := form.Value["directory"][0]

	var newFiles []structs.File

	for _, file := range files {
		filename := helpers.GenerateRandomFileName(file.Filename)
		path := "images/" + directory + "/" + filename

		err := c.SaveUploadedFile(file, path)
		if err != nil {
			res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		storeRequest := structs.FileStoreRequest{
			Name:     filename,
			Location: directory,
		}

		newFile, err := h.fileService.SaveFile(storeRequest)
		if err != nil {
			res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		newFiles = append(newFiles, newFile)
	}

	formatter := structs.FilesSummaryResponse(newFiles)
	res := helpers.ResponseAPI("Files successfully uploaded.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}
