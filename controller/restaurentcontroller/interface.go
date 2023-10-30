package restaurentcontroller


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

type Controller interface {
	Reader
	Writer
}