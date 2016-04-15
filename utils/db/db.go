package db

import (
	"database/sql"
	"errors"
)

var (
	dblist = DatabaseList{}
)

/* Use DBConn to connect to register a database
	Driver - name of the driver to use
	Name - name used later in program to point to this database conn.
	DBName - name of the database (for SQL)
	Host - Host for mysql or similar
	Username - Username for auth (Optional)
	Password - Password for auth (Optional)
*/
type DBConn struct {
	Driver   string
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}
// use this method to retrive connection string for sql.open method
func (conn *DBConn) GetConStr() string {
	cstr := ""
	if conn.Username != "" {
		cstr += conn.Username + ":"
	}
	if conn.Password != "" {
		cstr += conn.Password
	}
	if cstr != "" {
		cstr += "@"
	}
	if conn.Host != "" {
		cstr += conn.Host
		cstr += ":" + string(conn.Port)

	}
	if cstr != "" {
		cstr += "/"
	}
	cstr += conn.Name
	return cstr
}

// Use this to use multiple databases.
type DatabaseList struct {
	databases map[string]*sql.DB
}

func (dbl *DatabaseList) Get(name string) (*sql.DB, error) {
	db := dbl.databases[name]
	err := db.Ping()
	return db, err
}
func (dbl *DatabaseList) Put(name string, db *sql.DB) error {
	if dbl.databases == nil {
		dbl.databases = make(map[string]*sql.DB, 0)
	}
	dbl.databases[name] = db
	return nil
}

// use this method to register a database to the application
func Register(name string, conn DBConn) error {
	if conn.Driver == "" {
		return errors.New("SQL Driver not given given")
	}
	db, err := Connect(conn.Driver, conn.GetConStr())
	if err != nil {
		return err
	}
	return dblist.Put(name, db)
}