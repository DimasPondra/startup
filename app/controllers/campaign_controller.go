package controllers

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignController struct {
	campaignService services.CampaignService
	campaignImageService services.CampaignImageService
}

func NewCampaignController(campaignService services.CampaignService, campaignImageService services.CampaignImageService) *campaignController {
	return &campaignController{campaignService, campaignImageService}
}

func (h *campaignController) Index(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		res := helpers.ResponseAPI("Something went wrong", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := structs.ResponseCampaigns(campaigns)

	res := helpers.ResponseAPI("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Show(c *gin.Context) {
	slug := c.Param("slug")

	campaign, err := h.campaignService.GetCampaignBySlug(slug)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helpers.ResponseAPI("Campaign not found.", http.StatusNotFound, "error", errorMessage)
		c.JSON(http.StatusNotFound, res)
		return
	}

	formatter := structs.ResponseCampaign(campaign)

	res := helpers.ResponseAPI("Detail campaign.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Store(c *gin.Context) {
	var request structs.CampaignStoreRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to create a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	campaign, _ := h.campaignService.GetCampaignByName(request.Name)
	if campaign.ID != 0 {
		errorMessage := gin.H{"errors": "Name is already in use."}

		res := helpers.ResponseAPI("Failed to create a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	request.User = user
	_, err = h.campaignService.CreateCampaign(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Campaign successfully created.", http.StatusOK, "sucess", nil)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) Update(c *gin.Context) {
	var request structs.CampaignUpdateRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to update a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

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

	if campaign.Name != request.Name {
		checkCampaign, _ := h.campaignService.GetCampaignByName(request.Name)

		if checkCampaign.ID != 0 {
			errorMessage := gin.H{"errors": "Name is already in use."}

			res := helpers.ResponseAPI("Failed to update a campaign.", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	}

	_, err = h.campaignService.UpdateCampaign(request, campaign)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("Campaign successfully updated.", http.StatusOK, "sucess", nil)
	c.JSON(http.StatusOK, res)
}

func (h *campaignController) UploadImages(c *gin.Context) {
	campaign, err := h.campaignService.GetCampaignBySlug(c.Param("slug"))
	if err != nil {
		res := helpers.ResponseAPI("Campaign Not Found", http.StatusNotFound, "error", nil)
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
		errorMessage := gin.H{"errors": "files is required."}
		res := helpers.ResponseAPI("Failed to upload images.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	files := form.File["files[]"]

	if len(files) == 0 {
		errorMessage := gin.H{"errors": "files is required."}
		res := helpers.ResponseAPI("Failed to upload images.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.campaignImageService.DeleteImages(campaign.ID)
	if err != nil {
		res := helpers.ResponseAPI("Failed to delete images.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	for index, file := range files {
		primary := index == 0

		filename := helpers.GenerateRandomFileName(file.Filename)
		path := "images/campaigns/" + filename
		
		err := c.SaveUploadedFile(file, path)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			res := helpers.ResponseAPI("Failed to upload images.", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		_, err = h.campaignImageService.SaveImage(path, primary, campaign)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}

	res := helpers.ResponseAPI("Campaign image successfully uploaded.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, res)
}