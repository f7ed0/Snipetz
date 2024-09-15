package handlers

import (
	"net/http"
	"snipetz/commons/transaction"

	"github.com/gin-gonic/gin"
)

func TransactionRevert(c *gin.Context) {
	Qarr, ok := c.GetQueryArray("transaction_id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	trid := Qarr[0]
	err := transaction.TransactionRegistry.UndoTransaction(trid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusNoContent)
}

func TransactionClose(c *gin.Context) {
	Qarr, ok := c.GetQueryArray("transaction_id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	trid := Qarr[0]
	err := transaction.TransactionRegistry.CloseTransaction(trid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusNoContent)
}
