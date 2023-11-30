package structs

import "os"

type transactionSummaryResponse struct {
	ID         int                         `json:"id"`
	Amount     int                         `json:"amount"`
	Status     string                      `json:"status"`
	PaymentURL *string                     `json:"payment_url"`
	CreatedAt  string                      `json:"created_at"`
	Campaign   transactionCampaignResponse `json:"campaign"`
}

type transactionCampaignResponse struct {
	Name     string  `json:"name"`
	ImageURL *string `json:"image_url"`
}

func TransactionsSummaryResponse(transactions []Transaction) []transactionSummaryResponse {
	listTransactions := []transactionSummaryResponse{}

	appUrl := os.Getenv("APP_URL")
	var campaignImage *string

	for _, transaction := range transactions {
		createdAtFormatted := transaction.CreatedAt.Format("Monday 02, January 2006")

		transactionFormatter := transactionSummaryResponse{
			ID:         transaction.ID,
			Amount:     transaction.Amount,
			Status:     transaction.Status,
			PaymentURL: nil,
			CreatedAt:  createdAtFormatted,
		}

		if transaction.Status == "pending" {
			transactionFormatter.PaymentURL = &transaction.PaymentURL
		}

		if len(transaction.Campaign.CampaignImages) > 0 {
			file := transaction.Campaign.CampaignImages[0].File
			filename := appUrl + "images/" + file.Location + "/" + file.Name
			campaignImage = &filename
		}

		campaign := transactionCampaignResponse{
			Name:     transaction.Campaign.Name,
			ImageURL: campaignImage,
		}

		transactionFormatter.Campaign = campaign

		listTransactions = append(listTransactions, transactionFormatter)
	}

	return listTransactions
}

type transactionResponse struct {
	ID         int    `json:"id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func TransactionResponse(transaction Transaction) transactionResponse {
	formatter := transactionResponse{
		ID:         transaction.ID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}

	return formatter
}
