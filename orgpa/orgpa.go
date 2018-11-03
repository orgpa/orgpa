package orgpa

import (
	"net/http"
	"orgpa-database-api/configuration"
	"orgpa-database-api/database"
	"time"

	"github.com/gorilla/mux"
)

// Run run the sover server
func Run(databaseHandler database.DatabaseHandler, config configuration.ServiceConfig) error {
	handler := newEventHandler(databaseHandler)
	r := mux.NewRouter()
	listSubrouter := r.PathPrefix("/list").Subrouter()

	listSubrouter.Methods("GET").Path("").HandlerFunc(handler.getAllNotes)
	listSubrouter.Methods("POST").Path("").HandlerFunc(handler.addNote)
	listSubrouter.Methods("GET").Path("/{id}").HandlerFunc(handler.getNoteByID)
	listSubrouter.Methods("DELETE").Path("/{id}").HandlerFunc(handler.deleteNote)
	listSubrouter.Methods("PATCH").Path("/{id}").HandlerFunc(handler.patchNote)

	srv := http.Server{
		Addr:           config.EndpointAPI,
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
