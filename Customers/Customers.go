package Customers

import (
	"CurrencyServices/Utilitys"
	"CurrencyServices/Utilitys/DbUtils"
	"fmt"
	"github.com/pborman/uuid"
	"time"
)

type CustomerInterface interface {
	SignUp() (*Customer, error)
	SignIn() (*Customer, error)
	ChangePassword() (bool, error)
}
type Customer struct {
	id         uuid.UUID
	FirstName  string
	LastName   string
	UserName   string
	Password   string
	CellNo     string
	Email      string
	createDate time.Time
	exception  *[]Utilitys.JsonExceptions
}

func New(e *[]Utilitys.JsonExceptions) *Customer {
	result := new(Customer)
	result.id = uuid.NewRandom()
	result.createDate = time.Now()
	result.exception = e
	return result
}
func (c *Customer) SignUp() (*Customer, *Utilitys.JsonExceptions) {
	if err := Utilitys.CheckMail(c.Email); err != true {
		return nil, Utilitys.SelectException(10000, c.exception)
	}
	if err := Utilitys.CheckPassword(c.Password); err != true {
		return nil, Utilitys.SelectException(10001, c.exception)
	}
	if err := Utilitys.CheckPhoneNumber(c.CellNo); err != true {
		return nil, Utilitys.SelectException(10002, c.exception)
	}
	if err := Utilitys.CheckName(c.UserName); err != true {
		return nil, Utilitys.SelectException(10003, c.exception)
	}
	if err := Utilitys.CheckName(c.LastName); err != true {
		return nil, Utilitys.SelectException(10004, c.exception)
	}
	db := DbUtils.NewConnection(nil)
	db.Command = fmt.Sprintf("INSERT INTO public.Customers(ID, FirstName, SureName, UserName, Password, CellNo, Email, createDate)VALUES (%s, %s, %s, %s, %s, %s, %s, %s)", c.id, c.FirstName, c.LastName, c.UserName, c.Password, c.CellNo, c.Email, c.createDate)
	db.PgExecuteNonQuery()
	if db.Status.Key != 0 {
		return nil, Utilitys.SelectException(10009, c.exception)
	}

	return c, Utilitys.SelectException(0, c.exception)
}
func (c *Customer) SignIn() (*Customer, error) {
	return c, nil
}
func (c *Customer) ChangePassword() (bool, error) {
	return true, nil
}
