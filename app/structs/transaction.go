package structs

import (
	"os"
	"time"
)

type Transaction struct {
	ID         	int
	Amount     	int
	Status     	string
	Code       	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	CampaignID 	int
	UserID     	int
	Campaign    Campaign
	User		User
}

type transactionResponse struct {
	ID			int							`json:"id"`
	Amount		int							`json:"amount"`
	Status		string						`json:"status"`
	CreatedAt	string						`json:"created_at"`
	Campaign	campaignTransactionResponse	`json:"campaign"`
	User		userTransactionResponse		`json:"user"`
}

type campaignTransactionResponse struct {
	Name		string	`json:"name"`
	ImageURL	*string	`json:"image_url"`
}

type userTransactionResponse struct {
	Name		string	`json:"name"`
	ImageURL	string	`json:"image_url"`
}

func responseTransaction(transaction Transaction) transactionResponse {
	createdAt := transaction.CreatedAt.Format("Monday 02, January 2006")

	formatter := transactionResponse{
		ID: transaction.ID,
		Amount: transaction.Amount,
		Status: transaction.Status,
		CreatedAt: createdAt,
	}

	return formatter
}

func ResponseTransactions(transactions []Transaction) []transactionResponse {
	listTransactions := []transactionResponse{}

	for _, trx := range transactions {
		trxFormatter := responseTransaction(trx)

		var campaignImage *string
		appUrl := os.Getenv("APP_URL")

		if len(trx.Campaign.CampaignImages) > 0 {
			filename := appUrl + trx.Campaign.CampaignImages[0].FileName
			campaignImage = &filename
		}

		campaignTransaction := campaignTransactionResponse{
			Name: trx.Campaign.Name,
			ImageURL: campaignImage,
		}

		userTransaction := userTransactionResponse{
			Name: trx.User.Name,
			ImageURL: appUrl + trx.User.AvatarFileName,
		}

		trxFormatter.Campaign = campaignTransaction
		trxFormatter.User = userTransaction
		
		listTransactions = append(listTransactions, trxFormatter)
	}

	return listTransactions
}