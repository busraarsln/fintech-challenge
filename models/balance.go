package models

import (
	"database/sql"
	"time"
)

type Balance struct {
	ID          int          `json:"id"`
	CreditLimit float64      `json:"credit_limit"`
	AmountId    int          `json:"amount_id,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at,omitempty"`

	Amount Amount `json:"amount"`
}
