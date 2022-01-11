package TestPackages

import (
	"CurrencyServices/Utilitys"
	"fmt"
	"testing"
)

var ()

func TestRaiseError(t *testing.T) {
	//RaiseError()

	got := Utilitys.SelectException(10000, Utilitys.RaiseError())

	fmt.Println(got)
}
