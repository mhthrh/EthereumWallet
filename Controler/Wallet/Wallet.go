package Wallet

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Model/Accounts"
	"github.com/mhthrh/WalletServices/Model/Counsumers"
	"github.com/mhthrh/WalletServices/Model/Currencys"
	"github.com/mhthrh/WalletServices/Model/Customers"
	"github.com/mhthrh/WalletServices/Model/Networks"
	"github.com/mhthrh/WalletServices/Model/Transactions"
	"github.com/mhthrh/WalletServices/Utilitys"
	"io"
	"net/http"
)

var (
	methods *map[string]interface{}
)

func New() bool {

	methods = &map[string]interface{}{
		"signIn": func(i io.Reader) *Utilitys.LogInstance {
			user, err := Customers.New()
			if err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			if err := json.NewDecoder(i).Decode(&user); err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			if err := user.SignIn(); err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			return nil
		},
		"signUp": func(i io.Reader) *Utilitys.LogInstance {
			user, err := Customers.New()
			if err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			if err := json.NewDecoder(i).Decode(&user); err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			if err := user.SignUp(); err != nil {
				return Utilitys.Logger("signIn", "Error", user, err)
			}
			return nil
		},
		"createAcc": func(i io.Reader) *Utilitys.LogInstance {
			account, err := Accounts.New()
			if err != nil {
				return Utilitys.Logger("signIn", "Error", account, err)
			}
			if err := json.NewDecoder(i).Decode(&account); err != nil {
				return Utilitys.Logger("signIn", "Error", account, err)
			}
			if err := account.Create(); err != nil {
				return Utilitys.Logger("signIn", "Error", account, err)
			}
			return nil
		},
		"sendTrans": func(i io.Reader) *Utilitys.LogInstance {
			transaction, err := Transactions.New()
			if err != nil {
				return Utilitys.Logger("sendTrans", "Error", transaction, err)
			}
			if err := json.NewDecoder(i).Decode(&transaction); err != nil {
				return Utilitys.Logger("sendTrans", "Error", transaction, err)
			}

			if err := transaction.Send(); err != nil {
				return Utilitys.Logger("sendTrans", "Error", transaction, err)
			}
			return nil
		},
		"buyTrans": func(i io.Reader) *Utilitys.LogInstance {
			transaction, err := Transactions.New()
			if err != nil {
				return Utilitys.Logger("buyTrans", "Error", transaction, err)
			}
			if err := json.NewDecoder(i).Decode(&transaction); err != nil {
				return Utilitys.Logger("buyTrans", "Error", transaction, err)
			}
			if err := transaction.Buy(); err != nil {
				return Utilitys.Logger("buyTrans", "Error", transaction, err)
			}

			return nil
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
			if f, ok := j.(func(io.Reader) *Utilitys.LogInstance); ok {
				if f(r.Body) == nil {
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
	transaction, err := Transactions.New()
	if err != nil {
		return
	}
	t, err := transaction.Load()
	if err != nil {
		return
	}
	b, _ := json.Marshal(t)
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
func LoadAccounts(w http.ResponseWriter, r *http.Request) {
	//var account Accounts.Account
	account, err := Accounts.New()
	a, err := account.Load()
	if err != nil {

	}
	b, _ := json.Marshal(a)
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
func GetTicket(w http.ResponseWriter, r *http.Request) {
	t := Consumers.New()
	if err := json.NewDecoder(r.Body).Decode(&t.InputParameter); err != nil {
		t.OutputParameter.Ticket = ""
		t.OutputParameter.Result.Code = "32"
		t.OutputParameter.Result.Description = "Error with your input"
		b, _ := json.Marshal(t.OutputParameter)
		fmt.Fprintf(w, "%s", b)
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := t.GetTicket(); err != nil {
		t.OutputParameter.Ticket = ""
		t.OutputParameter.Result.Code = "33"
		t.OutputParameter.Result.Description = "Error with your input date"
		b, _ := json.Marshal(t.OutputParameter)
		fmt.Fprintf(w, "%s", b)
		w.WriteHeader(http.StatusOK)
		return
	}
	b, _ := json.Marshal(t.OutputParameter)
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
func TicketIsValid(w http.ResponseWriter, r *http.Request) {
	t := Consumers.New()

	if err := t.IsValid(); err != nil {
		return
	}
	json.NewDecoder(r.Body).Decode(&t.InputParameter)
	t.IsValid()
	t.OutputParameter.Ticket = t.InputParameter.InputTicket
	b, _ := json.Marshal(t.OutputParameter)
	fmt.Fprintf(w, "%s", b)
	w.WriteHeader(http.StatusOK)

}
