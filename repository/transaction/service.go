package transaction

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

func (s *Service) GetTransactionsOfUser(userID any)(*[]domain.Transaction, error){
	return s.repo.GetTransactionsOfUser(userID)
}

func (s *Service) CreateTransaction(userID any, total float64) (*domain.Transaction, error){
	return s.repo.CreateTransaction(userID, total)
}

func (s *Service) UpdateTransaction(transactionID any, total float64) (*domain.Transaction, error){
	return s.repo.UpdateTransaction(transactionID, total)
}