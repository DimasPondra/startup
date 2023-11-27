package controllers

import (
	"net/http"
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
)

type webhookController struct {
	webhookService services.WebhookService
}

func NewWebhookController(webhookService services.WebhookService) *webhookController {
	return &webhookController{webhookService}
}

func (h *webhookController) MidtransNotification(c *gin.Context) {
	var request structs.PaymentNotificationRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.webhookService.MidtransNotification(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
