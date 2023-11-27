package structs

type TransactionStoreRequest struct {
	CampaignID int `json:"campaign_id" validate:"required,exists_in_campaigns"`
	Amount     int `json:"amount" validate:"required,gt=10000"`
	User       User
}