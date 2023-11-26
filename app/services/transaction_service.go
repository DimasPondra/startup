package services

import (
	"startup/app/helpers"
	"startup/app/repositories"
	"startup/app/structs"
)

type TransactionService interface {
	GetTransactions(userID int, campaignID int) ([]structs.Transaction, error)
	GetTransactionsByUserID(userID int) ([]structs.Transaction, error)
	GetTransactionByCode(code string) (structs.Transaction, error)
	CreateTransaction(request structs.TransactionStoreRequest) (structs.Transaction, error)
	UpdateTransaction(transaction structs.Transaction) (structs.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	paymentService PaymentService
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, paymentService PaymentService) *transactionService {
	return &transactionService{transactionRepo, paymentService}
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

func (s *transactionService) GetTransactionsByUserID(userID int) ([]structs.Transaction, error) {
	transactions, err := s.transactionRepo.FindTransactionsByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionByCode(code string) (structs.Transaction, error) {
	transaction, err := s.transactionRepo.FindTransactionByCode(code)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *transactionService) CreateTransaction(request structs.TransactionStoreRequest) (structs.Transaction, error) {
	code := helpers.GenerateRandomCode()
	
	transaction := structs.Transaction{
		Amount: request.Amount,
		Status: "pending",
		Code: code,
		CampaignID: request.CampaignID,
		UserID: request.User.ID,
	}

	newTransaction, err := s.transactionRepo.Create(transaction)
	if err != nil {
		return newTransaction, err
	}

	payment := structs.PaymentStoreRequest{
		Code: newTransaction.Code,
		Amount: newTransaction.Amount,
		Name: request.User.Name,
		Email: request.User.Email,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(payment)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	updatedTransaction, err := s.transactionRepo.Update(newTransaction)
	if err != nil {
		return updatedTransaction, err
	}

	return updatedTransaction, nil
}

func (s *transactionService) UpdateTransaction(transaction structs.Transaction) (structs.Transaction, error) {
	transaction, err := s.transactionRepo.Update(transaction)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}