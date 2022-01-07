package controller

import "net/http"

type CustomerController interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	AddCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}
