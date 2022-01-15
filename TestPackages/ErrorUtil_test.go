package TestPackages

import (
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"testing"
)

var ()

func TestRaiseError(t *testing.T) {
	//RaiseError()

	got := Utilitys.SelectException(10000, Utilitys.RaiseError())

	fmt.Println(got)
}
