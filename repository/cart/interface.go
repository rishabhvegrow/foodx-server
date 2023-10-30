package cart


import (
	"foodx-server/domain"
)

// Reader Interfce
type Reader interface {
	GetCartItemByID(itemID any)(*domain.CartItem, error)
	GetCartItemsOfAUser(userID any, isChecked bool)(*[]domain.CartItem, error)
}

type Writer interface {
	CreateCartItem(cartItem domain.CartItem)(*domain.CartItem, error)
	RemoveCartItemByID(id any)(error)
	CheckoutCart(userID any)(*domain.Transaction, error)
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