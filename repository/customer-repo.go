package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type CustomerRepository interface {
	AddCustomer(costumer *models.Customer) (int, error)
	GetCustomers() (*models.CustomerList, error)
	GetCustomer(customerId int) (*models.Customer, error)
	DeleteCustomer(id int) error
}
