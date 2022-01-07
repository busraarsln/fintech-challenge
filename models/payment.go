package models

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID        string       `json:"id"`
	Amount    Amount       `json:"-"`
	From      string       `json:"from"`
	To        string       `json:"to"`
	Booked    time.Time    `json:"booked"`
	Valued    sql.NullTime `json:"valued,omitempty"`
	AccountID int          `json:"account_id"`
	AmountID  int          `json:"amount_id"`
}

type PaymentList struct {
	TotalCount int        `json:"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	Payments   []*Payment `json:"payments"`
}
