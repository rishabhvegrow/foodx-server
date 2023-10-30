package restaurent

import (
	"foodx-server/domain"
)

// Reader Interfce
type Reader interface {
	GetRestaurents()(*[]domain.Restaurant, error)
	GetRestaurent(restID any)(*domain.Restaurant, error)
}


type Writer interface {
	CraeteRestaurent(restaurent domain.Restaurant)(*domain.Restaurant, error)
	DeleteRestaurent(restID any)(error)
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