package customer

import (
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_service "github.com/busraarsln/fintech-challenge/service"
)

type service struct{}

var (
	customerRepo repository.CustomerRepository
)

func NewCustomerService(
	customerRepository repository.CustomerRepository) _service.CustomerService {
	customerRepo = customerRepository

	return &service{}
}

func (*service) Validate(customer *models.Customer) error {
	if customer == nil {
		err := errors.New("the account is empty")
		return err
	}
	if customer.Email == "" {
		err := errors.New("the account is emptdy")
		return err
	}
	return nil
}
func (*service) AddCustomer(customer *models.Customer) (*models.Customer, error) {
	customerId, _ := customerRepo.AddCustomer(customer)
	return customerRepo.GetCustomer(customerId)

}
func (*service) GetCustomers() (*models.CustomerList, error) {
	return customerRepo.GetCustomers()
}

func (*service) DeleteCustomer(id int) error {
	_, err := customerRepo.GetCustomer(id)
	if err != nil {
		return errors.New(fmt.Sprintln("Customer not found", err.Error()))
	}

	return customerRepo.DeleteCustomer(id)
}
