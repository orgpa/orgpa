package dblayer

import (
	"orgpa-database-api/database"
	"orgpa-database-api/database/mongo"
	"orgpa-database-api/database/mysql"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	MYSQLDB DBTYPE = "mysql"
)

// NewDBLayer return a new database handler depending
// on the database we want to use.
func NewDBLayer(dbtype DBTYPE, connection, passwordMySQL, databaseName string) (database.DatabaseHandler, error) {
	switch dbtype {
	case MONGODB:
		return mongo.NewMongoLayer(connection)
	case MYSQLDB:
		return mysql.NewMysqlLayer(connection, passwordMySQL, databaseName)
	}
	return nil, nil
}
