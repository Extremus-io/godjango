package db

import "database/sql"

// Handles Connection Failing events and enables to add middleware
// TODO: implement reconnect attempts
func Connect(driver string, dbName string) (*sql.DB, error){
	return sql.Open(driver, dbName)
}