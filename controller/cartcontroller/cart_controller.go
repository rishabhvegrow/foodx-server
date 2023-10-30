package cartcontroller

import (
	"foodx-server/domain"
	"foodx-server/repository/cart"
)

type CartController struct {
	cartRepo *cart.Service
}

func NewCartController(cartRepo *cart.Service) *CartController {
	return &CartController{cartRepo: cartRepo}
}

func (cc *CartController) GetCartItemByID(itemID any)(*domain.CartItem, error) {
	cartitem, err := cc.cartRepo.GetCartItemByID(itemID)

	if err != nil {
		return nil, err
	}

	return cartitem, nil
}


func (cc *CartController) GetCartItemsOfAUser(userID any, isChecked bool)(*[]domain.CartItem, error){
	cartitems, err := cc.cartRepo.GetCartItemsOfAUser(userID,isChecked)

	if err != nil{
		return nil, err
	}

	return cartitems, nil
}

func (cc *CartController) CreateCartItem(cartItem domain.CartItem)(*domain.CartItem, error){
	cartitem, err := cc.cartRepo.CreateCartItem(cartItem)

	if err != nil{
		return nil, err
	}

	return cartitem, nil
}


func (cc *CartController) RemoveCartItemByID(itemID any)(error){
	err := cc.cartRepo.RemoveCartItemByID(itemID)
	
	if err != nil{
		return err
	}

	return nil
}


func (cc *CartController) CheckoutCart(userID any)(*domain.Transaction, error){
	transaction, err := cc.cartRepo.CheckoutCart(userID)
	
	if err != nil{
		return nil,err
	}

	return transaction,nil
}