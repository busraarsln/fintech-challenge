package service

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type AccountService interface {
	Validate(account *models.Account) error
	AddAccount(account *models.Account) (*models.Account, error)
	GetAccounts(customerId int) (*models.AccountList, error)
	DeleteAccount(id int) error
}
