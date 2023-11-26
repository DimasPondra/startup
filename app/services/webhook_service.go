package services

import (
	"startup/app/structs"
)

type WebhookService interface {
	MidtransNotification(request structs.PaymentNotificationRequest) error
}

type webhookService struct {
	transactionService TransactionService
	campaignService CampaignService
}

func NewWebhookService(transactionService TransactionService, campaignService CampaignService) *webhookService {
	return &webhookService{transactionService, campaignService}
}

func (s *webhookService) MidtransNotification(request structs.PaymentNotificationRequest) error {
	orderID := request.OrderID
	transactionStatus := request.TransactionStatus
	paymentType := request.PaymentType
	fraudStatus := request.FraudStatus

	var status string

	if transactionStatus == "capture" && paymentType == "credit_card" && fraudStatus == "accept" {
		status = "paid"
	} else if transactionStatus == "settlement" {
		status = "paid"
	} else if transactionStatus == "deny" || transactionStatus == "expire" || transactionStatus == "cancel" {
		status = "canceled"
	}

	transaction, err := s.transactionService.GetTransactionByCode(orderID)
	if err != nil {
		return err
	}

	transaction.Status = status
	_, err = s.transactionService.UpdateTransaction(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignService.GetCampaignByID(transaction.CampaignID)
	if err != nil {
		return err
	}

	if status == "paid" {
		campaign.BackerCount += 1
		campaign.CurrentAmount += transaction.Amount

		err = s.campaignService.UpdateCampaignFromWebhook(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}