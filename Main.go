package main

import (
	"flag"
	"fmt"
	"github.com/mhthrh/WalletServices/View/Services"
)

func main() {
	addr := flag.String("addr", ":8585", "the TCP address for the server to listen on, in the form 'host:port'")
	fmt.Println(fmt.Sprintf("initalising server on %s", *addr))
	Services.RunApi(*addr)

}
