package foodcontroller

import (
	"foodx-server/domain"
	"foodx-server/repository/food"
)

type FoodController struct {
	foodRepo *food.Service
}

func NewFoodController(foodRepo *food.Service) *FoodController {
	return &FoodController{foodRepo: foodRepo}
}

func (fc *FoodController) GetFoodItemByID(foodID any)(*domain.FoodItem,error){
	fooditem, err := fc.foodRepo.GetFoodItemByID(foodID)

	if err != nil{
		return nil, err
	}

	return fooditem, nil
}

func (fc *FoodController) GetMenuOfRestaurent(restaurentID any)(*[]domain.FoodItem, error){
	fooditems, err := fc.foodRepo.GetMenuOfRestaurent(restaurentID)

	if err != nil{
		return nil, err
	}

	return fooditems, nil
}

func (fc *FoodController) CreateFoodItem(foodItem domain.FoodItem)(*domain.FoodItem,error) {
	createdFooditem, err := fc.foodRepo.CreateFoodItem(foodItem)

	if err != nil{
		return nil, err
	}

	return createdFooditem, nil
}

func (fc *FoodController) UpdateFoodItem(foodID any,updatedFoodItem domain.FoodItem)(*domain.FoodItem, error){
	fooditem, err := fc.foodRepo.UpdateFoodItem(foodID, updatedFoodItem)

	if err != nil{
		return nil, err
	}

	return fooditem, nil
}

func (fc *FoodController) AddFoodToCart(foodID any, userID any)(error){
	err := fc.foodRepo.AddFoodToCart(foodID, userID)
	if err != nil{
		return err
	}

	return nil
}

func (fc *FoodController) RemoveFoodFromCart(foodID any, userID any)(error){
	err := fc.foodRepo.RemoveFoodFromCart(foodID, userID)
	if err != nil{
		return err
	}

	return nil
}

func (fc *FoodController) DeleteFoodItem(foodID any)(error){
	err := fc.foodRepo.DeleteFoodItem(foodID)
	if err != nil{
		return err
	}

	return nil
}