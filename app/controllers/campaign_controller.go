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
}

func NewCampaignController(campaignService services.CampaignService) *campaignController {
	return &campaignController{campaignService}
}

func (h *campaignController) Index(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		res := helpers.ResponseAPI("Something went wrong", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := structs.CampaignResponses(campaigns)

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

	formatter := structs.CampaignResponse(campaign)

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
	newCampaign, err := h.campaignService.CreateCampaign(request, user)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	getCampaign, _ := h.campaignService.GetCampaignBySlug(newCampaign.Slug)

	formatter := structs.CampaignResponse(getCampaign)
	res := helpers.ResponseAPI("Campaign successfully created.", http.StatusOK, "sucess", formatter)
	c.JSON(http.StatusOK, res)
}