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
	CreateDate    string
}

var (
	db  *DbUtils.GreSQLResult
	err *Utilitys.LogInstance
)

func init() {
	db, err = DbUtils.NewConnection(nil)
}

func New() (*Account, *Utilitys.LogInstance) {
	result := new(Account)
	if err != nil {
		return nil, Utilitys.Logger("NewConnection", "Db Connection error", result, err)
	}
	result.Id = uuid.NewRandom()
	result.CreateDate = time.Now().Format("21-12-2006")
	result.AccountName = "New Account"
	result.Balance = 0
	result.AccountString = ""
	return result, nil
}
func (a *Account) Create() *Utilitys.LogInstance {
	if a == nil {
		return Utilitys.Logger("Create", "Db Connection error", a, nil)
	}
	db.Command = fmt.Sprintf("insert into ........")

	if err := db.PgExecuteNonQuery(); err != nil {
		return Utilitys.Logger("Create", "Db Connection error", a, err)
	}
	return nil
}
func (a *Account) Load() (*[]Account, *Utilitys.LogInstance) {
	var result []Account
	var account Account
	db.Command = fmt.Sprintf("select  ........")

	if err := db.PgExecuteNonQuery(); err != nil {
		return nil, Utilitys.Logger("Create", "Db Connection error", a, err)
	}
	for db.ResultSet.(*sql.Rows).Next() {
		err := db.ResultSet.(*sql.Rows).Scan(&account.Id, account.CustomerID, account.AccountName, account.PrivateKey, account.AccountString, account.Balance, account.CreateDate)
		if err != nil {
			return nil, Utilitys.Logger("Create", "Db Connection error", a, err)
		}
		result = append(result, account)
	}
	return &result, nil
}
