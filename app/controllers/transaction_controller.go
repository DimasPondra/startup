package controllers

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
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

	transactions, err := h.transactionService.GetTransactions(userID)
	if err != nil {
		res := helpers.ResponseAPI("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helpers.ResponseAPI("List of transactions.", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, res)
}