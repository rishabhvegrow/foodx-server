package restaurent

import (
	"gorm.io/gorm"
	"foodx-server/domain"

    "foodx-server/utils/loggerutil"
)

var logger = loggerutil.NewLogger()

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{
		db: db,
	}
}

func (restaurent_repo *RestaurantRepository) GetRestaurents()(*[]domain.Restaurant, error){
    db := restaurent_repo.db
	var restaurants []domain.Restaurant
    if err := db.Find(&restaurants).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }
	return &restaurants, nil
}

func (restaurent_repo *RestaurantRepository) GetRestaurent(restID any)(*domain.Restaurant, error){
    db := restaurent_repo.db
    var restaurent domain.Restaurant
    if err := db.First(&restaurent, restID).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }

	return &restaurent, nil
}

func (restaurent_repo *RestaurantRepository) CraeteRestaurent(restaurent domain.Restaurant)(*domain.Restaurant, error){
    db := restaurent_repo.db
    if err := db.Create(&restaurent).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }
	return &restaurent, nil
}

func (restaurent_repo *RestaurantRepository) DeleteRestaurent(restID any)(error){
    db := restaurent_repo.db
	var restaurent domain.Restaurant
    if err:= db.Delete(&restaurent, restID).Error; err != nil {
        logger.Log.Error(err)
        return err
    }
	return nil
}