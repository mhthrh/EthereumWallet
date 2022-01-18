package Utilitys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	AppName    string `json:"AppName"`
	IsTest     bool   `json:"IsTest"`
	Version    string `json:"Version"`
	ExpireDate string `json:"ExpireDate"`
	MaxThread  int    `json:"MaxThread"`
}

func ReadConfig() *Config {
	byte, err := ioutil.ReadFile("Config.json")
	if err != nil {
		fmt.Println("Can't open Errors file!", err)
		ReadLine()
		os.Exit(0)
	}

	var jsonMap *Config
	err = json.Unmarshal(byte, &jsonMap)
	if err != nil {
		fmt.Println("Can't Unmarshal Errors file!", err)
		ReadLine()
		os.Exit(0)
	}
	return jsonMap
}
