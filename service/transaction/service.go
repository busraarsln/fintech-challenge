package transaction

import (
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_service "github.com/busraarsln/fintech-challenge/service"
)

type service struct{}

var (
	transacitonRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
)

func NewTransactionService(transactionRepository repository.TransactionRepository,
	accountRepository repository.AccountRepository) _service.TransactionService {
	transacitonRepo = transactionRepository
	accountRepo = accountRepository
	return &service{}
}

func (*service) Validate(transaction *models.Transaction) error {
	if transaction == nil {
		err := errors.New("the transaction is empty")
		return err
	}
	if transaction.From == "" {
		err := errors.New("the transaction is empty")
		return err
	}
	return nil
}
func (*service) AddTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	transactionId, _ := transacitonRepo.AddTransaction(transaction)
	return transacitonRepo.GetTransaction(transactionId)
}
func (*service) GetTransactions(accountId int) (*models.TransactionList, error) {
	return transacitonRepo.GetTransactions(accountId)
}

func (*service) MakeTransaction(transaction *models.Transaction) error {
	sourceAccount, err := accountRepo.GetAccountByIban(transaction.From)
	if err != nil || sourceAccount == nil {
		return errors.New(fmt.Sprintln("source account not found or error getting the account", err.Error()))
	}
	if transaction.Amount.Value > sourceAccount.Balance.CreditLimit+sourceAccount.Balance.Amount.Value {
		return errors.New(fmt.Sprintln("insufficient balance"))
	}
	targetAccount, err := accountRepo.GetAccountByIban(transaction.To)
	if err != nil || targetAccount == nil {
		return errors.New(fmt.Sprintln("target account not found or error getting the account", err.Error()))
	}
	targetAccount.Balance.Amount.Value += transaction.Amount.Value
	sourceAccount.Balance.Amount.Value -= transaction.Amount.Value
	return transacitonRepo.MakeTransaction(transaction, targetAccount, sourceAccount)
}
