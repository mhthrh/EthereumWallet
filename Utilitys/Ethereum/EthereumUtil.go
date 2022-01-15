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
	exceptions *[]Utilitys.Exceptions
	Status     *Utilitys.Exceptions
}

func New(s *[]Utilitys.Exceptions) *Ether {
	e := new(Ether)
	e.exceptions = s
	return e
}

func (e *Ether) GetPrivate() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}
	k := Utilitys.NewKey()
	if k.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}

	privateKeyByte := crypto.FromECDSA(privateKey)
	k.Text = string(privateKeyByte[:2])
	k.Encrypt()
	if k.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}
	e.BytString = k.Result
	k.Text = string(privateKeyByte[2:])
	k.Encrypt()
	if k.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}
	e.PrivateKey = k.Result
}
