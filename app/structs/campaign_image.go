package structs

import "time"

type CampaignImage struct {
	ID         int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CampaignID int
}