package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"foodx-server/controller/usercontroller"
	"foodx-server/domain"
)


type UserHandler struct {
	userController usercontroller.Controller
}

func NewUserHandler(userController usercontroller.Controller) *UserHandler {
	return &UserHandler{userController: userController}
}

func (uh *UserHandler) GetPing(c *gin.Context){
    c.JSON(http.StatusOK, "PONG")
}


func (uh *UserHandler) GetUsers(c *gin.Context){
	users, err := uh.userController.GetUsers()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}


func (uh *UserHandler) GetUser(c *gin.Context) {
    userID := c.Param("id")
    user, err := uh.userController.GetUser(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdUser, err := uh.userController.CraeteUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, createdUser)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
    userID := c.Param("id")
    
    err := uh.userController.DeleteUser(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}


func (uh *UserHandler) UserLogin(c *gin.Context) {
    var requestData map[string]interface{}
    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    email, emailExists := requestData["email"].(string)
    password, passwordExists := requestData["password"].(string)

    if !emailExists || !passwordExists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    response, err := uh.userController.UserLogin(email, password)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, response)
}