package transactions

import (
	"snipetz/auth/dbconnection"
	"snipetz/auth/models"
	snipetzerror "snipetz/commons/errors"
)

type RegisterData struct {
	models.User
}

func (d RegisterData) CommitTransaction() error {
	return dbconnection.Cntr.CreateUser(d.User)
}

func (d RegisterData) UndoTransaction() error {
	// TODO Revert transaction to DB
	return snipetzerror.ErrorNotImplemented
}
