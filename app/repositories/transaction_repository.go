package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]structs.Transaction, error)
	FindTransactionsByUserID(userID int) ([]structs.Transaction, error)
	FindTransactionsByCampaignID(campaignID int) ([]structs.Transaction, error)
	FindTransactionByCode(code string) (structs.Transaction, error)
	Create(transaction structs.Transaction) (structs.Transaction, error)
	Update(transaction structs.Transaction) (structs.Transaction, error)
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

	err := r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "is_primary = 1").Preload("Campaign.CampaignImages.File").Preload("User").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindTransactionsByCampaignID(campaignID int) ([]structs.Transaction, error) {
	var transactions []structs.Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("Campaign.CampaignImages", "is_primary = 1").Preload("User.File").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindTransactionByCode(code string) (structs.Transaction, error) {
	var transaction structs.Transaction

	err := r.db.Where("code = ?", code).First(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Create(transaction structs.Transaction) (structs.Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Update(transaction structs.Transaction) (structs.Transaction, error) {
	err := r.db.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
