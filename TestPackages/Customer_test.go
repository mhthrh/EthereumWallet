package TestPackages

import (
	"fmt"
	"github.com/mhthrh/WalletServices/Customers"
	"testing"
)

func TestCustomer_SignUp(t *testing.T) {
	tests := []struct {
		name   string
		fields *Customers.Customer
	}{
		{
			name: "test1",
			fields: &Customers.Customer{
				FirstName: "Mohsen",
				LastName:  "Taheri",
				UserName:  "mhthrh",
				Password:  "Qaz@123789",
				CellNo:    "+4477594488882",
				Email:     "mhthrh@gmail.com",
			},
		},
		{
			name: "test2",
			fields: &Customers.Customer{
				FirstName: "Mohsen2",
				LastName:  "Taheri",
				UserName:  "mhthrh",
				Password:  "Qaz@123789",
				CellNo:    "+4477594488882",
				Email:     "mhthrh@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		c, err := Customers.New()
		fmt.Println(err)
		c.FirstName = tt.fields.FirstName
		c.LastName = tt.fields.LastName
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password
		c.CellNo = tt.fields.CellNo
		c.Email = tt.fields.Email

		c.SignUp()
	}
}
func TestCustomer_SignIn(t *testing.T) {
	tests := []struct {
		name   string
		fields *Customers.Customer
	}{
		{
			name: "test1",
			fields: &Customers.Customer{
				UserName: "mhthrh",
				Password: "Qaz@123789",
			},
		},
		{
			name: "test2",
			fields: &Customers.Customer{
				UserName: "mhthrh1",
				Password: "Qaz@123789",
			},
		},
	}
	for _, tt := range tests {
		c, err := Customers.New()
		fmt.Println(err)
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password

		c.SignIn()
	}
}
func TestCustomer_ChangePassword(t *testing.T) {

	tests := []struct {
		name   string
		fields *Customers.Customer
	}{
		{
			name: "test1",
			fields: &Customers.Customer{
				UserName: "mhthrh",
				Password: "Qaz@123789",
			},
		},
		{
			name: "test2",
			fields: &Customers.Customer{
				UserName: "mhthrh1",
				Password: "Qaz@123789",
			},
		},
	}
	for _, tt := range tests {
		c, err := Customers.New()
		fmt.Println(err)
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password

		c.ChangePassword("newPassword1")
	}
}
