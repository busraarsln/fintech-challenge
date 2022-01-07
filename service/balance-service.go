package service

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type BalanceService interface {
	Validate(balance *models.Balance) error
	GetBalance(accountId int) (*models.Balance, error)
	UpdateBalance(accountId int, balance *models.Balance) error
}
