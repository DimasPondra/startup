package structs

type TransactionStoreRequest struct {
	CampaignID int `json:"campaign_id" binding:"required,number"`
	Amount     int `json:"amount" binding:"required,number,gt=0"`
	User       User
}