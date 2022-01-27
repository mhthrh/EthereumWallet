package TestPackages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_GetTicket(t *testing.T) {
	type Request struct {
		UserName string
		Password string
		IsTest   bool
	}
	var response interface{}
	req := &Request{
		UserName: "Mohsen",
		Password: "3456",
		IsTest:   true,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&req)
	res, err := http.Post("http://127.0.0.1:8585/api/wallet/getTicket", "application/json", &data)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &response)

	fmt.Println(err)
	fmt.Println(req)
	fmt.Println(response)
	fmt.Println(res.StatusCode)

}

func Test_IsValid(t *testing.T) {
	type Request struct {
		UserName    string
		Password    string
		InputTicket string
		IsTest      bool
	}
	var response interface{}
	req := &Request{
		UserName:    "Mohsen",
		Password:    "",
		InputTicket: "123456",
		IsTest:      true,
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&req)
	res, err := http.Post("http://127.0.0.1:8585/api/wallet/TicketValidation", "application/json", &data)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &response)

	fmt.Println(err)
	fmt.Println(req)
	fmt.Println(response)
	fmt.Println(res.StatusCode)

}
