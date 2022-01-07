package account

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/busraarsln/fintech-challenge/driver"
// 	"github.com/busraarsln/fintech-challenge/models"
// 	"github.com/busraarsln/fintech-challenge/repository"
// 	_accountRepo "github.com/busraarsln/fintech-challenge/repository/account"
// 	_amountRepo "github.com/busraarsln/fintech-challenge/repository/amount"
// 	_balanceRepo "github.com/busraarsln/fintech-challenge/repository/balance"
// 	"github.com/busraarsln/fintech-challenge/service"
// 	_accountService "github.com/busraarsln/fintech-challenge/service/account"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	db, _                                          = driver.CreateDbConnection()
// 	accountRepo       repository.AccountRepository = _accountRepo.NewMysqlRepository(db)
// 	balanceRepo       repository.AccountRepository = _balanceRepo.NewMysqlRepository(db)
// 	amountRepo        repository.AccountRepository = _amountRepo.NewMysqlRepository(db)
// 	accountSrv        service.AccountService       = _accountService.NewAccountService(accountRepo, balanceRepo, amountRepo)
// 	accountController AccountController            = NewAccountController(accountSrv)
// )

// func TestAddAccount(t *testing.T) {
// 	//Create a new HTTP POST request
// 	var json = []byte(`{"account_no":"1212","iban":"2112"}`)
// 	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(json))
// 	//Assign HTTP Handler function (controller AddAccount function)
// 	handler := http.HandlerFunc(accountController.AddAccount)
// 	//Record HTTP Response (httptest)
// 	response := httptest.NewRecorder()
// 	//Dispatch the HTTP request
// 	handler.ServeHTTP(response, req)
// 	//Add Assertions on the HTTP Status code and the response
// 	status := response.Code
// 	if status != http.StatusOK {
// 		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	//Docode the HTTP response
// 	var account models.Account
// 	//json.NewDecoder(io.Reader(response.Body)).Decode(&account)

// 	//Assert HTTP response
// 	assert.NotNil(t, account.ID)

// 	// //Clean up db
// 	// CleanUp(&account)

// }

// func TestGetAccounts(t *testing.T) {
// 	//Create a new HTTP POST request
// 	req, _ := http.NewRequest("GET", "/accounts", nil)
// 	//Assign HTTP Handler function (controller AddAccount function)
// 	handler := http.HandlerFunc(accountController.GetAccounts)
// 	//Record HTTP Response (httptest)
// 	response := httptest.NewRecorder()
// 	//Dispatch the HTTP request
// 	handler.ServeHTTP(response, req)
// 	//Add Assertions on the HTTP Status code and the response
// 	status := response.Code
// 	if status != http.StatusOK {
// 		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	//Docode the HTTP response
// 	var account models.Account
// 	json.NewDecoder(io.Reader(response.Body)).Decode(&account)

// 	//Assert HTTP response
// 	assert.NotNil(t, account.ID)

// 	// //Clean up db
// 	// CleanUp(&account)

// }

// // func CleanUp(account *models.account) {
// // 	accountRepo.Delete(account)
// // }
