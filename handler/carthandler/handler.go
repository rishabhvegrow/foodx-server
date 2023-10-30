package carthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"foodx-server/controller/cartcontroller"
)


type CartHandler struct {
	cartController cartcontroller.Controller
}

func NewCartHandler(cartController cartcontroller.Controller) *CartHandler {
	return &CartHandler{cartController: cartController}
}

func (ch *CartHandler) GetCartItemByID(c *gin.Context){
	cartitemID := c.Param("id")
	cartItem, err := ch.cartController.GetCartItemByID(cartitemID)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart Item not found"})
        return
    }
    c.JSON(http.StatusOK, cartItem)
}


func (ch *CartHandler) GetCartItemsOfAUser(c *gin.Context){
	userID := c.MustGet("user_id")

	cartItems, err := ch.cartController.GetCartItemsOfAUser(userID, false)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart Items not found"})
        return
    }
    c.JSON(http.StatusOK, cartItems)
}

func (ch *CartHandler) GetOrderedItemsOfAUser(c *gin.Context){
	userID := c.MustGet("user_id")

	cartItems, err := ch.cartController.GetCartItemsOfAUser(userID, true)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ordered Items not found"})
        return
    }
    c.JSON(http.StatusOK, cartItems)
}


func (ch *CartHandler) RemoveCartItemByID(c *gin.Context){
	cartitemID := c.Param("id")
	err := ch.cartController.RemoveCartItemByID(cartitemID)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart Items not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Cart Items removed from cart"})
}

func (ch *CartHandler) CheckoutCart(c *gin.Context) {

    userID := c.MustGet("user_id")

    cartItems, err := ch.cartController.GetCartItemsOfAUser(userID, false)
    if err != nil || len(*cartItems) == 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added to the cart"})
        return
    }

    transaction, err := ch.cartController.CheckoutCart(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to checkout, Please try again"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Transaction successful", "transaction": transaction})
}