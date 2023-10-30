package user

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

func (s *Service) GetUsers()(*[]domain.User, error){
	return s.repo.GetUsers()
}

func (s *Service) GetUser(userID any)(*domain.User, error){
	return s.repo.GetUser(userID)
}

func (s *Service) CraeteUser(user domain.User)(*domain.User, error){
	return s.repo.CraeteUser(user)
}

func (s *Service) DeleteUser(userID any)(error){
	return s.repo.DeleteUser(userID)
}

func (s *Service) UserLogin(email string, password string) (*domain.UserLoginResponse, error){
	return s.repo.UserLogin(email, password)
}