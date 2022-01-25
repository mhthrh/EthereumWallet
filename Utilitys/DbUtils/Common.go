package DbUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mhthrh/WalletServices/Utilitys"
)

type DatabaseInterfaces interface {
	NewConnection(*GreSQL)
	CloseConnection()
	PgExecuteNonQuery()
	PgLastInsertId()
}

type GreSQLResult struct {
	db        *sql.DB
	Command   string
	ResultSet interface{}
}
type GreSQL struct {
	Host   string
	Port   int32
	User   string
	Pass   string
	Dbname string
	Driver string
}

func NewConnection(g *GreSQL) (*GreSQLResult, *Utilitys.LogInstance) {
	r := new(GreSQLResult)
	if g == nil {
		g = &GreSQL{
			Host:   "localhost",
			Port:   5432,
			User:   "postgres",
			Pass:   "123456",
			Dbname: "Curency",
			Driver: "postgres",
		}
	}
	db, err := sql.Open(g.Driver, fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		g.Host, g.Port, g.User, g.Pass, g.Dbname))
	if err != nil {
		return nil, Utilitys.Logger("NewConnection", "Db Connection error", g, err)
	}
	r.db = db
	r.Command = ""
	r.ResultSet = nil
	return r, nil
}

func (d *GreSQLResult) CloseConnection() *Utilitys.LogInstance {

	if err := d.db.Close(); err != nil {
		return Utilitys.Logger("NewConnection", "Db Connection error", d, err)
	}
	return nil
}
