package Accounts

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
)

type AccountInterface interface {
	Create()
	Load()
}
type Account struct {
	Id            uuid.UUID
	CustomerID    uuid.UUID
	AccountName   string
	AccountType   bool
	PrivateKey    string
	Byt           string
	AccountString string
	Balance       float64
	CreateDate    string
}

var (
	db  *DbUtils.GreSQLResult
	err *Utilitys.LogInstance
)
