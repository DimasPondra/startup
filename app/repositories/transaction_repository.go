package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]structs.Transaction, error)
	FindTransactionsByUserID(userID int) ([]structs.Transaction, error)
	FindTransactionsByCampaignID(campaignID int) ([]structs.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindAll() ([]structs.Transaction, error) {
	var transactions []structs.Transaction

	err := r.db.Preload("Campaign.CampaignImages", "is_primary = 1").Preload("User").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindTransactionsByUserID(userID int) ([]structs.Transaction, error) {
	var transactions []structs.Transaction

	err := r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "is_primary = 1").Preload("User").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindTransactionsByCampaignID(campaignID int) ([]structs.Transaction, error) {
	var transactions []structs.Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("Campaign.CampaignImages", "is_primary = 1").Preload("User").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}