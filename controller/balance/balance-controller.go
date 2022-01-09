package balance

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
	balanceService service.BalanceService
)

func NewBalanceController(service service.BalanceService) c.BalanceController {
	balanceService = service
	return &controller{}
}

func (*controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, _ := strconv.Atoi(vars["accountId"])
	w.Header().Set("Content-Type", "application/json")
	balance, err := balanceService.GetBalance(accountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the transactions"})
		return
	} 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}

func (*controller) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, _ := strconv.Atoi(vars["accountId"])
	w.Header().Set("Content-Type", "application/json")
	var balance models.Balance
	err := json.NewDecoder(r.Body).Decode(&balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	err1 := balanceService.Validate(&balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	err2 := balanceService.UpdateBalance(accountId, &balance)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}
