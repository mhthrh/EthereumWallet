package Wallet

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Accounts"
	"github.com/mhthrh/WalletServices/Currencys"
	"github.com/mhthrh/WalletServices/Customers"
	"github.com/mhthrh/WalletServices/Networks"
	"github.com/mhthrh/WalletServices/Transactions"
	"github.com/mhthrh/WalletServices/Utilitys"
	"io"
	"net/http"
)

type RestApi interface {
	Login(http.ResponseWriter, *http.Request)
}

var (
	methods     *map[string]interface{}
	user        *Customers.Customer
	account     *Accounts.Account
	transaction *Transactions.Transaction
)

func New(e *[]Utilitys.Exceptions) bool {

	if account = Accounts.New(e); account.Status != nil {
		return false
	}
	if transaction = Transactions.New(e); transaction.Status != nil {
		return false
	}
	methods = &map[string]interface{}{
		"signIn": func(i io.Reader) *Utilitys.Exceptions {
			if user = Customers.New(e); user.Status != nil {
				return Utilitys.SelectException(10000, e)
			}
			json.NewDecoder(i).Decode(&user)
			user.SignIn()
			return user.Status
		},
		"signUp": func(i io.Reader) *Utilitys.Exceptions {
			if user = Customers.New(e); user.Status != nil {
				return Utilitys.SelectException(10000, e)
			}
			json.NewDecoder(i).Decode(&user)
			user.SignUp()
			return user.Status
		},
		"createAcc": func(i io.Reader) *Utilitys.Exceptions {
			json.NewDecoder(i).Decode(&account)
			account.Create()
			return user.Status
		},
		"sendTrans": func(i io.Reader) *Utilitys.Exceptions {
			json.NewDecoder(i).Decode(&transaction)
			transaction.Send()
			return user.Status
		},
		"buyTrans": func(i io.Reader) *Utilitys.Exceptions {
			json.NewDecoder(i).Decode(&transaction)
			transaction.Buy()
			return user.Status
		},
	}
	return true
}

func PostMethod(w http.ResponseWriter, r *http.Request) {
	op, ok := mux.Vars(r)["operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, j := range *methods {
		if i == op {
			if f, ok := j.(func(io.Reader) *Utilitys.Exceptions); ok {
				if f(r.Body).Key == 0 {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode("OK")
					return
				}
			}
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("Nok")
}

func AllNetwork(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(Networks.Load())
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)
}

func AllCurrency(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(Currencys.Load())
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
func LoadTransactions(w http.ResponseWriter, r *http.Request) {

	b, _ := json.Marshal(transaction.Load())
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
func LoadAccounts(w http.ResponseWriter, r *http.Request) {

	b, _ := json.Marshal(account.Load())
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
