package mysql

import (
	"database/sql"

	// MySQL drivers
	_ "github.com/go-sql-driver/mysql"
)

// MysqlDBLayer is the MySQL database layer.
// It contains the connection to the database and
// methodes to close and open it.
type MysqlDBLayer struct {
	session *sql.DB
}

// NewMysqlLayer returns a new connection to MySQL
func NewMysqlLayer(connection, dbPassword, dbName string) (*MysqlDBLayer, error) {
	dbDriver := "mysql"
	dbUser := "root"

	dbAdress := dbUser + ":" + dbPassword + "@tcp(" + connection + ")/" + dbName + "?parseTime=true"

	session, err := sql.Open(dbDriver, dbAdress)
	return &MysqlDBLayer{session}, err
}

// CloseConnection with the mysql database.
func (msql *MysqlDBLayer) CloseConnection() error {
	return msql.session.Close()
}
