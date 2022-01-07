package amount

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/busraarsln/fintech-challenge/models"
	"github.com/busraarsln/fintech-challenge/repository"
	_ "github.com/go-sql-driver/mysql"
)

type (
	mysqlRepo struct{ db *sql.DB }
)

//NewMysqlRepository creates a new repo
func NewMysqlRepository(db *sql.DB) repository.AmountRepository {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) AddAmount(amount *models.Amount) (int, error) {

	stmt, e := m.db.Prepare(createAmount)
	if e != nil {
		return 0, errors.New(fmt.Sprintln("Sql syntax error", e.Error()))
	}

	//execute
	res, e := stmt.Exec(amount.Currency, amount.Value)

	if e != nil {
		return 0, errors.New(fmt.Sprintln("Amount could not be created", e.Error()))
	}

	id, e := res.LastInsertId()

	if e != nil {
		return 0, errors.New(fmt.Sprintln("Cannot get lastInsertId", e.Error()))
	}

	return int(id), nil
}
