package structs

import (
	"os"
	"strings"
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

type campaignResponse struct {
	ID 				int 		`json:"id"`
	Name 			string 		`json:"name"`
	Description 	string		`json:"description"`
	GoalAmount 		int 		`json:"goal_amount"`
	CurrentAmount 	int 		`json:"current_amount"`
	Perks 			[]string 	`json:"perks"`
	BackerCount 	int 		`json:"backer_count"`
	Images 			[]string 	`json:"images"`
}

func CampaignResponses(campaigns []Campaign) []listCampaignResponse {
	listCampaigns := []listCampaignResponse{}

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

func CampaignResponse(campaign Campaign) campaignResponse {
	appUrl := os.Getenv("APP_URL")
	images := []string{}

	for _, image := range campaign.CampaignImages {
		getFileName := appUrl + image.FileName

		images = append(images, getFileName)
	}

	campaignFormatter := campaignResponse{
		ID: campaign.ID,
		Name: campaign.Name,
		Description: campaign.Description,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Perks: splitPerks(campaign.Perks),
		BackerCount: campaign.BackerCount,
		Images: images,
	}

	return campaignFormatter
}

func getImageUrl(campaignImages []CampaignImage) *string {
	appUrl := os.Getenv("APP_URL")

	if len(campaignImages) > 0 {
		FileName := appUrl + campaignImages[0].FileName
		return &FileName
	}

	return nil
}

func splitPerks(perks string) []string {
	splitOfPerks := strings.Split(perks, ", ")

	return splitOfPerks
}