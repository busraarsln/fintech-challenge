package payment

import (
	"errors"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_service "github.com/busraarsln/fintech-challenge/service"
)

type service struct{}

var (
	paymentRepo repository.PaymentRepository
)

func NewPaymentService(paymentRepository repository.PaymentRepository) _service.PaymentService {
	paymentRepo = paymentRepository
	return &service{}
}

func (*service) Validate(payment *models.Payment) error {
	if payment == nil {
		err := errors.New("the payment is empty")
		return err
	}
	if payment.From == "" {
		err := errors.New("the payment is emptdy")
		return err
	}
	return nil
}
func (*service) AddPayment(payment *models.Payment) (*models.Payment, error) {

	paymentId, err := paymentRepo.AddPayment(payment)
	if err != nil {
		return nil, err
	}
	return paymentRepo.GetPayment(paymentId)
}
func (*service) GetPayments(acountId int) (*models.PaymentList, error) {
	return paymentRepo.GetPayments(acountId)
}
