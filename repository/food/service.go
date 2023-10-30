package food


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

func (s *Service) GetFoodItemByID(foodID any)(*domain.FoodItem,error){
	return s.repo.GetFoodItemByID(foodID)
}

func (s *Service) GetMenuOfRestaurent(restaurentID any)(*[]domain.FoodItem, error){
	return s.repo.GetMenuOfRestaurent(restaurentID)
}

func (s *Service) DeleteFoodItem(foodID any)(error){
	return s.repo.DeleteFoodItem(foodID)
}

func (s *Service) AddFoodToCart(foodID any, userID any)(error){
	return s.repo.AddFoodToCart(foodID, userID)
}

func (s *Service) RemoveFoodFromCart(foodID any, userID any)(error){
	return s.repo.RemoveFoodFromCart(foodID, userID)
}

func (s *Service) CreateFoodItem(foodItem domain.FoodItem)(*domain.FoodItem,error){
	return s.repo.CreateFoodItem(foodItem)
}

func (s *Service) UpdateFoodItem(foodID any, updatedFoodItem domain.FoodItem)(*domain.FoodItem, error){
	return s.repo.UpdateFoodItem(foodID, updatedFoodItem)
}