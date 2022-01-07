package repository

import (
	"github.com/busraarsln/fintech-challenge/models"
)

type AmountRepository interface {
	AddAmount(amount *models.Amount) (int, error)
}
