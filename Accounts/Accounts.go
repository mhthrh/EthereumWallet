package accounts

import (
	"github.com/pborman/uuid"
	"time"
)

type AccountInterface interface {
	Create()
	ChangeStatus()
	AccountList()
}
type Account struct {
	ID            uuid.UUID
	AccountName   string
	AccountType   bool
	PrivateKey    string
	AccountString string
	Balance       float64
	CreateDate    time.Time
}

func New() {
	result := new(Account)
	result.ID = uuid.NewRandom()
	result.CreateDate = time.Now()
}
