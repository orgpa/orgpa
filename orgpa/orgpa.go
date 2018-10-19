package orgpa

import (
	"net/http"
	"time"

	"orgpa-database-api/configuration"
	"orgpa-database-api/database"

	"github.com/gorilla/mux"
)

// Run run the sover server
func Run(databaseHandler database.DatabaseHandler, config configuration.ServiceConfig) error {
	handler := newEventHandler(databaseHandler)
	r := mux.NewRouter()
	listSubrouter := r.PathPrefix("/list").Subrouter()

	listSubrouter.Methods("GET").Path("").HandlerFunc(handler.getList)
	listSubrouter.Methods("POST").Path("").HandlerFunc(handler.addNote)
	listSubrouter.Methods("GET").Path("/{id}").HandlerFunc(handler.getNoteByID)
	listSubrouter.Methods("DELETE").Path("/{id}").HandlerFunc(handler.deleteNote)

	srv := http.Server{
		Addr:           config.EndpointAPI,
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
