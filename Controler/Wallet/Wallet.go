package Wallet

import (
	"encoding/json"
	"fmt"
	"github.com/mhthrh/WalletServices/Model/Accounts"
	Consumers "github.com/mhthrh/WalletServices/Model/Counsumers"
	"github.com/mhthrh/WalletServices/Model/Currencys"
	"github.com/mhthrh/WalletServices/Model/Customers"
	"github.com/mhthrh/WalletServices/Model/Networks"
	"github.com/mhthrh/WalletServices/Model/Transactions"
	"github.com/mhthrh/WalletServices/Utilitys"
	"net/http"
)

var (
	WriteResponse = func(w http.ResponseWriter, i interface{}) error {
		b, _ := json.Marshal(i)
		fmt.Fprintf(w, "%s", b)
		(w).WriteHeader(http.StatusOK)
		return nil
	}
	ReadRequest = func(r *http.Request, i interface{}) (interface{}, error) {
		return i, json.NewDecoder(r.Body).Decode(&i)
	}
	WriteException = func(w http.ResponseWriter, i interface{}, h int) {
		b, _ := json.Marshal(i)
		fmt.Fprintf(w, "%s", b)
		(w).WriteHeader(h)

	}
)

func Login(w http.ResponseWriter, r *http.Request) {
	user, err := Customers.New()
	if err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
		WriteException(w, "", http.StatusBadRequest)
	}
	obj, err1 := ReadRequest(r, user)
	if err1 != nil {
		Utilitys.Logger("signIn", "Error", user, err1)
	}
	if err := user.SignIn(); err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
	}
	if err := WriteResponse(w, obj); err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
	}

	//user, err := Customers.New()
	if err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
	}
	if err := user.SignUp(); err != nil {
		Utilitys.Logger("signIn", "Error", user, err)
	}

}

func CreateAcc(w http.ResponseWriter, r *http.Request) {
	account, err := Accounts.New()
	if err != nil {
		Utilitys.Logger("signIn", "Error", account, err)
	}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		Utilitys.Logger("signIn", "Error", account, err)
	}
	if err := account.Create(); err != nil {
		Utilitys.Logger("signIn", "Error", account, err)
	}

}
func DealWith(w http.ResponseWriter, r *http.Request) {
	transaction, err := Transactions.New()
	if err != nil {
		Utilitys.Logger("sendTrans", "Error", transaction, err)
	}
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		Utilitys.Logger("sendTrans", "Error", transaction, err)
	}
	if err := transaction.Send(); err != nil {
		Utilitys.Logger("sendTrans", "Error", transaction, err)
	}

	//transaction, err := Transactions.New()
	if err != nil {
		Utilitys.Logger("buyTrans", "Error", transaction, err)
	}
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		Utilitys.Logger("buyTrans", "Error", transaction, err)
	}
	if err := transaction.Buy(); err != nil {
		Utilitys.Logger("buyTrans", "Error", transaction, err)
	}

}
func AllNetwork(w http.ResponseWriter, r *http.Request) {
	Utilitys.WriteResponse(w, Networks.Load())
}
func AllCurrency(w http.ResponseWriter, r *http.Request) {
	Utilitys.WriteResponse(w, Currencys.Load())
}
func LoadTransactions(w http.ResponseWriter, r *http.Request) {
	transaction, err := Transactions.New()
	if err != nil {

	}
	t, err := transaction.Load()
	if err != nil {

	}
	Utilitys.WriteResponse(w, t)

}
func LoadAccounts(w http.ResponseWriter, r *http.Request) {
	//var account Accounts.Account
	account, err := Accounts.New()
	a, err := account.Load()
	if err != nil {
	}
	Utilitys.WriteResponse(w, a)

}
func GetTicket(w http.ResponseWriter, r *http.Request) {

	t := Consumers.New()
	if err := json.NewDecoder(r.Body).Decode(&t.InputParameter); err != nil {
		t.OutputParameter.Ticket = ""
		t.OutputParameter.Result.Code = "32"
		t.OutputParameter.Result.Description = "Error with your input"
		Utilitys.WriteResponse(w, t.OutputParameter)
	}
	if err := t.GetTicket(); err != nil {
		t.OutputParameter.Ticket = ""
		t.OutputParameter.Result.Code = "33"
		t.OutputParameter.Result.Description = "Error with your input date"
		Utilitys.WriteResponse(w, t.OutputParameter)
	}
	Utilitys.WriteResponse(w, t.OutputParameter)

}
func TicketIsValid(w http.ResponseWriter, r *http.Request) {

	t := Consumers.New()

	if err := t.IsValid(); err != nil {
	}
	json.NewDecoder(r.Body).Decode(&t.InputParameter)
	t.IsValid()
	t.OutputParameter.Ticket = t.InputParameter.InputTicket
	Utilitys.WriteResponse(w, t.OutputParameter)
}
