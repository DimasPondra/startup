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
}

func NewTransactionController(transactionService services.TransactionService) *transactionController {
	return &transactionController{transactionService}
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