package TestPackages

import (
	"github.com/mhthrh/WalletServices/Accounts"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/pborman/uuid"
	"reflect"
	"testing"
	"time"
)

func TestAccount_Create(t *testing.T) {
	type fields struct {
		Id            uuid.UUID
		CustomerID    uuid.UUID
		AccountName   string
		AccountType   bool
		PrivateKey    string
		AccountString string
		Balance       float64
		CreateDate    time.Time
		exceptions    *[]Utilitys.Exceptions
		Status        *Utilitys.Exceptions
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Accounts.Account{
				Id:            tt.fields.Id,
				CustomerID:    tt.fields.CustomerID,
				AccountName:   tt.fields.AccountName,
				AccountType:   tt.fields.AccountType,
				PrivateKey:    tt.fields.PrivateKey,
				AccountString: tt.fields.AccountString,
				Balance:       tt.fields.Balance,
				CreateDate:    tt.fields.CreateDate,
				Status:        tt.fields.Status,
			}
			a.Create()
		})
	}
}

func TestAccount_Load(t *testing.T) {
	type fields struct {
		Id            uuid.UUID
		CustomerID    uuid.UUID
		AccountName   string
		AccountType   bool
		PrivateKey    string
		AccountString string
		Balance       float64
		CreateDate    time.Time
		exceptions    *[]Utilitys.Exceptions
		Status        *Utilitys.Exceptions
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]Accounts.Account
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Accounts.Account{
				Id:            tt.fields.Id,
				CustomerID:    tt.fields.CustomerID,
				AccountName:   tt.fields.AccountName,
				AccountType:   tt.fields.AccountType,
				PrivateKey:    tt.fields.PrivateKey,
				AccountString: tt.fields.AccountString,
				Balance:       tt.fields.Balance,
				CreateDate:    tt.fields.CreateDate,
				Status:        tt.fields.Status,
			}
			if got := a.Load(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
