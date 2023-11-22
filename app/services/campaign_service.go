package services

import (
	"startup/app/repositories"
	"startup/app/structs"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]structs.Campaign, error)
	GetCampaignBySlug(slug string) (structs.Campaign, error)
	CreateCampaign(request structs.CampaignStoreRequest) (structs.Campaign, error)
	GetCampaignByName(name string) (structs.Campaign, error)
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

func (s *campaignService) CreateCampaign(request structs.CampaignStoreRequest) (structs.Campaign, error) {
	slug := slug.Make(request.Name)
	
	campaign := structs.Campaign{
		Name: request.Name,
		Slug: slug,
		ShortDescription: request.ShortDescription,
		Description: request.Description,
		GoalAmount: request.GoalAmount,
		CurrentAmount: 0,
		Perks: request.Perks,
		BackerCount: 0,
		UserID: request.User.ID,
	}

	newCampaign, err := s.campaignRepo.Create(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *campaignService) GetCampaignByName(name string) (structs.Campaign, error) {
	campaign, err := s.campaignRepo.FindCampaignByName(name)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}