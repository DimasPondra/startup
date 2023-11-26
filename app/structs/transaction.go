package structs

import (
	"time"
)

type Transaction struct {
	ID         	int
	Amount     	int
	Status     	string
	Code       	string
	PaymentURL	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	CampaignID 	int
	UserID     	int
	Campaign    Campaign
	User		User
}