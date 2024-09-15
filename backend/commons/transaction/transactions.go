package transaction

import (
	"errors"
	"sync"
)

/*
	Transaction should auto commit and shall be reverted if needed
*/

// ------------------------------------

type TransactionData interface {
	UndoTransaction() error
	CommitTransaction() error
}

type Transaction struct {
	TransactionId string
	Data          TransactionData
	Commited      bool
}

type TransactionReg struct {
	reg map[string]*Transaction
	sync.RWMutex
}

func NewTransactionReg() TransactionReg {
	return TransactionReg{
		reg: make(map[string]*Transaction),
	}
}

func (t *TransactionReg) UndoTransaction(id string) error {
	_, ok := t.reg[id]
	if !ok {
		return errors.New("No transaction with this id")
	}
	err := t.reg[id].Data.UndoTransaction()
	if err != nil {
		return err
	}
	t.reg[id].Commited = false
	return nil
}

func (t *TransactionReg) CloseTransaction(id string) error {
	_, ok := t.reg[id]
	if !ok {
		return errors.New("No transaction with this id")
	}
	delete(t.reg, id)
	return nil
}

func (t *TransactionReg) AddTransaction(tr Transaction) error {
	_, ok := t.reg[tr.TransactionId]
	if ok {
		return errors.New("Already a transaction with this Id")
	}
	t.reg[tr.TransactionId] = new(Transaction)
	*t.reg[tr.TransactionId] = tr
	return nil
}

var TransactionRegistry TransactionReg = NewTransactionReg()
