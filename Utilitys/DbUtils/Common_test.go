package DbUtils

import (
	"CurrencyServices/Utilitys"
	"database/sql"
	"reflect"
	"testing"
)

func TestDatabaseProperty_CloseConnection(t *testing.T) {
	type fields struct {
		ConnectionString string
		Driver           string
		Db               *sql.DB
		exception        *[]Utilitys.JsonExceptions
	}
	tests := []struct {
		name   string
		fields fields
		want   *Utilitys.JsonExceptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseProperty{
				ConnectionString: tt.fields.ConnectionString,
				Driver:           tt.fields.Driver,
				Db:               tt.fields.Db,
				exception:        tt.fields.exception,
			}
			if got := d.CloseConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloseConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseProperty_PgExecuteNonQuery(t *testing.T) {
	type fields struct {
		ConnectionString string
		Driver           string
		Db               *sql.DB
		exception        *[]Utilitys.JsonExceptions
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Rows
		want1  *Utilitys.JsonExceptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseProperty{
				ConnectionString: tt.fields.ConnectionString,
				Driver:           tt.fields.Driver,
				Db:               tt.fields.Db,
				exception:        tt.fields.exception,
			}
			got, got1 := d.PgExecuteNonQuery(tt.args.query)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgExecuteNonQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PgExecuteNonQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseProperty_PgGetTimestamp(t *testing.T) {
	type fields struct {
		ConnectionString string
		Driver           string
		Db               *sql.DB
		exception        *[]Utilitys.JsonExceptions
	}
	tests := []struct {
		name   string
		fields fields
		want   *sql.Rows
		want1  *Utilitys.JsonExceptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseProperty{
				ConnectionString: tt.fields.ConnectionString,
				Driver:           tt.fields.Driver,
				Db:               tt.fields.Db,
				exception:        tt.fields.exception,
			}
			got, got1 := d.PgGetTimestamp()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgGetTimestamp() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PgGetTimestamp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseProperty_PgLastInsertId(t *testing.T) {
	type fields struct {
		ConnectionString string
		Driver           string
		Db               *sql.DB
		exception        *[]Utilitys.JsonExceptions
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  *Utilitys.JsonExceptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseProperty{
				ConnectionString: tt.fields.ConnectionString,
				Driver:           tt.fields.Driver,
				Db:               tt.fields.Db,
				exception:        tt.fields.exception,
			}
			got, got1 := d.PgLastInsertId()
			if got != tt.want {
				t.Errorf("PgLastInsertId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PgLastInsertId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewConnection(t *testing.T) {
	type args struct {
		cnn   string
		drive string
		e     *[]Utilitys.JsonExceptions
	}
	tests := []struct {
		name  string
		args  args
		want  *DatabaseProperty
		want1 *Utilitys.JsonExceptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewConnection(tt.args.cnn, tt.args.drive, tt.args.e)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewConnection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
