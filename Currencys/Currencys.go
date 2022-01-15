package Currencys

import "github.com/pborman/uuid"

type CurrencyInterface interface {
	Add()
	Active()
	DeActive()
	Load() *[]Currency
}
type Currency struct {
	ID           uuid.UUID
	Name         string
	Appreciation string
	Status       bool
}

func New() Currency {
	r := new(Currency)
	r.ID = uuid.NewRandom()
	return *r
}
