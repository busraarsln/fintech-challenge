package service

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type CustomerService interface {
	Validate(customer *models.Customer) error
	AddCustomer(customer *models.Customer) (*models.Customer, error)
	GetCustomers() (*models.CustomerList, error)
	DeleteCustomer(id int) error
}
