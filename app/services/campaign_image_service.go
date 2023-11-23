package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type CampaignImageService interface {
	SaveImage(filename string, isPrimary bool, campaign structs.Campaign) (structs.CampaignImage, error)
}

type campaignImageService struct {
	campaignImageRepo repositories.CampaignImageRepository
}

func NewCampaignImageService(campaignImageRepo repositories.CampaignImageRepository) *campaignImageService {
	return &campaignImageService{campaignImageRepo}
}

func (s *campaignImageService) SaveImage(filename string, isPrimary bool, campaign structs.Campaign) (structs.CampaignImage, error) {
	var campaignImage structs.CampaignImage
	var primary int

	if isPrimary {
		primary = 1
	} else {
		primary = 0
	}

	campaignImage.FileName = filename
	campaignImage.CampaignID = campaign.ID
	campaignImage.IsPrimary = primary

	newCampaignImage, err := s.campaignImageRepo.Create(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}