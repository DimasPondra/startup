package controllers

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
	"startup/app/structs"
	"strconv"

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
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaignID, _ := strconv.Atoi(c.Query("campaign_id"))

	transactions, err := h.transactionService.GetTransactions(userID, campaignID)
	if err != nil {
		res := helpers.ResponseAPI("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.ResponseTransactions(transactions)
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
		errorMessage := gin.H{"errors": "campaign not found."}

		res := helpers.ResponseAPI("Failed to create a transaction.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user := c.MustGet("currentUser").(structs.User)
	request.User = user
	transaction, err := h.transactionService.CreateTransaction(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		res := helpers.ResponseAPI("Something went wrong.", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := structs.CreateResponseTransaction(transaction)

	res := helpers.ResponseAPI("Transaction successfully created.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}