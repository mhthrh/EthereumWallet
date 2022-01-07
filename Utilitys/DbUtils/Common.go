package DbUtils

import (
	"CurrencyServices/Utilitys"
	"database/sql"
	_ "github.com/lib/pq"
)

type DatabaseInterfaces interface {
	NewConnection(string, string, *[]Utilitys.JsonExceptions) (*DatabaseProperty, *Utilitys.JsonExceptions)
	CloseConnection() *Utilitys.JsonExceptions
	PgExecuteNonQuery(string) (*sql.Rows, *Utilitys.JsonExceptions)
	PgGetTimestamp() (*sql.Rows, *Utilitys.JsonExceptions)
	PgLastInsertId() (int, *Utilitys.JsonExceptions)
}
type DatabaseProperty struct {
	ConnectionString string
	Driver           string
	Db               *sql.DB
	exception        *[]Utilitys.JsonExceptions
}

func NewConnection(cnn, drive string, e *[]Utilitys.JsonExceptions) (*DatabaseProperty, *Utilitys.JsonExceptions) {
	d, err := sql.Open(drive, cnn)
	if err != nil {
		return &DatabaseProperty{Driver: drive, ConnectionString: cnn, Db: d, exception: e}, Utilitys.SelectException(10005, e)
	}
	return &DatabaseProperty{Driver: drive, ConnectionString: cnn, Db: d, exception: e}, Utilitys.SelectException(0, e)
}

func (d *DatabaseProperty) CloseConnection() *Utilitys.JsonExceptions {
	if err := d.Db.Close(); err != nil {
		return Utilitys.SelectException(10006, d.exception)
	}
	return Utilitys.SelectException(0, d.exception)
}
