package customer

import (
	"encoding/json"
	"net/http"
	"strconv"

	c "github.com/busraarsln/fintech-challenge/controller"
	"github.com/busraarsln/fintech-challenge/errors"
	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/service"
	"github.com/gorilla/mux"
)

type controller struct{}

var (
	customerService service.CustomerService
)

func NewCustomerController(service service.CustomerService) c.CustomerController {
	customerService = service
	return &controller{}
}

func (*controller) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers, err := customerService.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the transactions"})
		return
	} 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func (*controller) AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := customerService.Validate(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := customerService.AddCustomer(&customer)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := customerService.DeleteCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error deleting the customer"})
	}
	w.WriteHeader(http.StatusOK)
}
