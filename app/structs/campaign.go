package structs

import (
	"os"
	"strings"
	"time"
)

type Campaign struct {
	ID               	int
	Name             	string
	Slug             	string
	ShortDescription 	string
	Description      	string
	GoalAmount       	int
	CurrentAmount    	int
	Perks            	string
	BackerCount      	int
	CreatedAt        	time.Time
	UpdatedAt 		 	time.Time
	UserID			 	int
	User				User
	CampaignImages		[]CampaignImage
}

type CampaignStoreRequest struct {
	Name 				string 		`json:"name" binding:"required"`
	ShortDescription 	string 		`json:"short_description" binding:"required"`
	Description 		string 		`json:"description" binding:"required"`
	GoalAmount 			int 		`json:"goal_amount" binding:"required,number,gt=0"`
	Perks 				string 		`json:"perks" binding:"required"`
	User          	 	User
}

type CampaignUpdateRequest struct {
	Name 				string 		`json:"name" binding:"required"`
	ShortDescription 	string 		`json:"short_description" binding:"required"`
	Description 		string 		`json:"description" binding:"required"`
	GoalAmount 			int 		`json:"goal_amount" binding:"required,number,gt=0"`
	Perks 				string 		`json:"perks" binding:"required"`
}

type campaignResponse struct {
	ID 				 	int 		`json:"id"`
	Name 			 	string 		`json:"name"`
	Slug 			 	string 		`json:"slug"`
	ShortDescription 	string 		`json:"short_description"`
	GoalAmount		 	int			`json:"goal_amount"`
	CurrentAmount	 	int			`json:"current_amount"`
	ImageURL			*string		`json:"image_url"`
}

type detailCampaignResponse struct {
	ID 					int 					`json:"id"`
	Name 				string 					`json:"name"`
	ShortDescription 	string					`json:"short_description"`
	Description 		string					`json:"description"`
	GoalAmount 			int 					`json:"goal_amount"`
	CurrentAmount 		int 					`json:"current_amount"`
	Perks 				[]string 				`json:"perks"`
	BackerCount 		int 					`json:"backer_count"`
	User 				campaignUserResponse	`json:"user"`
	Images 				[]campaignImageResponse	`json:"images"`
}

type campaignUserResponse struct {
	Name		string	`json:"name"`
	ImageURL	string	`json:"image_url"`
}

type campaignImageResponse struct {
	ImageURL	string	`json:"url"`
	IsPrimary	bool	`json:"is_primary"`
}

func response(campaign Campaign) campaignResponse {
	image_url := getImageUrl(campaign.CampaignImages)
	
	formatter := campaignResponse{
		ID: campaign.ID,
		Name: campaign.Name,
		Slug: campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		ImageURL: image_url,
	}

	return formatter
}

func ResponseCampaigns(campaigns []Campaign) []campaignResponse {
	listCampaigns := []campaignResponse{}

	for _, campaign := range campaigns {
		campaignFormatter := response(campaign)
		listCampaigns = append(listCampaigns, campaignFormatter) 
	}

	return listCampaigns
}

func ResponseCampaign(campaign Campaign) detailCampaignResponse {
	appUrl := os.Getenv("APP_URL")
	images := []campaignImageResponse{}

	user := campaignUserResponse{
		Name: campaign.User.Name,
		ImageURL: appUrl + campaign.User.AvatarFileName,
	}

	for _, image := range campaign.CampaignImages {
		isPrimary := image.IsPrimary != 0

		campaignImage := campaignImageResponse{
			ImageURL: appUrl + image.FileName,
			IsPrimary: isPrimary,
		}

		images = append(images, campaignImage)
	}

	campaignFormatter := detailCampaignResponse{
		ID: campaign.ID,
		Name: campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description: campaign.Description,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Perks: splitPerks(campaign.Perks),
		BackerCount: campaign.BackerCount,
		User: user,
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