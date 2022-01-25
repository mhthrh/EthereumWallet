package Services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Wallet"
	"net/http"
)

func RunApi(endpoint string) error {
	router := mux.NewRouter()
	if !RunApiOnRouter(router) {
		return errors.New("fucking wallet Not initialize")
	}
	return http.ListenAndServe(endpoint, router)
}

func RunApiOnRouter(r *mux.Router) bool {
	if !Wallet.New() {
		return false
	}
	sub := r.PathPrefix("/api/wallet").Subrouter()

	sub.Methods("POST").Path("/{operation:(?:signIn|signUp)}").HandlerFunc(Wallet.PostMethod)
	sub.Methods("POST").Path("/Create").HandlerFunc(Wallet.PostMethod)
	sub.Methods("GET").Path("/load").HandlerFunc(Wallet.LoadAccounts)
	sub.Methods("POST").Path("/{operation:(?:Send|Buy)}").HandlerFunc(Wallet.PostMethod)
	sub.Methods("GET").Path("/load").HandlerFunc(Wallet.LoadTransactions)
	sub.Methods("GET").Path("/allNetwork").HandlerFunc(Wallet.AllNetwork)
	sub.Methods("GET").Path("/allCurrency").HandlerFunc(Wallet.AllCurrency)
	sub.Methods("POST").Path("/getTicket").HandlerFunc(Wallet.GetTicket)
	sub.Methods("POST").Path("/TicketValidation").HandlerFunc(Wallet.TicketIsValid)

	r.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			b, _ := json.Marshal("Utilitys.SelectException(10000, Exc)")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s", b)
		})
	return true
}
