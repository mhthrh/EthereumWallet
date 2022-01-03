package Transactions

import (
	"github.com/pborman/uuid"
	"time"
)

type TransactionInterface interface {
	Send() (Transaction, error)
	Buy() (Transaction, error)
	Load() ([]Transaction, error)
}
type Transaction struct {
	ID          uuid.UUID
	Amount      float64
	Destination string
	TransDate   time.Time
	Status      string
}

func New() TransactionInterface {
	r := new(Transaction)
	r.ID = uuid.NewRandom()
	r.TransDate = time.Now()
	return r
}
func (t *Transaction) Send() (Transaction, error) {
	return *t, nil
}

func (t *Transaction) Buy() (Transaction, error) {
	return *t, nil
}

func (t *Transaction) Load() ([]Transaction, error) {
	return nil, nil
}
