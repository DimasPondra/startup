package controllers

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionService services.TransactionService
	campaignService services.CampaignService
}

func NewTransactionController(transactionService services.TransactionService, campaignService services.CampaignService) *transactionController {
	return &transactionController{transactionService, campaignService}
}

func (h *transactionController) Index(c *gin.Context) {
	user := c.MustGet("currentUser").(structs.User)

	transactions, err := h.transactionService.GetTransactionsByUserID(user.ID)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.TransactionsSummaryResponse(transactions)
	res := helpers.ResponseAPI("List of transactions.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *transactionController) Store(c *gin.Context) {
	var request structs.TransactionStoreRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helpers.ResponseAPI("Failed to create a transaction.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.campaignService.GetCampaignByID(request.CampaignID)
	if err != nil {
		var errors []string
		errorMessage := gin.H{"errors": append(errors, "Campaign not found.")}

		res := helpers.ResponseAPI("Failed to create a transaction.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	request.User = user

	transaction, err := h.transactionService.CreateTransaction(request)
	if err != nil {
		res := helpers.ResponseAPI("Server error, something went wrong.", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.TransactionResponse(transaction)
	res := helpers.ResponseAPI("Transaction successfully created.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}