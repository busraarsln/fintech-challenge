package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type AccountRepository interface {
	AddAccount(account *models.Account) (int, error)
	GetAccounts(customerId int) (*models.AccountList, error)
	GetAccount(id int) (*models.Account, error)
	GetAccountByIban(iban string) (*models.Account, error)
	DeleteAccount(id int) error
}
