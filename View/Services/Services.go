package Services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Controler/Wallet"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/Redis"
	"net/http"
)

func RunApi() error {
	var cfg Utilitys.Config

	c, _ := Redis.New().Get("cfg")
	json.Unmarshal([]byte(c), &cfg)

	s := fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port)
	fmt.Println("initialising server on: ", s)

	router := mux.NewRouter()
	RunApiOnRouter(router)
	return http.ListenAndServe(s, router)
}

func RunApiOnRouter(r *mux.Router) {

	sub := r.PathPrefix("/api/wallet").Subrouter()
	sub.Methods("POST").Path("/{operation:(?:signIn|signUp)}").HandlerFunc(Wallet.Login)
	sub.Methods("POST").Path("/Create").HandlerFunc(Wallet.CreateAcc)
	sub.Methods("GET").Path("/load").HandlerFunc(Wallet.LoadAccounts)
	sub.Methods("POST").Path("/{operation:(?:Send|Buy)}").HandlerFunc(Wallet.DealWith)
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
}
