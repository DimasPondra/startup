package services

import (
	"os"
	"startup/app/repositories"
	"startup/app/structs"
)

type CampaignImageService interface {
	SaveImage(filename string, isPrimary bool, campaign structs.Campaign) (structs.CampaignImage, error)
	DeleteImages(campaignID int) (bool, error)
}

type campaignImageService struct {
	campaignImageRepo repositories.CampaignImageRepository
}

func NewCampaignImageService(campaignImageRepo repositories.CampaignImageRepository) *campaignImageService {
	return &campaignImageService{campaignImageRepo}
}

func (s *campaignImageService) SaveImage(filename string, isPrimary bool, campaign structs.Campaign) (structs.CampaignImage, error) {
	var campaignImage structs.CampaignImage
	primary := 1

	if !isPrimary {
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

func (s *campaignImageService) DeleteImages(campaignID int) (bool, error) {
	images, _ := s.campaignImageRepo.FindImagesByCampaignID(campaignID)

	if len(images) > 0 {
		for _, image := range images {
			os.Remove(image.FileName)
		}

		_, err := s.campaignImageRepo.DeleteAllImages(images)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
}
