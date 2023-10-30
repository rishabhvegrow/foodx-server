package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"

	"foodx-server/handler/packages/middlewares"

	"foodx-server/database"
	"foodx-server/domain"
	"foodx-server/repository/user"
	"foodx-server/controller/usercontroller"
	"foodx-server/handler/userhandler"

	"foodx-server/repository/restaurent"
	"foodx-server/controller/restaurentcontroller"
	"foodx-server/handler/restaurenthandler"

	"foodx-server/repository/food"
	"foodx-server/controller/foodcontroller"
	"foodx-server/handler/foodhandler"

	"foodx-server/repository/cart"
	"foodx-server/controller/cartcontroller"
	"foodx-server/handler/carthandler"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	// Migrating DB models
	db := database.GetDB()
	db.AutoMigrate(&domain.User{}, &domain.Restaurant{}, &domain.FoodItem{}, &domain.CartItem{}, &domain.Transaction{})

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} 
    config.AllowMethods = []string{"*"}
    config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	// User Routes
	userRepository := user.NewUserRepository(db)
	userService := user.NewService(userRepository)
	userController := usercontroller.NewUserController(userService)
	userHandler := userhandler.NewUserHandler(userController)

	usersGroup := router.Group("/users")
	usersGroup.GET("/", userHandler.GetUsers)

	userGroup := router.Group("/user")
	userGroup.GET("/:id", userHandler.GetUser)
    userGroup.DELETE("/:id", userHandler.DeleteUser)

	// Base Routes
	base := router.Group("/")
    base.GET("/", userHandler.GetPing)
    base.POST("signup/", userHandler.CreateUser)
    base.POST("signin/", userHandler.UserLogin)


	// Restaurent Routes
	restRepository := restaurent.NewRestaurantRepository(db)
	restService := restaurent.NewService(restRepository)
	restController := restaurentcontroller.NewRestController(restService)
	restHandler := restaurenthandler.NewRestaurentHandler(restController)

	restGroup := router.Group("/restaurents")
    restGroup.GET("/", restHandler.GetRestaurents)
    restGroup.GET("/:id", restHandler.GetRestaurent)
    restGroup.POST("/", restHandler.CreateRestaurent)
    restGroup.DELETE("/:id", restHandler.DeleteRestaurent)

	// Food Routes
	foodRepo := food.NewFoodRepository(db)
	foodService := food.NewService(foodRepo)
	foodController := foodcontroller.NewFoodController(foodService)
	foodHanler := foodhandler.NewFoodHandler(foodController)

    foodGroup := router.Group("/food")
    foodGroup.Use(middlewares.JWTAuthMiddleware())
    foodGroup.GET("/:restid", foodHanler.GetMenuOfRestaurent)
    foodGroup.POST("/", foodHanler.CreateFoodItem)
    foodGroup.PUT("/:id", foodHanler.UpdateFoodItem)
    foodGroup.DELETE("/:id", foodHanler.DeleteFoodItem)
    foodGroup.POST("/add/:id", foodHanler.AddFoodToCart)
    foodGroup.POST("/remove/:id", foodHanler.RemoveFoodFromCart)


	// Cart Routes
	cartRepo := cart.NewCartRepository(db)
	cartService := cart.NewService(cartRepo)
	cartController := cartcontroller.NewCartController(cartService)
	cartHandler := carthandler.NewCartHandler(cartController)

    cartGroup := router.Group("/cart")
    cartGroup.Use(middlewares.JWTAuthMiddleware())
    cartGroup.GET("/", cartHandler.GetCartItemsOfAUser)
    cartGroup.DELETE("/remove/:id", cartHandler.RemoveCartItemByID)
    cartGroup.GET("/orders", cartHandler.GetOrderedItemsOfAUser)
    cartGroup.POST("/checkout", cartHandler.CheckoutCart)

    // transactionGroup := router.Group("/transactions")
    // transactionGroup.Use(middlewares.JWTAuthMiddleware())
    // transactionGroup.GET("/", controlers.GetTransactions)

	router.Run(":8080")
}
