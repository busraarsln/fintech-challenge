package balance

import (
	"errors"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_service "github.com/busraarsln/fintech-challenge/service"
)

type BalanceService interface {
	GetBalance() (*models.Balance, error)
}

type service struct{}

var (
	repo repository.BalanceRepository
)

func (*service) Validate(balance *models.Balance) error {
	if balance == nil {
		err := errors.New("the balance is empty")
		return err
	}
	if balance.Amount.Currency == "" {
		err := errors.New("the currency is empty")
		return err
	}
	return nil
}

func NewBalanceService(repository repository.BalanceRepository) _service.BalanceService {
	repo = repository
	return &service{}
}

func (*service) GetBalance(accountId int) (*models.Balance, error) {
	return repo.GetBalance(accountId)
}

func (*service) UpdateBalance(accountId int, balance *models.Balance) error {
	b, err := repo.GetBalance(accountId)
	balance.Amount.ID = b.Amount.ID
	if err != nil {
		return err
	}

	return repo.UpdateBalance(balance)
}
