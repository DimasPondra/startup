package structs

import (
	"time"
)

type Campaign struct {
	ID               int
	Name             string
	Slug             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UserID           int
	User             User
	CampaignImages   []CampaignImage
}
