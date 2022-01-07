package main

import (
	"fmt"
	"net/http"

	"github.com/busraarsln/fintech-challenge/controller"
	_accountController "github.com/busraarsln/fintech-challenge/controller/account"
	_balanceController "github.com/busraarsln/fintech-challenge/controller/balance"
	_customerController "github.com/busraarsln/fintech-challenge/controller/customer"
	_paymentController "github.com/busraarsln/fintech-challenge/controller/payment"
	_transactionController "github.com/busraarsln/fintech-challenge/controller/transaction"
	"github.com/busraarsln/fintech-challenge/driver"
	router "github.com/busraarsln/fintech-challenge/http"
	"github.com/busraarsln/fintech-challenge/repository"
	_accountRepo "github.com/busraarsln/fintech-challenge/repository/account"
	_amountRepo "github.com/busraarsln/fintech-challenge/repository/amount"
	_balanceRepo "github.com/busraarsln/fintech-challenge/repository/balance"
	_customerRepo "github.com/busraarsln/fintech-challenge/repository/customer"
	_paymentRepo "github.com/busraarsln/fintech-challenge/repository/payment"
	_transactionRepo "github.com/busraarsln/fintech-challenge/repository/transaction"
	"github.com/busraarsln/fintech-challenge/service"
	_accountService "github.com/busraarsln/fintech-challenge/service/account"
	_balanceService "github.com/busraarsln/fintech-challenge/service/balance"
	_customerService "github.com/busraarsln/fintech-challenge/service/customer"
	_paymentService "github.com/busraarsln/fintech-challenge/service/payment"
	_transactionService "github.com/busraarsln/fintech-challenge/service/transaction"
	"github.com/go-openapi/runtime/middleware"
)

var (
	db, _                                                  = driver.CreateDbConnection()
	accountRepository     repository.AccountRepository     = _accountRepo.NewMysqlRepository(db)
	balanceRepository     repository.BalanceRepository     = _balanceRepo.NewMysqlRepository(db)
	transactionRepository repository.TransactionRepository = _transactionRepo.NewMysqlRepository(db)
	customerRepository    repository.CustomerRepository    = _customerRepo.NewMysqlRepository(db)
	amountRepository      repository.AmountRepository      = _amountRepo.NewMysqlRepository(db)
	paymentRepository     repository.PaymentRepository     = _paymentRepo.NewMysqlRepository(db)
	accountService        service.AccountService           = _accountService.NewAccountService(accountRepository, balanceRepository, amountRepository)
	accountController     controller.AccountController     = _accountController.NewAccountController(accountService)
	balanceService        service.BalanceService           = _balanceService.NewBalanceService(balanceRepository)
	balanceController     controller.BalanceController     = _balanceController.NewBalanceController(balanceService)
	transactionService    service.TransactionService       = _transactionService.NewTransactionService(transactionRepository, accountRepository)
	transactionController controller.TransactionController = _transactionController.NewTransactionController(transactionService)
	customerService       service.CustomerService          = _customerService.NewCustomerService(customerRepository)
	customerController    controller.CustomerController    = _customerController.NewCustomerController(customerService)
	paymentService        service.PaymentService           = _paymentService.NewPaymentService(paymentRepository)
	paymentController     controller.PaymentController     = _paymentController.NewPaymentController(paymentService)
	httpRouter            router.Router                    = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"

	httpRouter.Handle("/swaggerv013.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swaggerv013.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	httpRouter.Handle("/docs", sh)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	
	httpRouter.GET("/customers", customerController.GetCustomers)
	httpRouter.POST("/customers", customerController.AddCustomer)
	httpRouter.DELETE("/customers/{id:[0-9]+}", customerController.DeleteCustomer)

	httpRouter.GET("/customers/{id:[0-9]+}/accounts", accountController.GetAccounts)
	httpRouter.POST("/customers/{id:[0-9]+}/accounts", accountController.AddAccount)
	httpRouter.DELETE("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}", accountController.DeleteAccount)

	httpRouter.GET("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/balance", balanceController.GetBalance)
	httpRouter.PUT("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/balance", balanceController.UpdateBalance)

	httpRouter.GET("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/transactions", transactionController.GetTransacitons)
	httpRouter.POST("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/transactions", transactionController.MakeTransaction)

	httpRouter.GET("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/payments", paymentController.GetPayments)
	httpRouter.POST("/customers/{id:[0-9]+}/accounts/{accountId:[0-9]+}/payments", paymentController.AddPayment)

	httpRouter.SERVE(port)
}
