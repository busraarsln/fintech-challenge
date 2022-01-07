# fintech-challenge with go 

A RESTful API example for a bank app. 

#### 👨‍💻 Full list what has been used:
* [rate](https://golang.org/x/time/rate) - Rate Limit
* [mux](https://github.com/gorilla/mux) - Mux
* [viper](https://github.com/spf13/viper) - Go configuration
* [uuid](https://github.com/google/uuid) - UUID
* [Docker](https://www.docker.com/) - Docker
* [Swagger](https://github.com/go-openapi/runtime/middleware) - Swagger
* [Mysql](https://github.com/go-sql-driver/mysql) - Mysql Driver



## Installation & Run
```bash
# Download this project
go get github.com/busraarsln/fintech-challenge
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config-local.yaml] (if you dont want to use docker) (https://github.com/busraarsln/fintech-go-challenge/blob/main/config/config-local.yml)
```
mysql:
  MysqlHost: 0.0.0.0
  MysqlPort: 3306
  MysqlUser: root
  MysqlPassword: admin
  MysqlDbname: fintech
  MysqlSslmode: false
  MysqlDriver: mysql

```

Before running app, you should run the script [script.sql](https://github.com/busraarsln/fintech-go-challenge/blob/main/script.sql)
```

```bash
#DB
docker run --name my-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=fintech -d mysql

# Build and Run
cd fintech-challenge
go run
server.go

# API Endpoint : http://localhost:8000
```

## API

#### /customers
* `GET` : Get all customers
* `POST` : Create a new customer
* `DELETE` : Delete a customer

#### //customers/:id/accounts
* `GET` : Get accounts of a customer
* `POST` : Create a new account
* `DELETE` : Delete a account

#### /customers/:id/accounts/:accountId/balance
* `GET` : Get a balance
* `PUT` : Update a balance

#### /customers/:id/accounts/:accountId/balance/transactions
* `GET` : Get transactions of a account
* `POST` : Make a transaction

#### /customers/:id/accounts/:accountId/payments
* `GET` :  Get payments of a account
* `POST` : Create a new payment


### SWAGGER UI:

https://localhost:8000/docs


## Todo

- [ ] Add RabbitMQ for payments.
