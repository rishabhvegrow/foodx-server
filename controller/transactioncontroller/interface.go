package transactioncontroller

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

type Controller interface {
	Reader
	Writer
}