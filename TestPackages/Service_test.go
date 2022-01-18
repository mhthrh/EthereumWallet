package TestPackages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type SignUp struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	CellNo    string
	Email     string
	Status    interface{}
}

func Test_SignUp(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "Qaz@123456",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/wallet/login/signUp", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_SignIn(t *testing.T) {
	person := &SignUp{
		UserName: "mhthrh",
		Password: "Qaz@123456",
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/wallet/login/signIn", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_CreateAcc(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_sendTrans(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_BuyTrans(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_AllNetwork(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}

func Test_AllCurrency(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_LoadTransactions(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
func Test_LoadAccounts(t *testing.T) {
	person := &SignUp{
		FirstName: "Mohsen",
		LastName:  "Taheri",
		UserName:  "mhthrh",
		Password:  "654321",
		CellNo:    "09190352044",
		Email:     "mhthrh@gmail.com",
		Status:    nil,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&person)
	res, err := http.Post("http://127.0.0.1:8585/api/Wallet/login/signin", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error: %v | Status: %s", err, res.Status)
	}
	fmt.Println("Ok")

}
