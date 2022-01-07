package models

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Surname     string       `json:"surname"`
	Password    string       `json:"password" mask:"password"`
	PhoneNumber string       `json:"phone_number"`
	Email       string       `json:"email"`
	Role        string       `json:"role"`
	Accounts    []Account    `json:"-"`
	Address     Address      `json:"-"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt   sql.NullTime `json:"updated_at,omitempty"`
	IsActive    bool         `json:"is_active,omitempty"`
}

type CustomerList struct {
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	Customers  []*Customer `json:"customers"`
}
