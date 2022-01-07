package models

type Amount struct {
	ID       int    `json:"id"`
	Currency string `json:"currency"`
	// to prevent loosing data
	Value float64 `json:"value"`
}
