package models

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int          `json:"id"`
	Type        string       `json:"type"`
	Status      string       `json:"status, omitempty"`
	AccountNo   string       `json:"account_no"`
	Iban        string       `json:"iban"`
	Currency    string       `json:"currency"`
	Description string       `json:"description"`
	Nickname    string       `json:"nickname"`
	CustomerID  int          `json:"customer_id"`
	BalanceID   int          `json:"balance_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	IsActive    bool         `json:"is_active"`

	Balance      Balance       `json:"balance"`
	Transactions []Transaction `json:"-"`
	Payments     []Payment     `json:"-"`
}

type AccountList struct {
	TotalCount int        `json:"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	Accounts   []*Account `json:"accounts"`
}
