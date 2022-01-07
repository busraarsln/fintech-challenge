package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type BalanceRepository interface {
	GetBalance(accountId int) (*models.Balance, error)
	AddBalance(balance *models.Balance) (id int, err error)
	UpdateBalance(balance *models.Balance) error
}
