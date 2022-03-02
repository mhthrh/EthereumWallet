package Ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mhthrh/WalletServices/Model/Accounts"
	"github.com/mhthrh/WalletServices/Model/Networks"
	"github.com/mhthrh/WalletServices/Utilitys"
	"math/big"
	"strconv"
)

type EtherTransInterface interface {
	SendTrans()
}

type EtherTransaction struct {
	To       string
	Amount   string
	Private  *ecdsa.PrivateKey
	GasLimit uint64
	Account  *Accounts.Account
	Network  *Networks.Network
}

func NewTransaction() (*EtherTransaction, *Utilitys.LogInstance) {
	e := new(EtherTransaction)
	c, err := Utilitys.NewKey()
	if err != nil {
		return nil, Utilitys.Logger("NewTransaction", "Db Connection error", e, err)
	}
	c.Text = e.Account.PrivateKey
	c.Decrypt()
	privy := c.Result
	c.Text = e.Account.Byt
	c.Decrypt()
	byt := c.Result
	prv, _ := crypto.ToECDSA(append([]byte(byt), []byte(privy)...))
	e.Network.NetworkAddress = "https://ropsten.infura.io/v3/0be4c1b7d14c418d9a85ff78a14674ed"
	e.Private = prv
	e.GasLimit = uint64(21000)
	return e, nil
}

func (t *EtherTransaction) SendTransEther() *Utilitys.LogInstance {
	client, err := ethclient.Dial(t.Network.NetworkAddress)
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	privateKey := t.Private
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	n, _ := strconv.ParseInt(t.Amount, 10, 64)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(t.To), big.NewInt(n*1000000000000000000), uint64(30000), gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return Utilitys.Logger("SendTransEther", " EtherTransaction error", t, err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return nil
}
