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
	UserID			 int
	User 			 User
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

type CampaignStoreRequest struct {
	Name 				string 		`json:"name" binding:"required"`
	ShortDescription 	string 		`json:"short_description" binding:"required"`
	Description 		string 		`json:"description" binding:"required"`
	GoalAmount 			int 		`json:"goal_amount" binding:"required,number,gt=0"`
	Perks 				string 		`json:"perks" binding:"required"`
	User          	 	User
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
	ID 					int 					`json:"id"`
	Name 				string 					`json:"name"`
	ShortDescription 	string					`json:"short_description"`
	Description 		string					`json:"description"`
	GoalAmount 			int 					`json:"goal_amount"`
	CurrentAmount 		int 					`json:"current_amount"`
	Perks 				[]string 				`json:"perks"`
	BackerCount 		int 					`json:"backer_count"`
	User 				campaignUserResponse	`json:"user"`
	Images 				[]campaignImageResponse `json:"images"`
}

type campaignUserResponse struct {
	Name		string	`json:"name"`
	ImageURL	string	`json:"image_url"`
}

type campaignImageResponse struct {
	ImageURL	string	`json:"image_url"`
	IsPrimary	bool	`json:"is_primary"`
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

	campaignFormatter := campaignResponse{
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

func CampaignStoreResponse(campaign Campaign) listCampaignResponse {
	formatter := listCampaignResponse{
		ID: campaign.ID,
		Name: campaign.Name,
		Slug: campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		UserID: campaign.UserID,
		ImageURL: getImageUrl(campaign.CampaignImages),
	}

	return formatter
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