package db

import (
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Extremus-io/godjango/utils/db/contract"
	//"database/sql"
)

func TestSQLiteConnect(t *testing.T) {
	err := Register("default", DBConn{
		Name:"TestDB.db",
		Driver:"sqlite3",
	})
	//_, err = db.Exec("CREATE TABLE soab (ID int, name varchar(255))")
	if err != nil {
		t.Error(err);
	}
}

func TestSQLTable_Create(t *testing.T) {
	table := contract.SQLTable{
		Name :"Test2",
	}
	table.NewStringField("Name", "default shit",433).NotNull=true
	table.NewBoolField("Active",false).Unique=true
	table.NewIntField("BullShit",2)
	t.Log(table.GetCreateSQL())
	err := Create(table)
	if err!=nil{
		t.Error(err)
	}
}
