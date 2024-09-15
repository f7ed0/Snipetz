package transactions

import "snipetz/auth/models"

type RegisterData struct {
	models.User
}

func (d RegisterData) CommitTransaction() error {
	// TODO commit transaction to DB
	return nil
}

func (d RegisterData) UndoTransaction() error {
	// TODO Revert transaction to DB
	return nil
}
