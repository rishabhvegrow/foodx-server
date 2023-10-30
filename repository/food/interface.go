package food

import (
	"foodx-server/domain"
)

// Reader Interfce
type Reader interface {
	GetFoodItemByID(foodID any)(*domain.FoodItem,error)
	GetMenuOfRestaurent(restaurentID any)(*[]domain.FoodItem, error)
}

type Writer interface {
	DeleteFoodItem(foodID any)(error)
	AddFoodToCart(foodID any, userID any)(error)
	RemoveFoodFromCart(foodID any, userID any)(error)
	CreateFoodItem(foodItem domain.FoodItem)(*domain.FoodItem,error)
	UpdateFoodItem(foodID any,updatedFoodItem domain.FoodItem)(*domain.FoodItem, error)
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