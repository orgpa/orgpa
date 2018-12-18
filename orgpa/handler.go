package orgpa

import (
	"orgpa-database-api/database"
)

type serviceHandler struct {
	dbHandler database.DatabaseHandler
}

// Create a new serviceHandler with the given database handler.
func newServiceHandler(databaseHandler database.DatabaseHandler) *serviceHandler {
	return &serviceHandler{
		dbHandler: databaseHandler,
	}
}
