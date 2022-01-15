package Wallet

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Customers"
	"net/http"
	"strings"
)

type WallInterface interface {
	SingUp(http.ResponseWriter, *http.Request)
}
type Services struct {
	User *Customers.Customer
}

func NewUser() *Services {
	return &Services{User: &Customers.Customer{
		FirstName: "",
		LastName:  "",
		UserName:  "",
		Password:  "",
		CellNo:    "",
		Email:     "",
		Status:    nil,
	}}
}

func (handler *Services) Login(w http.ResponseWriter, r *http.Request) {
	op, ok := mux.Vars(r)["operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "operation not Found.")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&handler)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not Decode request body by error: %v", err)
		return
	}

	switch strings.ToLower(op) {
	case "signin":
		handler.User.SignIn()
		if handler.User.Status.Key != 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not find by error: %v", err)
			return
		}

	case "signup":
		handler.User.SignUp()
		if handler.User.Status.Key != 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not Insert: %v", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handler.User)

}
