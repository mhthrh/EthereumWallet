package Currencys

import "github.com/pborman/uuid"

type Currency struct {
	ID           uuid.UUID
	Name         string
	Appreciation string
	Status       bool
}
