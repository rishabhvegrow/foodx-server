package transaction

import (
	"foodx-server/domain"
)

// Reader Interfce
type Reader interface {
	GetTransactionsOfUser(userID any)(*[]domain.Transaction, error)
}

type Writer interface {
	CreateTransaction(userID any, total float64) (*domain.Transaction, error)
	UpdateTransaction(transactionID any, total float64) (*domain.Transaction, error)
}

// Filter interface
type Filter interface {
}

// Repository interface
type Repository interface {
	Reader
	Writer
	Filter
}