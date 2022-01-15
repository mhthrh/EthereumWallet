package TestPackages

import (
	"github.com/mhthrh/WalletServices/Customers"
	"github.com/mhthrh/WalletServices/Utilitys"
	"testing"
)

func TestCustomer_SignUp(t *testing.T) {
	a := Utilitys.RaiseError()

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
				Status:    nil,
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
				Status:    nil,
			},
		},
	}
	for _, tt := range tests {
		c := Customers.New(a)

		c.FirstName = tt.fields.FirstName
		c.LastName = tt.fields.LastName
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password
		c.CellNo = tt.fields.CellNo
		c.Email = tt.fields.Email
		c.Status = tt.fields.Status

		c.SignUp()
	}
}
func TestCustomer_SignIn(t *testing.T) {
	a := Utilitys.RaiseError()

	tests := []struct {
		name   string
		fields *Customers.Customer
	}{
		{
			name: "test1",
			fields: &Customers.Customer{
				UserName: "mhthrh",
				Password: "Qaz@123789",
				Status:   nil,
			},
		},
		{
			name: "test2",
			fields: &Customers.Customer{
				UserName: "mhthrh1",
				Password: "Qaz@123789",
				Status:   nil,
			},
		},
	}
	for _, tt := range tests {
		c := Customers.New(a)

		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password

		c.SignIn()
	}
}
func TestCustomer_ChangePassword(t *testing.T) {
	a := Utilitys.RaiseError()

	tests := []struct {
		name   string
		fields *Customers.Customer
	}{
		{
			name: "test1",
			fields: &Customers.Customer{
				UserName: "mhthrh",
				Password: "Qaz@123789",
				Status:   nil,
			},
		},
		{
			name: "test2",
			fields: &Customers.Customer{
				UserName: "mhthrh1",
				Password: "Qaz@123789",
				Status:   nil,
			},
		},
	}
	for _, tt := range tests {
		c := Customers.New(a)

		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password

		c.ChangePassword("newPassword1")
	}
}
