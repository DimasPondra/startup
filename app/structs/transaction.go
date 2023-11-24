package structs

import "time"

type Transaction struct {
	ID         	int
	Amount     	int
	Status     	string
	Code       	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	CampaignID 	int
	UserID     	int
}