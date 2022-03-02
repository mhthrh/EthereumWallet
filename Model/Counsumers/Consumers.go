package Consumers

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"github.com/pborman/uuid"
	"time"
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
		if err := db.ResultSet.(*sql.Rows).Scan(&c.id, &c.status, &c.ip, &c.validDate, &c.OutputParameter.Ticket); err != nil {
			return Utilitys.Logger("GetTicket", "Db DB SET error", c, err)
		}
	}
	if c.validDate == "" {
		c.OutputParameter.Result.Code = "24"
		c.OutputParameter.Result.Description = "Cannot find user"
		return nil
	}
	if c.status != true {
		c.OutputParameter.Result.Code = "23"
		c.OutputParameter.Result.Description = "User is disable"
		return nil
	}
	if c.validDate < Utilitys.GetDate("date") {
		c.OutputParameter.Result.Code = "22"
		c.OutputParameter.Result.Description = "Date expired"
		return nil
	}
	if c.OutputParameter.Ticket != "" {
		c.OutputParameter.Result.Code = "00"
		c.OutputParameter.Result.Description = "Successful"
		return nil
	}

	crypt.Text = fmt.Sprintf("%s%s", time.Now().Format("21122006 134225"), Utilitys.RandomString(10))
	if err := crypt.Sha256(); err != nil {
		return Utilitys.Logger("GetTicket", "crypt", c, err)
	}
	c.OutputParameter.Ticket = crypt.Result
	db.Command = fmt.Sprintf("INSERT INTO public.\"CounsumerDetails\"(\"ID\", \"CousumerID\", \"Ticket\")VALUES ('%s', '%s', '%s')", uuid.NewUUID(), c.id, c.OutputParameter.Ticket)

	if err := db.PgExecuteNonQuery(); err != nil {
		return Utilitys.Logger("GetTicket", "Db DB SET error", c, err)
	}
	c.OutputParameter.Result.Code = "00"
	c.OutputParameter.Result.Description = "Successful"
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
