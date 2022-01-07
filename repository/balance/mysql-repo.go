package balance

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
func NewMysqlRepository(db *sql.DB) repository.BalanceRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) GetBalance(accountId int) (*models.Balance, error) {

	balance := models.Balance{}
	smtp, err := m.db.Prepare(getBalance)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(accountId).Scan(&balance.ID,
		&balance.CreditLimit,
		&balance.CreatedAt,
		&balance.Amount.ID,
		&balance.Amount.Currency,
		&balance.Amount.Value,
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Account not found", err.Error()))
	}

	return &balance, err
}

func (m *mysqlRepo) AddBalance(balance *models.Balance) (int, error) {

	stmt, err := m.db.Prepare(createBalance)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}

	res, err := stmt.Exec(balance.CreditLimit, balance.AmountId)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Balance could not be created", err.Error()))
	}

	id, e := res.LastInsertId()

	if e != nil {
		return 0, errors.New(fmt.Sprintln("Cannot get lastInsertId", e.Error()))
	}

	return int(id), nil
}

func (m *mysqlRepo) UpdateBalance(balance *models.Balance) error {

	stmt, err := m.db.Prepare(updateBalance)

	if err != nil {
		return errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}

	res, err := stmt.Exec(balance.Amount.Value, balance.Amount.Currency, balance.Amount.ID)
	_, err2 := res.RowsAffected()
	if err != nil || err2 != nil {
		return errors.New(fmt.Sprintln("Balance could not be updated", err.Error()))
	}

	return nil
}
