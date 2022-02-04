package Utilitys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Users struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
type Treading struct {
	MinThread int `json:"MinThread"`
	MaxThread int `json:"MaxThread"`
}
type DataBase struct {
	Name   string `json:"Name"`
	Host   string `json:"Host"`
	Port   int    `json:"Port"`
	User   Users
	Dbname string `json:"Dbname"`
	Driver string `json:"Driver"`
}
type FTP struct {
	FtpName string `json:"FtpName"`
	IP      string `json:"IP"`
	Port    int    `json:"Port"`
	User    Users  `json:"User"`
}
type Authenticate struct {
	User Users  `json:"User"`
	Role string `json:"Role"`
}
type Serv struct {
	IP   string `json:"IP"`
	Port int    `json:"Port"`
}

type Config struct {
	AppName    string       `json:"AppName"`
	IsTest     bool         `json:"IsTest"`
	Version    string       `json:"Version"`
	ExpireDate string       `json:"ExpireDate"`
	DB         *[]DataBase  `json:"DB"`
	Ftp        []FTP        `json:"Ftp"`
	Thread     Treading     `json:"Thread"`
	Login      Authenticate `json:"Login"`
	Server     Serv         `json:"Server"`
}

func ReadConfig(file string) *Config {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't open file!", err)
		ReadLine()
		os.Exit(0)
	}
	byt, err := ioutil.ReadFile(filepath.Join(path, file))
	if err != nil {
		fmt.Println("Can't open Exceptions file!", err)
		ReadLine()
		os.Exit(0)
	}
	var jsonMap *Config
	err = json.Unmarshal(func(b []byte) []byte {
		k, _ := NewKey()
		k.Text = string(b)
		k.Decrypt()
		return []byte(k.Result)
	}(byt), &jsonMap)
	if err != nil {
		fmt.Println("Can't Unmarshal Exceptions file!", err)
		ReadLine()
		os.Exit(0)
	}
	return jsonMap
}

func WriteConfig() {
	cfg := &Config{
		AppName:    "Crypto Services",
		IsTest:     true,
		Version:    "1.0.0.1",
		ExpireDate: "01-01-2023",
		DB: &[]DataBase{{
			Name: "PostgresSQL",
			Host: "127.0.0.1",
			Port: 5432,
			User: Users{
				UserName: "postgresql",
				Password: "123456",
			},
			Dbname: "Curency",
			Driver: "postgres",
		}, {
			Name: "Oracle",
			Host: "127.0.0.1",
			Port: 1501,
			User: Users{
				UserName: "admin",
				Password: "admin",
			},
			Dbname: "",
			Driver: "OraDB",
		}},
		Ftp: []FTP{{
			FtpName: "MyFtp",
			IP:      "127.0.0.1",
			Port:    21,
			User: Users{
				UserName: "FtpUser",
				Password: "FtpPAss",
			},
		}, {
			FtpName: "YourFtp",
			IP:      "127.0.0.1",
			Port:    21,
			User: Users{
				UserName: "admin",
				Password: "admin",
			},
		}},
		Thread: Treading{
			MinThread: 1,
			MaxThread: 25,
		},
		Login: Authenticate{
			User: Users{
				UserName: "myUser",
				Password: "Password",
			},
			Role: "Admin",
		}, Server: Serv{
			IP:   "localhost",
			Port: 8585,
		},
	}
	b, _ := json.Marshal(cfg)
	k, _ := NewKey()
	k.Text = string(b)
	k.Encrypt()
	fmt.Println(fmt.Sprintf("%s\n", b), k.Result)

}
