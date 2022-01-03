package Currencys

import "github.com/pborman/uuid"

type CurrencyInterface interface {
	Add() (Currency, error)
	Active() (bool, error)
	DeActive() (bool, error)
	Load() ([]Currency, error)
}
type Currency struct {
	ID           uuid.UUID
	Name         string
	Appreciation string
	Status       bool
}

func New() CurrencyInterface {
	r := new(Currency)
	r.ID = uuid.NewRandom()
	return r
}
func (c *Currency) Add() (Currency, error) {
	return *c, nil
}

func (c *Currency) Active() (bool, error) {
	return true, nil
}

func (c *Currency) DeActive() (bool, error) {
	return true, nil
}

func (c *Currency) Load() ([]Currency, error) {
	return nil, nil
}
