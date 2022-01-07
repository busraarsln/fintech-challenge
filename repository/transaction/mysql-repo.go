package transaction

import (
	"context"
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
func NewMysqlRepository(db *sql.DB) repository.TransactionRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) AddTransaction(transaction *models.Transaction) (int, error) {
	stmt, err := m.db.Prepare(createTransaction)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Failed to create transaction", err.Error()))
	}

	defer stmt.Close()

	//execute
	res, err := stmt.Exec(transaction.Valued,
		transaction.Status,
		transaction.Info,
		transaction.From,
		transaction.To,
		transaction.AccountID,
		transaction.AmountID)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Failed to create transaction", err.Error()))
	}
	id, err := res.LastInsertId()

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Failed to create transaction", err.Error()))
	}

	return int(id), nil
}

func (m *mysqlRepo) GetTransactions(accountId int) (*models.TransactionList, error) {

	var transactionList = make([]*models.Transaction, 0)
	rows, err := m.db.Query(getTransactions, accountId)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Failed to fetch transactions", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		n := models.Transaction{}
		if err = rows.Scan(&n.ID,
			&n.Valued,
			&n.Status,
			&n.Info,
			&n.From,
			&n.To,
			&n.AccountID,
			&n.AmountID); err != nil {
			return nil, errors.New(fmt.Sprintln("Failed to fetch transactions", err.Error()))
		}
		transactionList = append(transactionList, &n)
	}

	return &models.TransactionList{
		TotalCount:   0,
		TotalPages:   0,
		Page:         0,
		Size:         1,
		Transactions: transactionList,
	}, nil
}

func (m *mysqlRepo) GetTransaction(transactionId int) (*models.Transaction, error) {

	var transaction models.Transaction

	smtp, err := m.db.Prepare(getTransaction)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Failed to fetch transaction", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(transactionId).Scan(&transaction.ID,
		&transaction.Valued,
		&transaction.Status,
		&transaction.Info,
		&transaction.From,
		&transaction.To,
		&transaction.AccountID,
		&transaction.AmountID,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Failed to fetch transaction", err.Error()))
	}

	return &transaction, nil
}

func (m *mysqlRepo) MakeTransaction(transaction *models.Transaction, targetAccount *models.Account, sourceAccount *models.Account) error {
	ctx := context.Background()

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}
	_, err = tx.ExecContext(ctx, updateAmount, sourceAccount.Balance.Amount.Value, sourceAccount.Balance.Amount.ID)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}
	_, err = tx.ExecContext(ctx, updateAmount, targetAccount.Balance.Amount.Value, targetAccount.Balance.Amount.ID)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}

	//execute
	res, err := tx.ExecContext(ctx, createAmount,transaction.Amount.Currency, transaction.Amount.Value)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}

	id, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}

	_, err = tx.ExecContext(ctx,
		createTransaction,
		"to",
		transaction.Info,
		transaction.From,
		transaction.To,
		targetAccount.ID,
		id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}
	_, err = tx.ExecContext(ctx,
		createTransaction,
		"from",
		transaction.Info,
		transaction.From,
		transaction.To,
		sourceAccount.ID,
		id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}

	err = tx.Commit()
	if err != nil {
		return errors.New(fmt.Sprintln("transaction failed", err.Error()))
	}

	return nil
}
