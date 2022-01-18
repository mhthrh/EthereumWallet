package Networks

import (
	"github.com/pborman/uuid"
)

type Network struct {
	ID             uuid.UUID
	NetworkName    string
	NetworkAddress string
	Alias          string
	NetworkType    bool
}

func Load() *[]Network {
	return nil
}
