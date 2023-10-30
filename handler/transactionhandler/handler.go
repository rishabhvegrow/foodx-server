package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"foodx-server/controller/transactioncontroller"
)


type TransactionHandler struct {
	transactionController transactioncontroller.Controller
}

func NewTransactionHandler(transactionController transactioncontroller.Controller) *TransactionHandler {
	return &TransactionHandler{transactionController: transactionController}
}

func (th *TransactionHandler) GetTransactions(c *gin.Context){

    userID := c.MustGet("user_id").(uint)

    transactions, err := th.transactionController.GetTransactionsOfUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
        return
    }
    c.JSON(http.StatusOK, transactions)
}