package DbUtils

import (
	"github.com/mhthrh/WalletServices/Utilitys"
)

func (d *GreSQLResult) PgExecuteNonQuery() *Utilitys.LogInstance {
	var err error
	d.ResultSet, err = d.db.Query(d.Command)
	if err != nil {
		return Utilitys.Logger("PgExecuteNonQuery", "SQL execution Error", d, err)
	}
	return nil
}

func (d *GreSQLResult) PgLastInsertId() *Utilitys.LogInstance {
	err := d.db.QueryRow(d.Command).Scan(&d.ResultSet)
	if err != nil {
		return Utilitys.Logger("PgExecuteNonQuery", "SQL execution Error", d, err)
	}
	return nil
}
