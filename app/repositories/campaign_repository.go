package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]structs.Campaign, error)
	FindCampaignsByUserID(userID int) ([]structs.Campaign, error)
	FindCampaignBySlug(slug string) (structs.Campaign, error)
	FindCampaignByName(name string) (structs.Campaign, error)
	FindCampaignByID(ID int) (structs.Campaign, error)
	Create(campaign structs.Campaign) (structs.Campaign, error)
	Update(campaign structs.Campaign) (structs.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) FindAll() ([]structs.Campaign, error) {
	var campaigns []structs.Campaign

	err := r.db.Preload("CampaignImages", "is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindCampaignsByUserID(userID int) ([]structs.Campaign, error) {
	var campaigns []structs.Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindCampaignBySlug(slug string) (structs.Campaign, error) {
	var campaign structs.Campaign

	err := r.db.Where("slug = ?", slug).Preload("CampaignImages").Preload("User").First(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) FindCampaignByName(name string) (structs.Campaign, error) {
	var campaign structs.Campaign

	err := r.db.Where("name = ?", name).First(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) FindCampaignByID(ID int) (structs.Campaign, error) {
	var campaign structs.Campaign

	err := r.db.First(&campaign, ID).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) Create(campaign structs.Campaign) (structs.Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) Update(campaign structs.Campaign) (structs.Campaign, error) {
	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}