package Utilitys

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPInterfaces interface {
}

type HTTP struct {
}

func NewHttp() HTTP {
	return HTTP{}

}
func WriteResponse(w http.ResponseWriter, i interface{}) error {
	b, _ := json.Marshal(i)
	fmt.Fprintf(w, "%s", b)
	(w).WriteHeader(http.StatusOK)
	return nil
}

func ReadRequest(r *http.Request, i interface{}) interface{} {
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return nil
	}
	return i
}
