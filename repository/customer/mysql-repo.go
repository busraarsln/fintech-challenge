package customer

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	masker "github.com/ggwhite/go-masker"
	_ "github.com/go-sql-driver/mysql"
)

type (
	mysqlRepo struct{ db *sql.DB }
)

//NewMysqlRepository creates a new repo
func NewMysqlRepository(db *sql.DB) repository.CustomerRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) AddCustomer(customer *models.Customer) (int, error) {
	stmt, e := m.db.Prepare(createCustomer)
	defer stmt.Close()

	//execute
	res, e := stmt.Exec(customer.Name,
		customer.Surname,
		customer.PhoneNumber,
		customer.Email,
		customer.Role,
		customer.Password)

	id, e := res.LastInsertId()

	if e != nil {
		panic("create balance error")
	}

	return int(id), nil
}

func (m *mysqlRepo) GetCustomers() (*models.CustomerList, error) {

	var customerList = make([]*models.Customer, 0)
	rows, err := m.db.Query(getCustomers)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Cannot get customers", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		n := models.Customer{}
		if err = rows.Scan(&n.ID,
			&n.Name,
			&n.Surname,
			&n.PhoneNumber,
			&n.Email,
			&n.Role,
			&n.CreatedAt,
			&n.UpdatedAt,
			&n.IsActive,
			&n.Password); err != nil {
			return nil, errors.New(fmt.Sprintln("Could not get customers", err.Error()))
		}
		password := masker.Password(n.Password)
		n.Password = password
		email := masker.Email(n.Email)
		n.Email = email
		phone := masker.Telephone(n.PhoneNumber)
		n.PhoneNumber = phone
		customerList = append(customerList, &n)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get customers", err.Error()))
	}
	return &models.CustomerList{
		TotalCount: 0,
		TotalPages: 0,
		Page:       0,
		Size:       1,
		Customers:  customerList,
	}, nil
}

func (m *mysqlRepo) DeleteCustomer(id int) error {

	stmt, err := m.db.Prepare(deleteCustomer)
	if err != nil {
		return errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return errors.New(fmt.Sprintln("Customer could not be deleted", err.Error()))
	}
	raff, err := res.RowsAffected()

	if err != nil || raff < 1 {
		return errors.New(fmt.Sprintln("Customer could not be deleted", err.Error()))
	}

	return nil
}

func (m *mysqlRepo) GetCustomer(customerId int) (*models.Customer, error) {

	var customer models.Customer

	smtp, err := m.db.Prepare(getCustomer)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Sql syntax error", err.Error()))
	}
	defer smtp.Close()

	err = smtp.QueryRow(customerId).Scan(&customer.ID,
		&customer.Name,
		&customer.Surname,
		&customer.PhoneNumber,
		&customer.Email,
		&customer.Role,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.IsActive,
		&customer.Password,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintln("Could not get customer", err.Error()))
	}

	return &customer, err
}
