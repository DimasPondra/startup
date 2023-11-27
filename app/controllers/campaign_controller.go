package controllers

import (
	"net/http"
	"startup/app"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type campaignController struct {
	campaignService services.CampaignService
	campaignImageService services.CampaignImageService
	transactionService services.TransactionService
}

func NewCampaignController(campaignService services.CampaignService, campaignImageService services.CampaignImageService, transactionService services.TransactionService) *campaignController {
	return &campaignController{campaignService, campaignImageService, transactionService}
}

func (h *campaignController) Index(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.CampaignsSummaryResponse(campaigns)

	res := helpers.ResponseAPI("List of campaigns.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Show(c *gin.Context) {
	slug := c.Param("slug")

	campaign, err := h.campaignService.GetCampaignBySlug(slug)
	if err != nil {
		res := helpers.ResponseAPI("Campaign not found.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	formatter := structs.CampaignResponse(campaign)
	res := helpers.ResponseAPI("Detail campaign.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Store(c *gin.Context) {
	var request structs.CampaignStoreRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterCampaignNameAvailableValidation(validate, h.campaignService, c)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to create a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	request.User = user

	_, err = h.campaignService.CreateCampaign(request)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Campaign successfully created.", http.StatusOK, "sucess", nil)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Update(c *gin.Context) {
	var request structs.CampaignUpdateRequest

	slug := c.Param("slug")
	campaign, err := h.campaignService.GetCampaignBySlug(slug)
	if err != nil {
		res := helpers.ResponseAPI("Campaign not found.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	if campaign.UserID != user.ID {
		res := helpers.ResponseAPI("Can't access this data.", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, res)
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ResponseAPI("Something wrong with the request.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = app.RegisterCampaignNameAvailableValidation(validate, h.campaignService, c)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		errors := helpers.FormatMessageValidationErrors(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to update a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.campaignService.UpdateCampaign(request, campaign)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Campaign successfully updated.", http.StatusOK, "sucess", nil)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) UploadImages(c *gin.Context) {
	campaign, err := h.campaignService.GetCampaignBySlug(c.Param("slug"))
	if err != nil {
		res := helpers.ResponseAPI("Campaign Not Found.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	if user.ID != campaign.UserID {
		res := helpers.ResponseAPI("Can't access this data.", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, res)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		var errors []string
		errorMessage := gin.H{"errors": append(errors, "Field Files is required.")}
		
		res := helpers.ResponseAPI("Failed to upload images.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	files := form.File["files[]"]

	if len(files) == 0 {
		var errors []string
		errorMessage := gin.H{"errors": append(errors, "Field Files is required.")}
		
		res := helpers.ResponseAPI("Failed to upload images.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.campaignImageService.DeleteImages(campaign.ID)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	for index, file := range files {
		primary := index == 0

		filename := helpers.GenerateRandomFileName(file.Filename)
		path := "images/campaigns/" + filename
		
		err := c.SaveUploadedFile(file, path)
		if err != nil {
			res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		_, err = h.campaignImageService.SaveImage(path, primary, campaign)
		if err != nil {
			res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}

	res := helpers.ResponseAPI("Campaign image successfully uploaded.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) ShowTransactions(c *gin.Context) {
	slug := c.Param("slug")

	campaign, err := h.campaignService.GetCampaignBySlug(slug)
	if err != nil {
		res := helpers.ResponseAPI("Campaign not found.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	if campaign.UserID != user.ID {
		res := helpers.ResponseAPI("Can't access this data.", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, res)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByCampaignID(campaign.ID)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.CampaignTransactionsResponse(transactions)
	res := helpers.ResponseAPI("List of transactions by campaign.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}