package Customers

import (
	"github.com/pborman/uuid"
	"time"
)

type CustomerInterface interface {
	SignUp() (Customer, error)
	SignIn() (Customer, error)
	ChangePassword() (bool, error)
}
type Customer struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	UserName      string
	Password      string
	CellNo        string
	Email         string
	LastLoginDate time.Time
	CreateDate    time.Time
}

func New() CustomerInterface {
	result := new(Customer)
	result.ID = uuid.NewRandom()
	result.CreateDate = time.Now()
	return result
}
func (c *Customer) SignUp() (Customer, error) {
	return *c, nil
}
func (c *Customer) SignIn() (Customer, error) {
	return *c, nil
}
func (c *Customer) ChangePassword() (bool, error) {
	return true, nil
}
