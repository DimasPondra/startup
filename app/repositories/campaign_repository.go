package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]structs.Campaign, error)
	FindByUserID(userID int) ([]structs.Campaign, error)
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

func (r *campaignRepository) FindByUserID(userID int) ([]structs.Campaign, error) {
	var campaigns []structs.Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}