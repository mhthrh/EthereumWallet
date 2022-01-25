package Ethereum

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mhthrh/WalletServices/Utilitys"
)

type EtherInterface interface {
	GetPrivate()
}

type Ether struct {
	PrivateKey string
	BytString  string
}

func New() (*Ether, *Utilitys.LogInstance) {
	return new(Ether), nil
}

func (e *Ether) GetPrivate() *Utilitys.LogInstance {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Utilitys.Logger("GetPrivate", "Generate key bega", e, err)
	}
	k, err2 := Utilitys.NewKey()
	if err2 != nil {
		return Utilitys.Logger("GetPrivate", "Generate key bega", e, err2)
	}

	privateKeyByte := crypto.FromECDSA(privateKey)
	k.Text = string(privateKeyByte[:2])

	if err := k.Encrypt(); err != nil {
		return Utilitys.Logger("GetPrivate", "Generate key bega", e, err)
	}
	e.BytString = k.Result
	k.Text = string(privateKeyByte[2:])

	if err := k.Encrypt(); err != nil {
		return Utilitys.Logger("GetPrivate", "Generate key bega", e, err)
	}
	e.PrivateKey = k.Result
	return nil
}
