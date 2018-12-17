package orgpa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orgpa-database-api/database"
	"orgpa-database-api/orgpa/message"
	"strconv"

	"github.com/gorilla/mux"
)

func (sh *serviceHandler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := sh.dbHandler.GetAllNotes()
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	jsonNotes, err := json.Marshal(notes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s, "number_of_record": %d}`, string(jsonNotes), len(notes))
}

func (sh *serviceHandler) getNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	note, err := sh.dbHandler.GetNoteByID(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	jsonNote, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s`, jsonNote)
}

func (sh *serviceHandler) addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	note := database.Note{}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}
	note, err = sh.dbHandler.AddNote(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	jsonNote, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonNote))
}

func (sh *serviceHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	err = sh.dbHandler.DeleteNote(ID)
	if err.Error() == message.NoDataFoundError.Message {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.NoDataFoundError.JSON())
		return
	}
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true}`)
}

func (sh *serviceHandler) patchNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	note := database.Note{}
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&note)
	if !ok || err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	err = sh.dbHandler.PatchNote(ID, note.Content)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true}`)
}
