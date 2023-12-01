package structs

import (
	"os"
	"strings"
)

type campaignSummaryResponse struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Slug             string  `json:"slug"`
	ShortDescription string  `json:"short_description"`
	GoalAmount       int     `json:"goal_amount"`
	CurrentAmount    int     `json:"current_amount"`
	ImageURL         *string `json:"image_url"`
}

func CampaignsSummaryResponse(campaigns []Campaign) []campaignSummaryResponse {
	listCampaigns := []campaignSummaryResponse{}

	for _, campaign := range campaigns {
		image_url := getImageUrl(campaign.CampaignImages)

		campaignFormatter := campaignSummaryResponse{
			ID:               campaign.ID,
			Name:             campaign.Name,
			Slug:             campaign.Slug,
			ShortDescription: campaign.ShortDescription,
			GoalAmount:       campaign.GoalAmount,
			CurrentAmount:    campaign.CurrentAmount,
			ImageURL:         image_url,
		}

		listCampaigns = append(listCampaigns, campaignFormatter)
	}

	return listCampaigns
}

func getImageUrl(campaignImages []CampaignImage) *string {
	appUrl := os.Getenv("APP_URL")

	if len(campaignImages) > 0 {
		FileName := appUrl + "images/" + campaignImages[0].File.Location + "/" + campaignImages[0].File.Name
		return &FileName
	}

	return nil
}

type campaignDetailResponse struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	GoalAmount       int                     `json:"goal_amount"`
	CurrentAmount    int                     `json:"current_amount"`
	Perks            []string                `json:"perks"`
	BackerCount      int                     `json:"backer_count"`
	User             campaignUserResponse    `json:"user"`
	Images           []campaignImageResponse `json:"images"`
}

type campaignUserResponse struct {
	Name     string  `json:"name"`
	ImageURL *string `json:"image_url"`
}

type campaignImageResponse struct {
	ImageURL  string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignResponse(campaign Campaign) campaignDetailResponse {
	appUrl := os.Getenv("APP_URL")
	images := []campaignImageResponse{}

	user := campaignUserResponse{
		Name:     campaign.User.Name,
		ImageURL: nil,
	}

	if campaign.User.FileID != nil {
		avatarUrl := appUrl + "images/" + campaign.User.File.Location + "/" + campaign.User.File.Name
		user.ImageURL = &avatarUrl
	}

	for _, image := range campaign.CampaignImages {
		isPrimary := image.IsPrimary != 0

		campaignImage := campaignImageResponse{
			ImageURL:  appUrl + "images/" + image.File.Location + "/" + image.File.Name,
			IsPrimary: isPrimary,
		}

		images = append(images, campaignImage)
	}

	formatter := campaignDetailResponse{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Perks:            splitPerks(campaign.Perks),
		BackerCount:      campaign.BackerCount,
		User:             user,
		Images:           images,
	}

	return formatter
}

func splitPerks(perks string) []string {
	results := []string{}

	for _, perk := range strings.Split(perks, ",") {
		results = append(results, strings.TrimSpace(perk))
	}

	return results
}

type campaignTransactionResponse struct {
	ID        int                             `json:"id"`
	Amount    int                             `json:"amount"`
	Status    string                          `json:"status"`
	CreatedAt string                          `json:"created_at"`
	User      campaignTransactionUserResponse `json:"user"`
}

type campaignTransactionUserResponse struct {
	Name     string  `json:"name"`
	ImageURL *string `json:"image_url"`
}

func CampaignTransactionsResponse(transactions []Transaction) []campaignTransactionResponse {
	listTransactions := []campaignTransactionResponse{}
	appUrl := os.Getenv("APP_URL")

	for _, transaction := range transactions {
		if transaction.Status == "paid" {
			createdAtFormatted := transaction.CreatedAt.Format("Monday 02, January 2006")

			transactionFormatter := campaignTransactionResponse{
				ID:        transaction.ID,
				Amount:    transaction.Amount,
				Status:    transaction.Status,
				CreatedAt: createdAtFormatted,
			}

			user := campaignTransactionUserResponse{
				Name:     transaction.User.Name,
				ImageURL: nil,
			}

			if transaction.User.File.ID != 0 {
				filename := appUrl + "images/" + transaction.User.File.Location + "/" + transaction.User.File.Name
				user.ImageURL = &filename
			}

			transactionFormatter.User = user

			listTransactions = append(listTransactions, transactionFormatter)
		}
	}

	return listTransactions
}
