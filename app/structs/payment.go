package structs

type PaymentStoreRequest struct {
	Code   string
	Amount int
	Name   string
	Email  string
}

type PaymentNotificationRequest struct {
	OrderID           string `json:"order_id" binding:"required"`
	TransactionStatus string `json:"transaction_status" binding:"required"`
	PaymentType       string `json:"payment_type" binding:"required"`
	FraudStatus       string `json:"fraud_status" binding:"-"`
}