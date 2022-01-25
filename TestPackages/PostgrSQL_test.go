package TestPackages

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/WalletServices/Utilitys/DbUtils"
	"testing"
)

var (
	send    DbUtils.GreSQL
	receive *DbUtils.GreSQLResult
)

func init() {
	send.Host = "localhost"
	send.Port = 5432
	send.User = "postgres"
	send.Pass = "123456"
	send.Dbname = "Curency"
	send.Driver = "postgres"

}
func newConnection(t *testing.T, s *DbUtils.GreSQL) {
	receive, _ = DbUtils.NewConnection(&send)
}

func TestPgExecuteNonQuery(t *testing.T) {
	var f1, f2, f3, f4, f5, f6 string
	res, _ := DbUtils.NewConnection(&send)
	res.Command = "SELECT * FROM public.users"
	res.PgExecuteNonQuery()
	for res.ResultSet.(*sql.Rows).Next() {
		err := res.ResultSet.(*sql.Rows).Scan(&f1, &f2, &f3, &f4, &f5, &f6)
		fmt.Println(f1, f2, f3, f4, f5, f6)
		fmt.Println(err)
	}
}

func TestNewConnection(t *testing.T) {
	type args struct {
		g *DbUtils.GreSQL
	}
	tests := []struct {
		name string
		args args
		want *DbUtils.GreSQLResult
	}{
		{
			name: "test1",
			args: args{g: &DbUtils.GreSQL{Host: "localhost", Port: 5432, User: "postgres", Pass: "123456", Dbname: "calhounio_demo", Driver: "postgres"}},
			want: &DbUtils.GreSQLResult{
				Command:   "",
				ResultSet: nil,
			},
		},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		got, _ := DbUtils.NewConnection(tt.args.g)
		if got == tt.want {
			t.Errorf("Be ga raftimmmmmmm, Fucking test")
		}

		//})
	}
}
