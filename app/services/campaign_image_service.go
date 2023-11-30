package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type CampaignImageService interface {
	SaveImage(requets structs.CampaignImageStoreRequest) (structs.CampaignImage, error)
	DeleteImages(campaignID int) (bool, error)
}

type campaignImageService struct {
	campaignImageRepo repositories.CampaignImageRepository
}

func NewCampaignImageService(campaignImageRepo repositories.CampaignImageRepository) *campaignImageService {
	return &campaignImageService{campaignImageRepo}
}

func (s *campaignImageService) SaveImage(request structs.CampaignImageStoreRequest) (structs.CampaignImage, error) {
	campignImage := structs.CampaignImage{
		IsPrimary:  request.IsPrimary,
		CampaignID: request.CampaignID,
		FileID:     request.FileID,
	}

	newCampaignImage, err := s.campaignImageRepo.Create(campignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}

func (s *campaignImageService) DeleteImages(campaignID int) (bool, error) {
	images, _ := s.campaignImageRepo.FindImagesByCampaignID(campaignID)

	if len(images) > 0 {
		_, err := s.campaignImageRepo.DeleteAllImages(images)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
}
