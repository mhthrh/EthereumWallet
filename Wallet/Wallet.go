package Wallet

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Customers"
	"github.com/mhthrh/WalletServices/Utilitys"
	"io"
	"net/http"
)

type RestApi interface {
	Login(http.ResponseWriter, *http.Request)
}

var (
	methods map[string]interface{}
)

type RestApiHandler struct {
	User *Customers.Customer
}

func New(e *[]Utilitys.Exceptions) {
	methods = map[string]interface{}{
		"signin": func(i io.Reader) *Utilitys.Exceptions {
			a := Customers.New(e)
			if a.Status.Key != 0 {
				return a.Status
			}
			json.NewDecoder(i).Decode(&a)
			a.SignIn()
			return a.Status

		},
		"signup": func(i io.Reader) *Utilitys.Exceptions {
			a := Customers.New(e)
			if a.Status.Key != 0 {
				return a.Status
			}
			json.NewDecoder(i).Decode(&a)
			a.SignUp()
			return a.Status
		},
	}
}

func (handler *RestApiHandler) PostMethod(w http.ResponseWriter, r *http.Request) {
	op, ok := mux.Vars(r)["operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, j := range methods {
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
