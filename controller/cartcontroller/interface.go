package cartcontroller

import (
	"foodx-server/domain"
)

type Reader interface {
	GetCartItemByID(itemID any)(*domain.CartItem, error)
	GetCartItemsOfAUser(userID any, isChecked bool)(*[]domain.CartItem, error)
}

type Writer interface {
	CreateCartItem(cartItem domain.CartItem)(*domain.CartItem, error)
	RemoveCartItemByID(id any)(error)
	CheckoutCart(userID any)(*domain.Transaction, error)
}


type Controller interface {
	Reader
	Writer
}