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
	handler := newServiceHandler(databaseHandler)
	r := mux.NewRouter()

	// Notes routes
	noteSubrouter := r.PathPrefix("/notes").Subrouter()
	noteSubrouter.Methods("GET").Path("").HandlerFunc(handler.getAllNotes)
	noteSubrouter.Methods("POST").Path("").HandlerFunc(handler.addNote)
	noteSubrouter.Methods("GET").Path("/{id}").HandlerFunc(handler.getNoteByID)
	noteSubrouter.Methods("DELETE").Path("/{id}").HandlerFunc(handler.deleteNote)
	noteSubrouter.Methods("PATCH").Path("/{id}").HandlerFunc(handler.patchNote)

	// Todos routes
	todoSubrouter := r.PathPrefix("/todos").Subrouter()
	todoSubrouter.Methods("GET").Path("").HandlerFunc(handler.getAllTodos)
	todoSubrouter.Methods("POST").Path("").HandlerFunc(handler.addTodo)
	todoSubrouter.Methods("GET").Path("/{id}").HandlerFunc(handler.getTodoByID)
	todoSubrouter.Methods("DELETE").Path("/{id}").HandlerFunc(handler.deleteTodo)
	todoSubrouter.Methods("PATCH").Path("/{id}").HandlerFunc(handler.patchTodo)

	srv := http.Server{
		Addr:           config.EndpointAPI,
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
