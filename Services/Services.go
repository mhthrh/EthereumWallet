package Services

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/WalletServices/Wallet"
	"net/http"
)

func RunApi(endpoint string) error {
	r := mux.NewRouter()
	RunApiOnRouter(r)
	fmt.Println("Server Started ...")
	return http.ListenAndServe(endpoint, r)
}

func RunApiOnRouter(r *mux.Router) {
	apiRouter := r.PathPrefix("/api/Wallet").Subrouter()
	apiRouter.Methods("POST").Path("/{operation:(?:SignIn|SignUp)}").HandlerFunc(Wallet.NewUser().Login)
	//apiRouter.Methods("GET").Path("/SignIn/{search}").HandlerFunc(Wallet.NewUser().SingUp)

}
