package service

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type PaymentService interface {
	AddPayment(account *models.Payment) (*models.Payment, error)
	GetPayments(accountId int) (*models.PaymentList, error)
	Validate(payment *models.Payment) error
}
