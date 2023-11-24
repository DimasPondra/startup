package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type TransactionService interface {
	GetTransactions(userID int, campaignID int) ([]structs.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository) *transactionService {
	return &transactionService{transactionRepo}
}

func (s *transactionService) GetTransactions(userID int, campaignID int) ([]structs.Transaction, error) {
	if userID != 0 {
		transactions, err := s.transactionRepo.FindTransactionsByUserID(userID)

		if err != nil {
			return transactions, err
		}

		return transactions, nil
	}

	if campaignID != 0 {
		transactions, err := s.transactionRepo.FindTransactionsByCampaignID(campaignID)

		if err != nil {
			return transactions, err
		}

		return transactions, nil
	}

	transactions, err := s.transactionRepo.FindAll()

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}