package structs

import "os"

type transactionSummaryResponse struct {
	ID        int                         `json:"id"`
	Amount    int                         `json:"amount"`
	Status    string                      `json:"status"`
	CreatedAt string                      `json:"created_at"`
	Campaign  transactionCampaignResponse `json:"campaign"`
	User      transactionUserResponse     `json:"user"`
}

type transactionCampaignResponse struct {
	Name     string  `json:"name"`
	ImageURL *string `json:"image_url"`
}

type transactionUserResponse struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TransactionsSummaryResponse(transactions []Transaction) []transactionSummaryResponse {
	listTransactions := []transactionSummaryResponse{}

	appUrl := os.Getenv("APP_URL")
	var campaignImage *string

	for _, transaction := range transactions {
		createdAtFormatted := transaction.CreatedAt.Format("Monday 02, January 2006")

		transactionFormatter := transactionSummaryResponse{
			ID:        transaction.ID,
			Amount:    transaction.Amount,
			Status:    transaction.Status,
			CreatedAt: createdAtFormatted,
		}

		if len(transaction.Campaign.CampaignImages) > 0 {
			filename := appUrl + transaction.Campaign.CampaignImages[0].FileName
			campaignImage = &filename
		}

		campaign := transactionCampaignResponse{
			Name: transaction.Campaign.Name,
			ImageURL: campaignImage,
		}

		user := transactionUserResponse{
			Name: transaction.User.Name,
			ImageURL: "",
		}

		if transaction.User.AvatarFileName != "" {
			user.ImageURL = appUrl + transaction.User.AvatarFileName
		}

		transactionFormatter.Campaign = campaign
		transactionFormatter.User = user

		listTransactions = append(listTransactions, transactionFormatter)
	}

	return listTransactions
}

type transactionResponse struct {
	ID			int		`json:"id"`
	Amount		int		`json:"amount"`
	Status		string	`json:"status"`
	Code		string	`json:"code"`
	PaymentURL	string	`json:"payment_url"`
}

func TransactionResponse(transaction Transaction) transactionResponse {
	formatter := transactionResponse{
		ID: transaction.ID,
		Amount: transaction.Amount,
		Status: transaction.Status,
		Code: transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}

	return formatter
}