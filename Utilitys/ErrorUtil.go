package Utilitys

import (
	"fmt"
	"time"
)

type CustomError struct {
	when  time.Time
	Where string
	Dic   ErrorDictionary
}
type ErrorDictionary struct {
	Priority   int
	Code       int
	What       string
	Suggestion string
}

func GetErrorCode(C int) CustomError {

}

func NewError() CustomError {
	err := new(CustomError)
	err.when = time.Now()
	return *err
}

func (e *CustomError) RaiseError() string {
	return fmt.Sprintf("%v: %s: %d: %d: %s: %s", e.when, e.Dic.What, e.Priority, e.Code, e.Suggestion, e.Where)
}
