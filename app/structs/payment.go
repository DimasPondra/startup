package structs

type PaymentStoreRequest struct {
	Code   string
	Amount int
	Name   string
	Email  string
}