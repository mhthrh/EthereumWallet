package Transactions

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/pborman/uuid"
	"time"
)

type TransactionInterface interface {
	Send()
	Buy()
	Load()
}
type Transaction struct {
	ID          uuid.UUID
	customerId  uuid.UUID
	AccountId   uuid.UUID
	Amount      float64
	Destination string
	TransDate   time.Time
	exception   *[]Utilitys.Exceptions
	status      *Utilitys.Exceptions
}

func New(e *[]Utilitys.Exceptions) *Transaction {
	r := new(Transaction)
	r.ID = uuid.NewRandom()
	r.exception = e
	r.TransDate = Utilitys.GetDate("20/01/2020")
	return r
}
func (t *Transaction) Send() {
	return
}

func (t *Transaction) Buy() {
	return
}

func (t *Transaction) Load() {
	return
}
