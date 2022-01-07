package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type PaymentRepository interface {
	AddPayment(payment *models.Payment) (int, error)
	GetPayments(accountId int) (*models.PaymentList, error)
	GetPayment(paymentId int) (*models.Payment, error)
}
