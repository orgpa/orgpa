package orgpa

import (
	"net/http"
	"time"

	"../lib/database"
	"github.com/gorilla/mux"
)

// Run run the sover server
func Run(databaseHandler database.DatabaseHandler) error {
	handler := newEventHandler(databaseHandler)
	r := mux.NewRouter()
	listSubrouter := r.PathPrefix("/list").Subrouter()

	listSubrouter.Methods("GET").Path("").HandlerFunc(handler.getList)
	listSubrouter.Methods("POST").Path("").HandlerFunc(handler.addNote)
	listSubrouter.Methods("GET").Path("/{id}").HandlerFunc(handler.getNoteByID)
	listSubrouter.Methods("DELETE").Path("/{id}").HandlerFunc(handler.deleteNote)

	srv := http.Server{
		Addr:           "localhost:8000",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
