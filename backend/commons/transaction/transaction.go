package transaction

import "errors"

type TransactionData interface {
	UndoTransaction() error
	CommitTransaction() error
}

type Transaction struct {
	TransactionId string
	Data          TransactionData
	Commited      bool
}

func (tr Transaction) UndoTransaction() (err error) {
	if !tr.Commited {
		return errors.New("Transaction not commited")
	}
	err = tr.Data.UndoTransaction()
	if err != nil {
		return
	}
	tr.Commited = false
	return
}

func (tr Transaction) CommitTransaction() (err error) {
	if tr.Commited {
		return errors.New("Transaction already commited")
	}
	err = tr.Data.CommitTransaction()
	if err != nil {
		return
	}
	tr.Commited = true
	return
}
