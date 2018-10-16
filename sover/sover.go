package sover

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

	r.Methods("GET").Path("/").HandlerFunc(handler.homePage)
	listSubrouter.Methods("GET").Path("").HandlerFunc(handler.getList)

	srv := http.Server{
		Addr:           "localhost:8000",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
