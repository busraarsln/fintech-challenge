package models

import "time"

type Transaction struct {
	ID        int       `json:"id,omitempty"`
	Valued    time.Time `json:"valued,omitempty"`
	Status    string    `json:"status"`
	Info      string    `json:"info"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	AccountID int       `json:"account_id"`
	AmountID  int       `json:"amount_id"`

	Amount  Amount  `json:"amount"`
	Account Account `json:"-"`
}

type TransactionList struct {
	TotalCount   int            `json:"total_count"`
	TotalPages   int            `json:"total_pages"`
	Page         int            `json:"page"`
	Size         int            `json:"size"`
	Transactions []*Transaction `json:"transactions"`
}
