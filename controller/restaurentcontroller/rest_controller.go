package restaurentcontroller

import (
	"foodx-server/domain"
	"foodx-server/repository/restaurent"
)

type RestController struct {
	restRepo *restaurent.Service
}

func NewRestController(restRepo *restaurent.Service) *RestController {
	return &RestController{restRepo: restRepo}
}

func (rc *RestController) GetRestaurents()(*[]domain.Restaurant, error){
	restaurents, err := rc.restRepo.GetRestaurents()

	if err != nil {
		return nil, err
	}

	return restaurents, nil

}

func (rc *RestController) GetRestaurent(restID any)(*domain.Restaurant, error){
	restaurent, err := rc.restRepo.GetRestaurent(restID)

	if err != nil {
		return nil, err
	}

	return restaurent, nil
}

func (rc *RestController) CraeteRestaurent(restaurent domain.Restaurant)(*domain.Restaurant, error){
	createdRestaurent, err := rc.restRepo.CraeteRestaurent(restaurent)

	if err != nil {
		return nil, err
	}

	return createdRestaurent, nil
}

func (rc *RestController) DeleteRestaurent(restID any)(error){

	err := rc.restRepo.DeleteRestaurent(restID)

	if err != nil {
		return err
	}

	return nil
}