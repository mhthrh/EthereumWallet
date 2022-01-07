package Customers_test

import (
	"CurrencyServices/Customers"
	"CurrencyServices/Utilitys"
	"fmt"
	"testing"
)

var (
	c Customers.Customer
)

func TestSingUp(t *testing.T) {

	c := Customers.New(Utilitys.RaiseError())
	d, err := c.SignUp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
