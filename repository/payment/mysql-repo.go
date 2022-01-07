package payment

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
func NewMysqlRepository(db *sql.DB) repository.PaymentRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) AddPayment(payment *models.Payment) (int, error) {
	stmt, err := m.db.Prepare(createPayment)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}

	defer stmt.Close()

	//execute
	res, err := stmt.Exec(payment.From,
		payment.To,
		payment.Booked,
		payment.Valued,
		payment.AccountID,
		payment.AmountID)

	if err != nil {
		return 0, errors.New(fmt.Sprintln("Payment could not be created", err.Error()))
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, errors.New(fmt.Sprintln("could not get lastInsertId", err.Error()))
	}

	return int(id), nil
}

func (m *mysqlRepo) GetPayments(accountId int) (*models.PaymentList, error) {

	var paymentList = make([]*models.Payment, 0)
	rows, err := m.db.Query(getPayments, accountId)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get Payments", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		n := models.Payment{}
		if err = rows.Scan(&n.ID,
			&n.From,
			&n.To,
			&n.Booked,
			&n.Valued,
			&n.AccountID,
			&n.AmountID); err != nil {
			return nil, errors.New(fmt.Sprintln("Could not get Payments", err.Error()))
		}
		paymentList = append(paymentList, &n)
	}

	return &models.PaymentList{
		TotalCount: 0,
		TotalPages: 0,
		Page:       0,
		Size:       1,
		Payments:   paymentList,
	}, nil
}

func (m *mysqlRepo) GetPayment(paymentId int) (*models.Payment, error) {

	var payment models.Payment

	smtp, err := m.db.Prepare(getPayment)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get Payment", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(paymentId).Scan(&payment.ID,
		&payment.From,
		&payment.To,
		&payment.Booked,
		&payment.Valued,
		&payment.AccountID,
		&payment.AmountID,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get Payment", err.Error()))
	}

	return &payment, nil
}
