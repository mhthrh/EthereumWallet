package Customers

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
	"time"
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
	exceptions *[]Utilitys.Exceptions
	Status     *Utilitys.Exceptions
}

var (
	db *DbUtils.GreSQLResult
)

func init() {
	db = DbUtils.NewConnection(nil)
}
func New(e *[]Utilitys.Exceptions) *Customer {
	result := new(Customer)
	result.id = uuid.NewRandom()
	result.createDate = time.Now().String()
	result.exceptions = e
	return result
}
func (c *Customer) SignUp() {
	if err := Utilitys.CheckMail(c.Email); err != true {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	if err := Utilitys.CheckPassword(c.Password); err != true {
		c.Status = Utilitys.SelectException(10001, c.exceptions)
		return
	}
	if err := Utilitys.CheckPhoneNumber(c.CellNo); err != true {
		c.Status = Utilitys.SelectException(10002, c.exceptions)
		return
	}
	if err := Utilitys.CheckName(c.UserName); err != true {
		c.Status = Utilitys.SelectException(10003, c.exceptions)
		return
	}
	if err := Utilitys.CheckName(c.LastName); err != true {
		c.Status = Utilitys.SelectException(10004, c.exceptions)
		return
	}

	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	db.Command = fmt.Sprintf("INSERT INTO public.\"Customers\"(\"ID\", \"FirstName\", \"SureName\", \"UserName\", \"Password\", \"CellNo\", \"Email\", \"createDate\")VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", c.id, c.FirstName, c.LastName, c.UserName, c.Password, c.CellNo, c.Email, c.createDate)
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	c.Status = Utilitys.SelectException(0, c.exceptions)

}
func (c *Customer) SignIn() {
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	db.Command = fmt.Sprintf("SELECT \"ID\", \"FirstName\", \"SureName\", \"UserName\", \"Password\", \"CellNo\", \"Email\", \"createDate\" FROM public.\"Customers\" c where c.\"UserName\"='%s' and c.\"Password\"='%s'", c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&c.id, &c.FirstName, &c.LastName, &c.UserName, &c.Password, &c.CellNo, &c.Email, &c.createDate); err != nil {
			c.Status = Utilitys.SelectException(10000, c.exceptions)
			return
		}
		counter++
	}
	if counter != 1 { //Not found
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	c.Status = Utilitys.SelectException(0, c.exceptions)
}
func (c *Customer) ChangePassword(newPass string) {
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	db.Command = fmt.Sprintf("SELECT count(*) FROM public.\"Customers\" c where c.\"UserName\"='%s' and c.\"Password\"='%s'", c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}
	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&counter); err != nil {
			c.Status = Utilitys.SelectException(10000, c.exceptions)
			return
		}
	}
	if counter != 1 { //not found
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}

	db.Command = fmt.Sprintf("UPDATE public.\"Customers\" c SET  \"Password\"='%s' WHERE c.\"UserName\"='%s' and c.\"Password\"='%s' ", newPass, c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		c.Status = Utilitys.SelectException(10000, c.exceptions)
		return
	}

	c.Status = Utilitys.SelectException(0, c.exceptions)
}
