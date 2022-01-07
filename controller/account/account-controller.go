package account

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
	accountService service.AccountService
)

func NewAccountController(service service.AccountService) c.AccountController {
	accountService = service
	return &controller{}
}

func (*controller) GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	customerId, _ := strconv.Atoi(vars["id"])
	accounts, err := accountService.GetAccounts(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func (*controller) AddAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	customerId, _ := strconv.Atoi(vars["id"])
	var account models.Account
	account.CustomerID = customerId
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := accountService.Validate(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := accountService.AddAccount(&account)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*controller) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	accountId, _ := strconv.Atoi(vars["accountId"])
	err := accountService.DeleteAccount(accountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error deleting the account"})
	}
	w.WriteHeader(http.StatusOK)
}
