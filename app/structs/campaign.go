package structs

import (
	"os"
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
	UpdatedAt 		 time.Time
	UserID 			 int
	CampaignImages 	 []CampaignImage
}

type CampaignImage struct {
	ID 			int
	FileName 	string
	IsPrimary 	int
	CreatedAt   time.Time
	UpdatedAt 	time.Time
	CampaignID 	int
}

type listCampaignResponse struct {
	ID 				 int 		`json:"id"`
	Name 			 string 	`json:"name"`
	Slug 			 string 	`json:"slug"`
	ShortDescription string 	`json:"short_description"`
	GoalAmount		 int		`json:"goal_amount"`
	CurrentAmount	 int		`json:"current_amount"`
	UserID			 int		`json:"user_id"`
	ImageURL 		 *string 	`json:"image_url"`
}

func CampaignResponses(campaigns []Campaign) []listCampaignResponse {
	var listCampaigns []listCampaignResponse

	for _, campaign := range campaigns {
		
		campaignFormatter := listCampaignResponse{
			ID: campaign.ID,
			Name: campaign.Name,
			Slug: campaign.Slug,
			ShortDescription: campaign.ShortDescription,
			GoalAmount: campaign.GoalAmount,
			CurrentAmount: campaign.CurrentAmount,
			UserID: campaign.UserID,
			ImageURL: getImageUrl(campaign.CampaignImages),
		}

		listCampaigns = append(listCampaigns, campaignFormatter) 
	}

	return listCampaigns
}

func getImageUrl(campaignImages []CampaignImage) *string {
	appUrl := os.Getenv("APP_URL")

	if len(campaignImages) > 0 {
		FileName := appUrl + campaignImages[0].FileName
		return &FileName
	}

	return nil
}