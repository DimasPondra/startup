package structs

import "time"

type CampaignImage struct {
	ID         int
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CampaignID int
	FileID     int
	File       File
}
