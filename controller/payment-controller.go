package controller

import "net/http"

type PaymentController interface {
	GetPayments(w http.ResponseWriter, r *http.Request)
	AddPayment(w http.ResponseWriter, r *http.Request)
}
