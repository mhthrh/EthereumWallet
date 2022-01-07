package DbUtils

import (
	"CurrencyServices/Utilitys"
	"database/sql"
)

func (d *DatabaseProperty) PgExecuteNonQuery(query string) (*sql.Rows, *Utilitys.JsonExceptions) {
	r, err := d.Db.Query(query)
	if err != nil {
		return nil, Utilitys.SelectException(10007, d.exception)
	}
	return r, Utilitys.SelectException(0, d.exception)
}
func (d *DatabaseProperty) PgGetTimestamp() (*sql.Rows, *Utilitys.JsonExceptions) {
	r, err := d.Db.Query("SELECT CURRENT_TIMESTAMP")
	if err != nil {
		return nil, Utilitys.SelectException(10007, d.exception)
	}
	return r, Utilitys.SelectException(0, d.exception)
}

func (d *DatabaseProperty) PgLastInsertId() (int, *Utilitys.JsonExceptions) {
	var id int
	err := d.Db.QueryRow("SELECT CURRENT_TIMESTAMP").Scan(&id)
	if err != nil {
		return 0, Utilitys.SelectException(10007, d.exception)
	}
	return id, Utilitys.SelectException(0, d.exception)
}
