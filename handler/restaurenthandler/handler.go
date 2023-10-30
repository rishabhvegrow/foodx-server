package restaurenthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"foodx-server/controller/restaurentcontroller"
	"foodx-server/domain"
)


type RestaurentHandler struct {
	restController restaurentcontroller.Controller
}

func NewRestaurentHandler(restController restaurentcontroller.Controller) *RestaurentHandler {
	return &RestaurentHandler{restController: restController}
}

func (rh *RestaurentHandler) GetRestaurents(c *gin.Context){
	restaurents, err := rh.restController.GetRestaurents()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Restaurents"})
        return
    }
    c.JSON(http.StatusOK, restaurents)
}


func (rh *RestaurentHandler) CreateRestaurent(c *gin.Context) {
    var restaurent domain.Restaurant
    if err := c.ShouldBindJSON(&restaurent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdRestaurent, err := rh.restController.CraeteRestaurent(restaurent)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, createdRestaurent)
}


func (rh *RestaurentHandler) GetRestaurent(c *gin.Context) {
    restID := c.Param("id")
    restaurent, err := rh.restController.GetRestaurent(restID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Restaurent not found"})
        return
    }
    c.JSON(http.StatusOK, restaurent)
}


func (rh *RestaurentHandler) DeleteRestaurent(c *gin.Context) {
    restID := c.Param("id")
    err := rh.restController.DeleteRestaurent(restID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Unable to delete"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Restaurent deleted"})
}