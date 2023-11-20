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