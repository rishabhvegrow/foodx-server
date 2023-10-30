package transactioncontroller

import (
	"foodx-server/domain"
	"foodx-server/repository/transaction"
)

type TransactionController struct {
	transactRepo *transaction.Service
}

func NewTransactionController(transactRepo *transaction.Service) *TransactionController {
	return &TransactionController{transactRepo: transactRepo}
}

func (tc *TransactionController) GetTransactionsOfUser(userID any)(*[]domain.Transaction, error){
	transactions, err := tc.transactRepo.GetTransactionsOfUser(userID)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tc *TransactionController) CreateTransaction(userID any, total float64) (*domain.Transaction, error){
	transaction, err := tc.transactRepo.CreateTransaction(userID, total)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (tc *TransactionController) UpdateTransaction(transactionID any, total float64) (*domain.Transaction, error){
	transaction, err := tc.transactRepo.UpdateTransaction(transactionID, total)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}