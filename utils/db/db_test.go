package db

import (
	"testing"
	_ "github.com/mattn/go-sqlite3"
//"database/sql"
)


func TestSQLiteConnect(t *testing.T) {
	err := Register("default",DBConn{
		Name:"TestDB.db",
	})
	//_, err = db.Exec("CREATE TABLE soab (ID int, name varchar(255))")
	if err != nil {
		t.Error(err);
	}
}
