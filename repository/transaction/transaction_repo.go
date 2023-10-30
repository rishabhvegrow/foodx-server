package transaction

import (
	"gorm.io/gorm"
	"foodx-server/domain"

    "foodx-server/utils/loggerutil"
)

var logger = loggerutil.NewLogger()

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (transaction_repo *TransactionRepository) GetTransactionsOfUser(userID any)(*[]domain.Transaction, error){
    db := transaction_repo.db
	var transactions []domain.Transaction
	if err := db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil{
        logger.Log.Error(err)
		return nil, err
	}
	return &transactions, nil
}

func (transaction_repo *TransactionRepository) CreateTransaction(userID any, total float64) (*domain.Transaction, error) {
    db := transaction_repo.db

    transaction := domain.Transaction{
        UserID: userID.(uint),
        Total:  total,
    }
    
    if err := db.Create(&transaction).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }

    return &transaction, nil
}

func (transaction_repo *TransactionRepository) UpdateTransaction(transactionID any, total float64) (*domain.Transaction, error) {
    db := transaction_repo.db

    var existingTransaction domain.Transaction
    if err := db.First(&existingTransaction, transactionID).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }

    existingTransaction.Total = total
    if err := db.Save(&existingTransaction).Error; err != nil {
        logger.Log.Error(err)
        return nil, err
    }

    return &existingTransaction, nil
}

