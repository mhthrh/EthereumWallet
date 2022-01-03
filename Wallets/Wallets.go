package Wallets

import (
	"CurrencyServices/Accounts"
	"CurrencyServices/Currencys"
	"CurrencyServices/Customers"
	"CurrencyServices/Networks"
	"CurrencyServices/Transactions"
)

type WalletInterface interface {
}

var (
	customer    Customers.CustomerInterface
	account     Accounts.AccountInterface
	currency    Currencys.CurrencyInterface
	transaction Transactions.TransactionInterface
	network     Networks.NetworkInterface
)

type Wallet struct {
	Cust  Customers.Customer
	Acc   []Accounts.Account
	Cure  []Currencys.Currency
	Trans []Transactions.Transaction
	Net   []Networks.Network
}

func New() WalletInterface {
	w := new(Wallet)
	w.Cust, _ = customer.SignUp()
	w.Acc, _ = account.Load()
	w.Cure, _ = currency.Load()
	w.Trans, _ = transaction.Load()
	w.Net, _ = network.Load()
	return w
}
