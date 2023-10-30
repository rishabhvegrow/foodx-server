package cart

import (
    "strconv"

	"gorm.io/gorm"
	"foodx-server/domain"

    "foodx-server/utils/loggerutil"
)

var logger = loggerutil.NewLogger()

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (cart_repo *CartRepository) GetCartItemByID(itemID any)(*domain.CartItem, error){
    db := cart_repo.db
	var cartItem domain.CartItem
	if err := db.Where("id = ?", itemID).First(&cartItem).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }

	return &cartItem, nil
}


func (cart_repo *CartRepository) RemoveCartItemByID(id any)(error) {
    db := cart_repo.db
	if err := db.Where("id = ?", id).Delete(&domain.CartItem{}).Error; err != nil {
        logger.Log.Error(err)
        return err
    }
    return nil
}

func (cart_repo *CartRepository) CreateCartItem(cartItem domain.CartItem)(*domain.CartItem, error){
    db := cart_repo.db
    if err := db.Create(&cartItem).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }
	return &cartItem, nil
}


func (cart_repo *CartRepository) GetCartItemsOfAUser(userID any, isChecked bool)(*[]domain.CartItem, error){
    db := cart_repo.db
	var cartItems []domain.CartItem

    if err := db.
        Preload("FoodItem").
        Where("user_id = ? AND is_checked_out = ?", userID, isChecked).
        Find(&cartItems).Error; err != nil {
            logger.Log.Error(err)
            return nil, err
    }

	return &cartItems, nil
}

func (cart_repo *CartRepository) CheckoutCart(userID any)(*domain.Transaction, error){
    db := cart_repo.db
	tx := db.Begin()
	if tx.Error != nil {
        return nil,tx.Error
    }
	cartItems, err := cart_repo.GetCartItemsOfAUser(userID, false)

	if err != nil || len(*cartItems) == 0 {
        tx.Rollback()
        logger.Log.Error(err)
        return nil,err
	}

    parsedUserIDStr, ok := userID.(string)
    if !ok {
        tx.Rollback()
        logger.Log.Error(err)
        return nil,err
    }

    parsedUserID, err := strconv.ParseUint(parsedUserIDStr, 10, 64)
    if err != nil {
        tx.Rollback()
        logger.Log.Error(err)
        return nil,err
    }

	transaction := domain.Transaction{UserID: uint(parsedUserID), Total: 0}
	if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback()
        logger.Log.Error(err)
        return nil,err
    }

	var totalSum float32

    for _, item := range *cartItems {
        totalSum += item.Price
        item.IsCheckedOut = true
        item.TransactionID = transaction.ID
        if err := tx.Save(&item).Error; err != nil {
            tx.Rollback()
            logger.Log.Error(err)
            return nil,err
        }
    }


    transaction.Total = float64(totalSum)
    if err := tx.Save(&transaction).Error; err != nil {
        tx.Rollback()
        logger.Log.Error(err)
        return nil,err
    }

    tx.Commit()

	return &transaction,nil
}

