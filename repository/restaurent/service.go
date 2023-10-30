package restaurent

import (
	"foodx-server/domain"
)

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) GetRestaurents()(*[]domain.Restaurant, error) {
	return s.repo.GetRestaurents()
}

func (s *Service) GetRestaurent(restID any)(*domain.Restaurant, error){
	return s.repo.GetRestaurent(restID)
}

func (s *Service) CraeteRestaurent(restaurent domain.Restaurant)(*domain.Restaurant, error){
	return s.repo.CraeteRestaurent(restaurent)
}

func (s *Service) DeleteRestaurent(restID any)(error){
	return s.repo.DeleteRestaurent(restID)
}