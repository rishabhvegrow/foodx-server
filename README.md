# FoodX Server Routes

This document provides an overview of the routes and endpoints in the FoodX Server application.

## User Routes

### Get User By ID
- **Endpoint:** `GET /user/:id`
- **Description:** Retrieves user information by ID.

### Delete User By ID
- **Endpoint:** `DELETE /user/:id`
- **Description:** Deletes a user by ID.

### Get All Users
- **Endpoint:** `GET /users`
- **Description:** Retrieves a list of all users.

### User Registration
- **Endpoint:** `POST /signup`
- **Description:** Registers a new user.

### User Login
- **Endpoint:** `POST /signin`
- **Description:** Logs in a user.

## Restaurant Routes

### Get All Restaurants
- **Endpoint:** `GET /restaurents`
- **Description:** Retrieves a list of all restaurants.

### Get Restaurant By ID
- **Endpoint:** `GET /restaurents/:id`
- **Description:** Retrieves restaurant information by ID.

### Create Restaurant
- **Endpoint:** `POST /restaurents`
- **Description:** Creates a new restaurant.

### Delete Restaurant By ID
- **Endpoint:** `DELETE /restaurents/:id`
- **Description:** Deletes a restaurant by ID.

## Food Routes

### Get Menu of a Restaurant
- **Endpoint:** `GET /food/:restid`
- **Description:** Retrieves the menu of a specific restaurant.

### Create Food Item
- **Endpoint:** `POST /food`
- **Description:** Creates a new food item.

### Update Food Item
- **Endpoint:** `PUT /food/:id`
- **Description:** Updates a food item by ID.

### Delete Food Item
- **Endpoint:** `DELETE /food/:id`
- **Description:** Deletes a food item by ID.

### Add Food to Cart
- **Endpoint:** `POST /food/add/:id`
- **Description:** Adds a food item to the user's cart.

### Remove Food from Cart
- **Endpoint:** `POST /food/remove/:id`
- **Description:** Removes a food item from the user's cart.

## Cart Routes

### Get User's Cart Items
- **Endpoint:** `GET /cart`
- **Description:** Retrieves items in the user's cart.

### Remove Cart Item By ID
- **Endpoint:** `DELETE /cart/remove/:id`
- **Description:** Removes a cart item by ID.

### Get Ordered Items
- **Endpoint:** `GET /cart/orders`
- **Description:** Retrieves ordered items.

### Checkout Cart
- **Endpoint:** `POST /cart/checkout`
- **Description:** Checks out the user's cart and creates a transaction.
