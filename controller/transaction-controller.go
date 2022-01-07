package controller

import "net/http"

type TransactionController interface {
	GetTransacitons(w http.ResponseWriter, r *http.Request)
	AddTransaction(w http.ResponseWriter, r *http.Request)
	MakeTransaction(w http.ResponseWriter, r *http.Request)
}
