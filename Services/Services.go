package Services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Wallet"
	"net/http"
)

var (
	Exc *[]Utilitys.Exceptions
)

func init() {
	Exc = Utilitys.RaiseError()
}

func RunApi(endpoint string) error {
	router := mux.NewRouter()
	if !RunApiOnRouter(router) {
		return errors.New("fucking wallet initiation")
	}
	return http.ListenAndServe(endpoint, router)
}

func RunApiOnRouter(r *mux.Router) bool {
	if !Wallet.New(Exc) {
		return false
	}

	r.PathPrefix("/api/wallet/login").Subrouter().Methods("POST").Path("/{operation:(?:signIn|signUp)}").HandlerFunc(Wallet.PostMethod)
	r.PathPrefix("/api/wallet/account").Subrouter().Methods("POST").Path("/Create").HandlerFunc(Wallet.PostMethod)
	r.PathPrefix("/api/wallet/account").Subrouter().Methods("GET").Path("/load").HandlerFunc(Wallet.LoadAccounts)
	r.PathPrefix("/api/wallet/Transaction").Subrouter().Methods("POST").Path("/{operation:(?:Send|Buy)}").HandlerFunc(Wallet.PostMethod)
	r.PathPrefix("/api/wallet/Transaction").Subrouter().Methods("GET").Path("/load").HandlerFunc(Wallet.LoadTransactions)
	r.PathPrefix("/api/wallet/network").Subrouter().Methods("GET").Path("/allNetwork").HandlerFunc(Wallet.AllNetwork)
	r.PathPrefix("/api/wallet/currency").Subrouter().Methods("GET").Path("/allCurrency").HandlerFunc(Wallet.AllCurrency)

	r.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			b, _ := json.Marshal(Utilitys.SelectException(10000, Exc))
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s", b)
		})
	return true
}
