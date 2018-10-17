package dblayer

import (
	"github.com/frouioui/orgpa/lib/database"
	"github.com/frouioui/orgpa/lib/database/mongo"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
)

// NewDBLayer return a new database handler depending
// on the database we want to use.
func NewDBLayer(dbtype DBTYPE, connection string) (database.DatabaseHandler, error) {
	switch dbtype {
	case MONGODB:
		return mongo.NewMongoLayer(connection)
	}
	return nil, nil
}
