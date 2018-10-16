package orgpa

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../lib/database"
)

type eventServiceHandler struct {
	dbHandler database.DatabaseHandler
}

func newEventHandler(databaseHandler database.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbHandler: databaseHandler,
	}
}

func (eh *eventServiceHandler) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func (eh *eventServiceHandler) getList(w http.ResponseWriter, r *http.Request) {
	notes, err := eh.dbHandler.GetAllNotes()
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: %s}", err)
		return
	}
	err = json.NewEncoder(w).Encode(&notes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) addNote(w http.ResponseWriter, r *http.Request) {
	note := database.Notes{}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to get the data %s}", err)
		return
	}
	note, err = eh.dbHandler.AddNote(note)
	fmt.Fprintln(w, note)
}
