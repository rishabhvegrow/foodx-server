# API Endpoints Documentation

## Root Endpoints

### 1. `/` (GET)
- **Description**: This is the root endpoint for the application.
- **Action**: It is used to check the availability of the server.
- **Controller Function**: `controlers.GetPing`

### 2. `/signup/` (POST)
- **Description**: Used for user registration (sign up).
- **Action**: Creates a new user account.
- **Controller Function**: `controlers.CreateUser`

### 3. `/signin/` (POST)
- **Description**: Used for user authentication (sign in).
- **Action**: Authenticates a user and generates a JWT token for subsequent requests.
- **Controller Function**: `controlers.Login`

## User Endpoints

### 4. `/users` (GET)
- **Description**: Get a list of all users.
- **Action**: Retrieves a list of all registered users.
- **Controller Function**: `controlers.GetUsers`

### 5. `/users/:id` (GET)
- **Description**: Get user by ID.
- **Action**: Retrieves a specific user's details by their ID.
- **Controller Function**: `controlers.GetUser`

### 6. `/users/:id` (DELETE)
- **Description**: Delete user by ID.
- **Action**: Deletes a specific user by their ID.
- **Controller Function**: `controlers.DeleteUser`

## Restaurant Endpoints

### 7. `/restaurants` (GET)
- **Description**: Get a list of all restaurants.
- **Action**: Retrieves a list of all registered restaurants.
- **Controller Function**: `controlers.GetRestaurants`

### 8. `/restaurants/:id` (GET)
- **Description**: Get restaurant by ID.
- **Action**: Retrieves details of a specific restaurant by its ID.
- **Controller Function**: `controlers.GetRestaurant`

### 9. `/restaurants` (POST)
- **Description**: Create a new restaurant.
- **Action**: Creates a new restaurant entity.
- **Controller Function**: `controlers.CreateRestaurant`

### 10. `/restaurants/:id` (DELETE)
- **Description**: Delete restaurant by ID.
- **Action**: Deletes a specific restaurant by its ID.
- **Controller Function**: `controlers.DeleteRestaurant`

## Food Endpoints

### 11. `/food/:restid` (GET)
- **Description**: Get food items for a restaurant.
- **Action**: Retrieves a list of food items associated with a specific restaurant.
- **Controller Function**: `controlers.GetFoodItemOfRestaurant`

### 12. `/food` (POST)
- **Description**: Create a new food item.
- **Action**: Creates a new food item to be associated with a restaurant.
- **Controller Function**: `controlers.CreateFoodItem`

### 13. `/food/:id` (PUT)
- **Description**: Update a food item.
- **Action**: Updates the details of a specific food item.
- **Controller Function**: `controlers.UpdateFoodItem`

### 14. `/food/:id` (DELETE)
- **Description**: Delete food item by ID.
- **Action**: Deletes a specific food item by its ID.
- **Controller Function**: `controlers.DeleteFoodItem`

### 15. `/food/add/:id` (POST)
- **Description**: Add food item to cart.
- **Action**: Adds a food item to the user's shopping cart.
- **Controller Function**: `controlers.AddFoodToCart`

### 16. `/food/remove/:id` (POST)
- **Description**: Remove food item from cart.
- **Action**: Removes a food item from the user's shopping cart.
- **Controller Function**: `controlers.RemoveFoodFromCart`

## Cart Endpoints

### 17. `/cart` (GET)
- **Description**: Get details of the shopping cart.
- **Action**: Retrieves the details of the user's shopping cart.
- **Controller Function**: `controlers.GetCartDetails`

### 18. `/cart/remove/:id` (DELETE)
- **Description**: Remove item from cart.
- **Action**: Removes a specific item from the user's shopping cart.
- **Controller Function**: `controlers.RemoveCartItem`

### 19. `/cart/orders` (GET)
- **Description**: Get a list of ordered items.
- **Action**: Retrieves a list of items that have been ordered by the user.
- **Controller Function**: `controlers.GetOrderedItems`

### 20. `/cart/checkout` (POST)
- **Description**: Checkout the shopping cart.
- **Action**: Completes the order and checkout process for items in the user's shopping cart.
- **Controller Function**: `controlers.CheckoutCart`

## Transaction Endpoints

### 21. `/transactions` (GET)
- **Description**: Get a list of user transactions.
- **Action**: Retrieves a list of user transactions, such as order history.
- **Controller Function**: `controlers.GetTransactions`

These are the various endpoints exposed by the Go application, each associated with specific controller functions for handling various actions and operations. Additionally, some endpoints are protected with JWT authentication using middleware (`middlewares.JWTAuthMiddleware()`).
