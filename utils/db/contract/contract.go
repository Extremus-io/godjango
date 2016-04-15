package contract

import (
	"errors"
	"strings"
	"strconv"
)

// use this struct to create or alter a table.
// Here is the meaning of every field of the struct
// 	Name - name of the table
//	Fields - array of SQLField structs
//	Database - registered name of database in the application. if not provided, defaults to 'default'
type SQLTable struct {
	Name     string
	Fields   []SQLField
	Database string
}
// add a field to table if there is a conflict return error
func (t *SQLTable) AddField(f SQLField) error {
	if !f.Validate() {
		return errors.New("Attempt to add Invalid SQL Field to " + t.Name + " table")
	}
	for i := 0; i < len(t.Fields); i++ {
		if t.Fields[i].Name == f.Name {
			return errors.New("field with name '" + f.Name + "' Already exists")
		}
	}
	t.Fields = append(t.Fields, f)
	return nil
}
// GetField
func (t *SQLTable) GetField(n string) *SQLField {
	for i := 0; i < len(t.Fields); i++ {
		if (t.Fields[i].Name == n) {
			return &t.Fields[i]
		}
	}
	return nil
}
// Helper methods to add new Fields
// TODO: add More Fields here
func (t *SQLTable) NewIntField(name string, def int64) *SQLField {
	t.AddField(SQLField{
		Name: name,
		SQLType: "int",
		Default: strconv.FormatInt(def, 10),
	})
	return t.GetField(name)
}
func (t *SQLTable) NewBoolField(name string, def bool) *SQLField {
	var defaults string
	if def {
		defaults = "1"
	} else {
		defaults = "0"

	}
	t.AddField(SQLField{
		Name: name,
		SQLType: "int",
		Default: defaults,
	})
	return t.GetField(name)
}
func (t *SQLTable) NewStringField(name, def string, max_length int) *SQLField {
	t.AddField(SQLField{
		Name:name,
		SQLType:"varchar(" + strconv.FormatInt(int64(max_length), 10) + ")",
		Default: "'" + def + "'",
	})
	return t.GetField(name)
}
// Generate sql string to create the table
func (t *SQLTable) GetCreateSQL() string{
	sqlstr := "CREATE TABLE " + t.Name + "("
	var fields []string
	for i := 0; i < len(t.Fields); i++ {
		fields = append(fields, t.Fields[i].GetSQL());
	}
	sqlstr += strings.Join(fields, " , ")
	sqlstr += ")"
	return sqlstr
}

type SQLField struct {
	Name    string
	SQLType string
	Unique  bool
	NotNull bool
	Default string
}
// use this method to check if the SQLField is filled enough.
func (f *SQLField) Validate() bool {
	return f.Name != "" && f.SQLType != ""
}
// use this method to get a string to put in create table clause.
// returns nil if couldn't create a string
func (f *SQLField) GetSQL() string {
	if !f.Validate() {
		return ""
	} else {
		output := f.Name
		output += " " + f.SQLType
		if f.NotNull {
			output += "NOT NULL"
		}
		if f.Unique {
			output += " UNIQUE"
		}
		if f.Default != "" {
			output += " DEFAULT " + f.Default
		}
		return output
	}
}
