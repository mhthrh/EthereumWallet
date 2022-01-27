package Customers

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
)

type CustomerInterface interface {
	SignUp()
	SignIn()
	ChangePassword(string)
}
type Customer struct {
	id         uuid.UUID
	FirstName  string
	LastName   string
	UserName   string
	Password   string
	CellNo     string
	Email      string
	createDate string
}

var (
	db  *DbUtils.GreSQLResult
	err *Utilitys.LogInstance
)
