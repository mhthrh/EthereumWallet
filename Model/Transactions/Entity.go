package Transactions

import "github.com/pborman/uuid"

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
