package Utilitys

import (
	"fmt"
	"testing"
)

var ()

func TestRaiseError(t *testing.T) {
	//RaiseError()
	got := SelectException(10000, RaiseError())

	fmt.Println(got)
}
