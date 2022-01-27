package Transactions

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/pborman/uuid"
)

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
