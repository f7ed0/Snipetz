package authhandlers

import (
	"net/http"
	"snipetz/auth/dbconnection"
	"snipetz/auth/models"
	"snipetz/auth/transactions"
	"snipetz/commons/schema"
	"snipetz/commons/transaction"

	"github.com/f7ed0/golog/lg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	// TODO Use the schema to transmit error kind
	var content schema.AuthRegisterForm
	err := c.ShouldBindBodyWithJSON(&content)
	if err != nil {
		lg.Info.Println("Request rejected, malformed json.", err.Error())
		c.Status(400)
		return
	}

	lg.Debug.Println(content)

	if !content.AllFieldValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	u, err := models.GenerateUserFromRegisterForm(content)

	if err != nil {
		lg.Error.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	isValid, err := dbconnection.Cntr.UserCreationValid(u)

	if err != nil {
		lg.Error.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if !isValid {
		c.JSON(http.StatusOK, schema.AuthRegisterResponse{
			Status:        "invalid",
			InvalidReason: "all",
		})
		return
	}

	tr := transaction.Transaction{
		TransactionId: uuid.NewString(),
		Data:          transactions.RegisterData{User: u},
		Commited:      false,
	}

	if tr.Data.CommitTransaction() != nil {
		c.JSON(http.StatusOK, schema.AuthRegisterResponse{
			Status: "error",
		})
	}

	transaction.TransactionRegistry.AddTransaction(tr)

	c.JSON(http.StatusOK, schema.AuthRegisterResponse{
		Status: "valid",
		Uuid:   u.Uid,
		DefaultMSResponse: schema.DefaultMSResponse{
			TransactionId: tr.TransactionId,
		},
	})
}
