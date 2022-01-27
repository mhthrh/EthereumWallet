package Consumers

import (
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
)

type ConsumerInterface interface {
	GetTicket() *Utilitys.LogInstance
	IsValid() *Utilitys.LogInstance
}
type ConsumerInput struct {
	UserName    string
	Password    string
	InputTicket string
	IsTest      bool
}
type ConsumerOutput struct {
	Ticket string
	Result *Utilitys.ResultSet
}

type Consumer struct {
	id              uuid.UUID
	InputParameter  *ConsumerInput
	OutputParameter *ConsumerOutput
	validDate       string
	ip              string
	status          bool
}

var (
	db    *DbUtils.GreSQLResult
	crypt *Utilitys.Crypto
	err   *Utilitys.LogInstance
)
