package main

import (
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/View/Services"
)

func main() {
	//Utility's.WriteConfig()
	Utilitys.SetConsoleTitle("Crypto Services")
	cfg := Utilitys.ReadConfig("ApplicationFiles/ConfigCoded.json")
	s := fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port)
	fmt.Println("initialising server on: ", s)
	Services.RunApi(s)
}
