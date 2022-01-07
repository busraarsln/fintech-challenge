package account

import (
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_service "github.com/busraarsln/fintech-challenge/service"
	"github.com/google/uuid"
)

type service struct{}

var (
	accountRepo repository.AccountRepository
	balanceRepo repository.BalanceRepository
	amountRepo  repository.AmountRepository
)

func NewAccountService(
	accountRepository repository.AccountRepository,
	balanceRepository repository.BalanceRepository,
	amountRepository repository.AmountRepository) _service.AccountService {
	accountRepo = accountRepository
	balanceRepo = balanceRepository
	amountRepo = amountRepository
	return &service{}
}

func (*service) Validate(account *models.Account) error {
	if account == nil {
		err := errors.New("the account is empty")
		return err
	}
	if account.Iban == "" {
		err := errors.New("the account is emptdy")
		return err
	}
	return nil
}
func (*service) AddAccount(account *models.Account) (*models.Account, error) {
	amountId, err := amountRepo.AddAmount(&account.Balance.Amount)
	if err != nil {
		return nil, err
	}
	account.Balance.AmountId = amountId
	balanceId, err := balanceRepo.AddBalance(&account.Balance)
	if err != nil {
		return nil, err
	}
	account.BalanceID = balanceId

	account.Iban = uuid.NewString()
	account.AccountNo = account.Iban[len(account.Iban)-16:]

	accountId, err := accountRepo.AddAccount(account)
	if err != nil {
		return nil, err
	}
	return accountRepo.GetAccount(accountId)

}
func (*service) GetAccounts(customerId int) (*models.AccountList, error) {
	return accountRepo.GetAccounts(customerId)
}

func (*service) DeleteAccount(id int) error {
	account, err := accountRepo.GetAccount(id)
	if err != nil {
		return err
	}
	if int(account.Balance.Amount.Value) != 0 {
		return errors.New(fmt.Sprintln("Balance must be 0"))
	}
	return accountRepo.DeleteAccount(id)
}
