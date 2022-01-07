package payment

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
	paymentService service.PaymentService
)

func NewPaymentController(service service.PaymentService) c.PaymentController {
	paymentService = service
	return &controller{}
}

func (*controller) GetPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	accountId, _ := strconv.Atoi(vars["accountId"])
	payments, err := paymentService.GetPayments(accountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the payments"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func (*controller) AddPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := paymentService.Validate(&payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := paymentService.AddPayment(&payment)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the payment"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
