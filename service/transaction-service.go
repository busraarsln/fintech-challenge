package service

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type TransactionService interface {
	AddTransaction(account *models.Transaction) (*models.Transaction, error)
	GetTransactions(accountId int) (*models.TransactionList, error)
	MakeTransaction(transaction *models.Transaction) error
	Validate(transaction *models.Transaction) error
}
