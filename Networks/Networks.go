package Networks

import "github.com/pborman/uuid"

type NetworkInterface interface {
	IsValid() (bool, error)
	Load() ([]Network, error)
}
type Network struct {
	ID             uuid.UUID
	NetworkName    string
	NetworkAddress string
	Alias          string
	NetworkType    bool
	Status         bool
}

func (n *Network) IsValid() (bool, error) {
	return true, nil
}

func (n *Network) Active() ([]Network, error) {
	return nil, nil
}
