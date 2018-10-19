package orgpa

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"orgpa-database-api/database"

	"github.com/gorilla/mux"
)

type eventServiceHandler struct {
	dbHandler database.DatabaseHandler
}

func newEventHandler(databaseHandler database.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbHandler: databaseHandler,
	}
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
		fmt.Fprintf(w, "{error: Error occured while trying encode notes to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) getNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No ID found}`)
		return
	}

	ID, err := hex.DecodeString(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: %s}", err)
		return
	}

	note, err := eh.dbHandler.GetNoteByID(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to get the data %s}", err)
		return
	}
	err = json.NewEncoder(w).Encode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode the note to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	note := database.Notes{}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to get the data %s}", err)
		return
	}
	note, err = eh.dbHandler.AddNote(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to insert the data into the database %s", err)
		return
	}
	err = json.NewEncoder(w).Encode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode the note to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No ID found}`)
		return
	}

	ID, err := hex.DecodeString(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: %s}", err)
		return
	}
	err = eh.dbHandler.DeleteNote(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to delete the note %s}", err)
		return
	}
}
