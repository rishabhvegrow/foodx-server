package cart

import (
	"foodx-server/domain"
)

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) GetCartItemByID(itemID any)(*domain.CartItem, error){
	return s.repo.GetCartItemByID(itemID)
}

func (s *Service) GetCartItemsOfAUser(userID any, isChecked bool)(*[]domain.CartItem, error){
	return s.repo.GetCartItemsOfAUser(userID, isChecked)
}

func (s *Service) CreateCartItem(cartItem domain.CartItem)(*domain.CartItem, error){
	return s.repo.CreateCartItem(cartItem)
}

func (s *Service) RemoveCartItemByID(itemID any)(error){
	return s.repo.RemoveCartItemByID(itemID)
}

func (s *Service) CheckoutCart(userID any)(*domain.Transaction, error){
	return s.repo.CheckoutCart(userID)
}