package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]structs.Campaign, error)
	GetCampaignBySlug(slug string) (structs.Campaign, error)
}

type campaignService struct {
	campaignRepo repositories.CampaignRepository
}

func NewCampaignSevice(campaignRepo repositories.CampaignRepository) *campaignService {
	return &campaignService{campaignRepo}
}

func (s *campaignService) GetCampaigns(userID int) ([]structs.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.campaignRepo.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.campaignRepo.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *campaignService) GetCampaignBySlug(slug string) (structs.Campaign, error) {
	campaign, err := s.campaignRepo.FindBySlug(slug)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}