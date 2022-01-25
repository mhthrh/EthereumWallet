package Transactions

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/pborman/uuid"
)

type TransactionInterface interface {
	Send()
	Buy()
	Load() *[]Transaction
}
type Transaction struct {
	ID          uuid.UUID
	customerId  uuid.UUID
	AccountId   uuid.UUID
	Amount      float64
	Destination string
	TransDate   string
}

func New() (*Transaction, *Utilitys.LogInstance) {
	r := new(Transaction)
	r.ID = uuid.NewRandom()
	r.TransDate = Utilitys.GetDate("date")
	return r, nil
}
func (t *Transaction) Send() *Utilitys.LogInstance {
	return nil
}

func (t *Transaction) Buy() *Utilitys.LogInstance {
	return nil
}

func (t *Transaction) Load() (*[]Transaction, *Utilitys.LogInstance) {
	return nil, nil
}
