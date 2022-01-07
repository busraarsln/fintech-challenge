package controller

import "net/http"

type BalanceController interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
	UpdateBalance(w http.ResponseWriter, r *http.Request)
}
