package db

import "github.com/Extremus-io/godjango/utils/db/contract"

// uses the data in the struct and executes a create table command on database.
// if no database is mentioned, uses 'default' database.
func  Create(t *contract.SQLTable) error {
	sqlstr := t.GetCreateSQL()
	if t.Database == "" {
		t.Database = "default"
	}
	db, err := dblist.Get(t.Database)
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlstr)
	return err
}