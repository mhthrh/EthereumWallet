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
}

var (
	db  *DbUtils.GreSQLResult
	err *Utilitys.LogInstance
)

func New() (*Customer, *Utilitys.LogInstance) {
	if db, err = DbUtils.NewConnection(nil); err != nil {
		return nil, Utilitys.Logger("New", "Connection Error", db, err)
	}
	result := new(Customer)
	result.id = uuid.NewRandom()
	result.createDate = time.Now().String()
	return result, nil
}
func (c *Customer) SignUp() *Utilitys.LogInstance {

	f := func(m string) *Utilitys.LogInstance {
		return Utilitys.Logger("SignUp", m, c, err)
	}
	if err = Utilitys.CheckMail(c.Email); err != nil {
		f("Invalid Email")
	}
	if err = Utilitys.CheckPassword(c.Password); err != nil {
		f("Invalid Password")
	}
	if err = Utilitys.CheckPhoneNumber(c.CellNo); err != nil {
		f("Invalid Phone number")
	}
	if err = Utilitys.CheckName(c.UserName); err != nil {
		f("Invalid User Name")
	}
	if err = Utilitys.CheckName(c.LastName); err != nil {
		f("Invalid Last Name")
	}

	db.Command = fmt.Sprintf("INSERT INTO public.\"Customers\"(\"ID\", \"FirstName\", \"SureName\", \"UserName\", \"Password\", \"CellNo\", \"Email\", \"createDate\")VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", c.id, c.FirstName, c.LastName, c.UserName, c.Password, c.CellNo, c.Email, c.createDate)

	if err = db.PgExecuteNonQuery(); err != nil {
		f("Invalid SQL Command")
	}
	return nil
}
func (c *Customer) SignIn() *Utilitys.LogInstance {
	f := func(m string) *Utilitys.LogInstance {
		return Utilitys.Logger("SignIn", m, c, err)
	}
	db.Command = fmt.Sprintf("SELECT \"ID\", \"FirstName\", \"SureName\", \"UserName\", \"Password\", \"CellNo\", \"Email\", \"createDate\" FROM public.\"Customers\" c where c.\"UserName\"='%s' and c.\"Password\"='%s'", c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if err = db.PgExecuteNonQuery(); err != nil {
		f("SQL command Error")
	}
	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&c.id, &c.FirstName, &c.LastName, &c.UserName, &c.Password, &c.CellNo, &c.Email, &c.createDate); err != nil {
			f("SQL command Error")
		}
		counter++
	}
	if counter > 1 { // Error in Result
		f("SQL command Error")
	}
	return nil
}
func (c *Customer) ChangePassword(newPass string) *Utilitys.LogInstance {
	f := func(m string) *Utilitys.LogInstance {
		return Utilitys.Logger("ChangePassword", m, c, err)
	}
	db.Command = fmt.Sprintf("SELECT count(*) FROM public.\"Customers\" c where c.\"UserName\"='%s' and c.\"Password\"='%s'", c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if err = db.PgExecuteNonQuery(); err != nil {
		f("SQL command Error")
	}
	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&counter); err != nil {
			return Utilitys.Logger("ChangePassword", "Result error", c, err)
		}
	}
	if counter != 1 { //not found
		f("Cannot find Customer")
	}

	db.Command = fmt.Sprintf("UPDATE public.\"Customers\" c SET  \"Password\"='%s' WHERE c.\"UserName\"='%s' and c.\"Password\"='%s' ", newPass, c.UserName, c.Password)
	db.PgExecuteNonQuery()
	if err = db.PgExecuteNonQuery(); err != nil {
		f("Execute Command Error")
	}
	return nil
}
