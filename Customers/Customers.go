package Customers

import (
	"CurrencyServices/Utilitys"
	"errors"
	"github.com/pborman/uuid"
	"time"
)

var (
	err Utilitys.CustomError
)

type CustomerInterface interface {
	SignUp() (*Customer, error)
	SignIn() (*Customer, error)
	ChangePassword() (bool, error)
}
type Customer struct {
	id            uuid.UUID
	FirstName     string
	LastName      string
	UserName      string
	Password      string
	CellNo        string
	Email         string
	lastLoginDate time.Time
	createDate    time.Time
}

func New() *Customer {
	result := new(Customer)
	result.id = uuid.NewRandom()
	result.createDate = time.Now()
	return result
}
func (c *Customer) SignUp() (*Customer, error) {
	err := Utilitys.NewError()
	if err1 := Utilitys.CheckMail(c.Email); err1 != true {
		err.Code = Utilitys.Email
		errors.New("mail address is not valid")
	}
	if err1 := Utilitys.CheckPassword(c.Password); err1 != true {
		return nil, errors.New("password is not valid")
	}
	if err1 := Utilitys.CheckPhoneNumber(c.CellNo); err1 != true {
		return nil, errors.New("phone number is not valid")
	}
	return c, nil
}
func (c *Customer) SignIn() (*Customer, error) {
	return c, nil
}
func (c *Customer) ChangePassword() (bool, error) {
	return true, nil
}
