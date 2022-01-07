package account

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_ "github.com/go-sql-driver/mysql"
)

type (
	mysqlRepo struct{ db *sql.DB }
)

//NewMysqlRepository creates a new repo
func NewMysqlRepository(db *sql.DB) repository.AccountRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) AddAccount(account *models.Account) (int, error) {
	stmt, e := m.db.Prepare(createAccount)
	if e != nil {
		return 0, errors.New(fmt.Sprintln("Sql syntax error", e.Error()))
	}
	defer stmt.Close()

	//execute
	res, e := stmt.Exec(account.Type,
		account.AccountNo,
		account.Iban,
		account.Currency,
		account.Description,
		account.Status,
		account.Nickname,
		account.CustomerID,
		account.BalanceID)

	if e != nil {
		return 0, errors.New(fmt.Sprintln("Account could not be created", e.Error()))
	}
	id, e := res.LastInsertId()

	if e != nil {
		return 0, errors.New(fmt.Sprintln("Cannot get lastInsertId", e.Error()))
	}

	return int(id), nil
}

func (m *mysqlRepo) GetAccounts(customerId int) (*models.AccountList, error) {

	var accountList = make([]*models.Account, 0)
	rows, err := m.db.Query(getAccounts, customerId)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error getting account", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		n := models.Account{}
		if err = rows.Scan(&n.ID,
			&n.Type,
			&n.AccountNo,
			&n.Iban,
			&n.Currency,
			&n.Description,
			&n.Nickname,
			&n.CreatedAt,
			&n.UpdatedAt,
			&n.IsActive,
			&n.CustomerID,
			&n.BalanceID,
			&n.Status,
			&n.Balance.ID,
			&n.Balance.CreditLimit,
			&n.Balance.AmountId,
			&n.Balance.Amount.ID,
			&n.Balance.Amount.Currency,
			&n.Balance.Amount.Value,
		); err != nil {
			return nil, errors.New(fmt.Sprintln("Account not found", err.Error()))
		}
		accountList = append(accountList, &n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &models.AccountList{
		TotalCount: 0,
		TotalPages: 0,
		Page:       0,
		Size:       1,
		Accounts:   accountList,
	}, nil
}

func (m *mysqlRepo) DeleteAccount(id int) error {

	stmt, e := m.db.Prepare(deleteAccount)
	if e != nil {
		return errors.New(fmt.Sprintln("Sql syntax error", e.Error()))
	}
	defer stmt.Close()

	//execute
	res, e := stmt.Exec(id)
	if e != nil {
		return errors.New("account could not be deleted")
	}
	raff, e := res.RowsAffected()

	if e != nil || raff < 1 {
		return errors.New("account could not be deleted")
	}

	return nil
}

func (m *mysqlRepo) GetAccount(id int) (*models.Account, error) {

	var account models.Account

	smtp, err := m.db.Prepare(getAccount)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(id).Scan(&account.ID,
		&account.Type,
		&account.AccountNo,
		&account.Iban,
		&account.Currency,
		&account.Description,
		&account.Nickname,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.IsActive,
		&account.CustomerID,
		&account.BalanceID,
		&account.Status,
		&account.Balance.ID,
		&account.Balance.CreditLimit,
		&account.Balance.AmountId,
		&account.Balance.Amount.ID,
		&account.Balance.Amount.Currency,
		&account.Balance.Amount.Value,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get account", err.Error()))
	}

	return &account, nil
}

func (m *mysqlRepo) GetAccountByIban(iban string) (*models.Account, error) {

	var account models.Account

	smtp, err := m.db.Prepare(getAccountByIban)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("get account by iban prepare error", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(iban).Scan(&account.ID,
		&account.Type,
		&account.AccountNo,
		&account.Iban,
		&account.Currency,
		&account.Description,
		&account.Nickname,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.IsActive,
		&account.CustomerID,
		&account.BalanceID,
		&account.Status,
		&account.Balance.ID,
		&account.Balance.CreditLimit,
		&account.Balance.AmountId,
		&account.Balance.Amount.ID,
		&account.Balance.Amount.Currency,
		&account.Balance.Amount.Value,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Account not found", err.Error()))
	}

	return &account, nil
}
