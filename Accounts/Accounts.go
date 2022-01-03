package Accounts

import (
	"github.com/pborman/uuid"
	"time"
)

type AccountInterface interface {
	Create() (Account, error)
	Load() ([]Account, error)
	ChangeStatus() (bool, error)
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

func New() AccountInterface {
	result := new(Account)
	result.ID = uuid.NewRandom()
	result.CreateDate = time.Now()
	return result
}
func (a *Account) Create() (Account, error) {
	return *a, nil
}
func (a *Account) Load() ([]Account, error) {
	return nil, nil
}
func (a *Account) ChangeStatus() (bool, error) {
	return false, nil
}
