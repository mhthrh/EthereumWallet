package Accounts

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
	"time"
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
	CreateDate    time.Time
	exceptions    *[]Utilitys.Exceptions
	Status        *Utilitys.Exceptions
}

var (
	//Ether *Ethereum.Ether
	db *DbUtils.GreSQLResult
)

func init() {
	db = DbUtils.NewConnection(nil)
}

func New(e *[]Utilitys.Exceptions) *Account {
	result := new(Account)
	result.Id = uuid.NewRandom()
	result.CreateDate = time.Now()
	result.AccountName = "New Account"
	result.Balance = 0
	result.AccountString = ""
	result.CreateDate = Utilitys.GetDate("15/10/2020")
	result.exceptions = e
	return result
}
func (a *Account) Create() {
	if a == nil {
		a.Status = Utilitys.SelectException(10000, a.exceptions)
		return
	}
	db.Command = fmt.Sprintf("insert into ........")
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		a.Status = Utilitys.SelectException(10000, a.exceptions)
		return
	}
	a.Status = Utilitys.SelectException(0, a.exceptions)
}
func (a *Account) Load() *[]Account {
	var result []Account
	var account Account
	db.Command = fmt.Sprintf("select  ........")
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		a.Status = Utilitys.SelectException(10000, a.exceptions)
		return nil
	}
	for db.ResultSet.(*sql.Rows).Next() {
		err := db.ResultSet.(*sql.Rows).Scan(&account.Id, account.CustomerID, account.AccountName, account.PrivateKey, account.AccountString, account.Balance, account.CreateDate)
		if err != nil {
			a.Status = Utilitys.SelectException(10000, a.exceptions)
			return nil
		}
		result = append(result, account)

	}
	a.Status = Utilitys.SelectException(0, a.exceptions)
	return &result
}
