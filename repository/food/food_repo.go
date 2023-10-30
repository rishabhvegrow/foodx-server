package food

import (
	"gorm.io/gorm"
	"foodx-server/domain"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (food_repo *FoodRepository) CreateFoodItem(foodItem domain.FoodItem)(*domain.FoodItem,error){
	db := food_repo.db
    if err := db.Create(&foodItem).Error; err != nil{
        return nil,err
    }
	return &foodItem, nil
}

func (food_repo *FoodRepository) GetFoodItemByID(foodID any)(*domain.FoodItem,error){
    db := food_repo.db
	var foodItem domain.FoodItem

    if err := db.First(&foodItem, foodID).Error; err != nil {
        return nil, err
    }
	return &foodItem, nil
}

func (food_repo *FoodRepository) UpdateFoodItem(foodID any,updatedFoodItem domain.FoodItem)(*domain.FoodItem, error){
    db := food_repo.db
	foodItem, err := food_repo.GetFoodItemByID(foodID)
    if err != nil {
        return nil,err
    }

	foodItem.Name = updatedFoodItem.Name
    foodItem.Price = updatedFoodItem.Price
    foodItem.Description = updatedFoodItem.Description
    foodItem.RestaurantID = updatedFoodItem.RestaurantID
    foodItem.ImageUrl = updatedFoodItem.ImageUrl

    if err:=db.Save(&foodItem).Error; err != nil {
        return nil, err
    }

	return foodItem, nil
}

func (food_repo *FoodRepository) DeleteFoodItem(foodID any)(error){
    db := food_repo.db
    var foodItem domain.FoodItem
    if err := db.Delete(&foodItem, foodID).Error; err != nil {
        return err
    }
    return nil
}

func (food_repo *FoodRepository) GetMenuOfRestaurent(restaurentID any)(*[]domain.FoodItem, error){
    db := food_repo.db
    var foodItems []domain.FoodItem

    if err := db.Where("restaurant_id = ?", restaurentID).Find(&foodItems).Error; err != nil{
        return nil, err
    }

    return &foodItems, nil
}

func (food_repo *FoodRepository) AddFoodToCart(foodID any, userID any)(error){
    db := food_repo.db
    var foodItem domain.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        return err
    }

    var cartItem domain.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        newItem := domain.CartItem{
            UserID:     userID.(uint),
            FoodItemID: foodItem.ID,
            Quantity:   1,
            Price:     foodItem.Price,
        }
        if err := db.Create(&newItem).Error; err != nil {
            return err
        }
    } else {
        cartItem.Quantity++
        cartItem.Price += foodItem.Price
        if err := db.Save(&cartItem).Error; err != nil {
            return err
        }
    }
    return nil
}

func (food_repo *FoodRepository) RemoveFoodFromCart(foodID any, userID any)(error){
    db := food_repo.db
    var foodItem domain.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        return err
    }

    var cartItem domain.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        return err
    }

    cartItem.Quantity--
    cartItem.Price -= foodItem.Price
    if cartItem.Quantity == 0 {
        if err := db.Delete(&cartItem).Error; err != nil {
            return err
        }
    } else {
        if err := db.Save(&cartItem).Error; err != nil {
            return err
        }
    }

    return nil
}