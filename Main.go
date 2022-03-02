package main

import (
	"encoding/json"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/Redis"
	"github.com/mhthrh/WalletServices/View/Services"
)

func main() {
	//Utility's.WriteConfig()
	Utilitys.SetConsoleTitle("Crypto Services")

	cfg := Utilitys.ReadConfig("ApplicationFiles/ConfigCoded.json")
	b, _ := json.Marshal(cfg)

	r := Redis.New()
	if r.Ping() != nil {
		return
	}
	r.Set("cfg", string(b))

	Services.RunApi()
}
