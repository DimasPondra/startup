package services

import (
	"os"
	"startup/app/structs"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	GetPaymentURL(request structs.PaymentStoreRequest) (string, error)
}

type paymentService struct{}

func NewPaymentService() *paymentService {
	return &paymentService{}
}

func (s *paymentService) GetPaymentURL(request structs.PaymentStoreRequest) (string, error) {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")

	if os.Getenv("MIDTRANS_ENVIRONMENT") == "PRODUCTION" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.Code,
			GrossAmt: int64(request.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.Name,
			Email: request.Email,
		},
	}

	snapResp, err := snap.CreateTransaction(req)

	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil
}
