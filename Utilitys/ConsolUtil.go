package Utilitys

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}
func ReadLine() string {

	isDebug := func() bool {
		// gops executable must be in the path. See https://github.com/google/gops
		gopsOut, err := exec.Command("gops", strconv.Itoa(os.Getppid())).Output()
		if err == nil && strings.Contains(string(gopsOut), "\\dlv.exe") {
			// our parent process is (probably) the Delve debugger
			return true
		}
		return false
	}()
	reader := bufio.NewReader(os.Stdin)
	if isDebug {
		fmt.Println("ddd")
		l, err := net.Listen("tcp", "127.0.0.1:1234")
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer l.Close()

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return ""
		}

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Print("-> ", string(netData))

	}

	text, _ := reader.ReadString('\r')
	text = strings.Replace(text, "\r", "", -1)
	return text
}
