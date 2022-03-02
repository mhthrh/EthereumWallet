package Utilitys

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ServerQuery() []byte {
	conn, err := net.Dial("tcp", ":85")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return nil
	}
	return []byte(str)
}
