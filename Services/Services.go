package Services

import (
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
	RunApiOnRouter(router)
	return http.ListenAndServe(endpoint, router)
}

func RunApiOnRouter(r *mux.Router) {
	var handler Wallet.RestApiHandler
	Wallet.New(Exc)
	r.PathPrefix("/api/Wallet/login").Subrouter().Methods("POST").Path("/{operation:(?:signin|signup)}").HandlerFunc(handler.PostMethod)
	r.PathPrefix("/api/Wallet/account").Subrouter().Methods("POST").Path("/{operation:(?:create|load)}").HandlerFunc(handler.PostMethod)
}
