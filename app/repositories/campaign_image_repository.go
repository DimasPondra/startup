package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type CampaignImageRepository interface {
	Create(campaignImage structs.CampaignImage) (structs.CampaignImage, error)
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