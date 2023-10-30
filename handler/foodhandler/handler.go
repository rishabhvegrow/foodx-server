package foodhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"foodx-server/controller/foodcontroller"
	"foodx-server/domain"
)


type FoodHandler struct {
	foodController foodcontroller.Controller
}

func NewFoodHandler(foodController foodcontroller.Controller) *FoodHandler {
	return &FoodHandler{foodController: foodController}
}

func (fh *FoodHandler) CreateFoodItem(c *gin.Context){
    var foodItem domain.FoodItem
    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
	createdFoodItem, err := fh.foodController.CreateFoodItem(foodItem)

    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, createdFoodItem)
}

func (fh *FoodHandler) UpdateFoodItem(c *gin.Context){
    
    fooditemID := c.Param("id")
    var foodItem domain.FoodItem

    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedFoodItem, err := fh.foodController.UpdateFoodItem(fooditemID, foodItem)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item"})
        return
    }
    
    c.JSON(http.StatusOK, updatedFoodItem)
}


func (fh *FoodHandler) GetMenuOfRestaurent(c *gin.Context){
    restID := c.Param("restid")
    foodItems, err := fh.foodController.GetMenuOfRestaurent(restID)

    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find Records"})
    }
    c.JSON(http.StatusOK, foodItems)
}


func (fh *FoodHandler) AddFoodToCart(c *gin.Context){
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

	err := fh.foodController.AddFoodToCart(foodID, userID)

    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func (fh *FoodHandler) RemoveFoodFromCart(c *gin.Context){
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

	err := fh.foodController.RemoveFoodFromCart(foodID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}


func (fh *FoodHandler) DeleteFoodItem(c *gin.Context) {
    fooditemID := c.Param("id")
    err := fh.foodController.DeleteFoodItem(fooditemID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "foodItem not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "foodItem deleted"})
}