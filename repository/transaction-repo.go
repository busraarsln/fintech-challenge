package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type TransactionRepository interface {
	AddTransaction(transaction *models.Transaction) (int, error)
	GetTransactions(accountId int) (*models.TransactionList, error)
	GetTransaction(transactionId int) (*models.Transaction, error)
	MakeTransaction(transaction *models.Transaction, targetAccount *models.Account, sourceAccount *models.Account) error
}
