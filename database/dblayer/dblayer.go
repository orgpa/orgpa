package dblayer

import (
	"orgpa-database-api/database"
	"orgpa-database-api/database/mysql"
)

// DBTYPE represent a type of database (MySQL, MongoDB...)
type DBTYPE string

const (
	// MONGODB defines the mongoDB DBTYPE
	MONGODB DBTYPE = "mongodb"

	// MYSQLDB defines the MySQL DBTYPE
	MYSQLDB DBTYPE = "mysql"
)

// NewDBLayer return a new database handler depending on the database
// you want to use.
// Takes the a few configuration parameters.
func NewDBLayer(dbtype DBTYPE, connection, passwordMySQL, databaseName string) (database.DatabaseHandler, error) {
	switch dbtype {
	case MYSQLDB:
		return mysql.NewMysqlLayer(connection, passwordMySQL, databaseName)
	}
	return nil, nil
}
