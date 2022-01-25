package Consumers

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
	"time"
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

type Consumer struct {
	id             uuid.UUID
	InputParameter *ConsumerInput
	Ticket         string
	validDate      string
	ip             string
	status         bool
}

var (
	db    *DbUtils.GreSQLResult
	crypt *Utilitys.Crypto
	err   *Utilitys.LogInstance
)

func init() {
	db, err = DbUtils.NewConnection(nil)
}
func New() *Consumer {
	crypt, err = Utilitys.NewKey()
	result := new(Consumer)
	result.id = uuid.NewUUID()
	return result
}
func (c *Consumer) GetTicket() *Utilitys.LogInstance {
	db.Command = fmt.Sprintf("select c.\"ID\",c.\"Status\",c.\"IP\",c.\"ValidDate\",COALESCE(d.\"Ticket\",'') from public.\"Counsumers\" c left join public.\"CounsumerDetails\"  d on c.\"ID\"=d.\"CousumerID\" Where c.\"UserName\"='%s' and c.\"Password\"='%s'", c.InputParameter.UserName, c.InputParameter.Password)

	if err := db.PgExecuteNonQuery(); err != nil {
		return Utilitys.Logger("GetTicket", "Db Execute error", c, err)

	}
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&c.id, &c.status, &c.ip, &c.validDate, &c.Ticket); err != nil {
			return Utilitys.Logger("GetTicket", "Db DB SET error", c, err)

		}
	}
	if c.status == true {
		if c.Ticket != "" {
			return nil
		}
		crypt.Text = fmt.Sprintf("%s%s", time.Now().Format("21122006 134225"), Utilitys.RandString(10))
		crypt.Sha256()
		c.Ticket = crypt.Result
		db.Command = fmt.Sprintf("INSERT INTO public.\"CounsumerDetails\"(\"ID\", \"CousumerID\", \"Ticket\")VALUES ('%s', '%s', '%s')", uuid.NewUUID(), c.id, c.Ticket)

		if err := db.PgExecuteNonQuery(); err != nil {
			return Utilitys.Logger("GetTicket", "Db DB SET error", c, err)

		}
	}
	return nil
}

func (c *Consumer) IsValid() *Utilitys.LogInstance {
	db.Command = fmt.Sprintf("select count(*) from public.\"Counsumers\" c left join public.\"CounsumerDetails\"  d on c.\"ID\"=d.\"CousumerID\" Where c.\"UserName\"='%s' and d.\"Ticket\"='%s'", c.InputParameter.UserName, c.InputParameter.InputTicket)

	if err := db.PgExecuteNonQuery(); err != nil {
		return Utilitys.Logger("IsValid", "Db DB SET error", c, err)
	}
	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&counter); err != nil {
			return Utilitys.Logger("IsValid", "Db DB SET error", c, err)
		}
		counter++
	}
	if counter != 1 { //Not found
		return Utilitys.Logger("IsValid", "Db DB SET error", c, err) //It must, will change
	}
	return nil
}
