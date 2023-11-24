package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type CampaignImageRepository interface {
	Create(campaignImage structs.CampaignImage) (structs.CampaignImage, error)
	FindImagesByCampaignID(campaignID int) ([]structs.CampaignImage, error)
	DeleteAllImages(campaignImages []structs.CampaignImage) (bool, error)
}

type campaignImageRepository struct {
	db *gorm.DB
}

func NewCampaignImageRepository(db *gorm.DB) *campaignImageRepository {
	return &campaignImageRepository{db}
}

func (r *campaignImageRepository) Create(campaignImage structs.CampaignImage) (structs.CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error

	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *campaignImageRepository) FindImagesByCampaignID(campaignID int) ([]structs.CampaignImage, error) {
	var campaignImages []structs.CampaignImage

	err := r.db.Where("campaign_id = ?", campaignID).Find(&campaignImages).Error

	if err != nil {
		return campaignImages, err
	}

	return campaignImages, nil
}

func (r *campaignImageRepository) DeleteAllImages(campaignImages []structs.CampaignImage) (bool, error) {
	err := r.db.Delete(&campaignImages).Error

	if err != nil {
		return false, err
	}

	return true, nil
}