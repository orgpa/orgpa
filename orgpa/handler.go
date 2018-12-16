package orgpa

import (
	"orgpa-database-api/database"
)

type serviceHandler struct {
	dbHandler database.DatabaseHandler
}

func newServiceHandler(databaseHandler database.DatabaseHandler) *serviceHandler {
	return &serviceHandler{
		dbHandler: databaseHandler,
	}
}
