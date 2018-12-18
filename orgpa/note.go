package orgpa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orgpa-database-api/database"
	"orgpa-database-api/message"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all the notes in the database and return them
func (sh *serviceHandler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get all the notes
	notes, err := sh.dbHandler.GetAllNotes()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// JSONify the found notes and return them
	jsonNotes, err := json.Marshal(notes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s, "number_of_record": %d}`, string(jsonNotes), len(notes))
}

// Return the note corresponding to the given ID in the URL.
// If the ID is malformated or no data is found and error
// will be returned over JSON.
// Otherwise, return the note over JSON.
func (sh *serviceHandler) getNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get the variable in the URL
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Check the ID
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Get the note in the database
	note, err := sh.dbHandler.GetNoteByID(ID)
	if err != nil {
		// Check if the error from the databaseHandler is a
		// NoDataFoundError, in this case we return the corresponding
		// JSON error.
		if message.IsNoDataErr(err) {
			w.WriteHeader(400)
			fmt.Fprintf(w, `{"success": false, "error": %s}`, message.NoDataFoundError.JSON())
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		}
		return
	}

	// Return a JSON of the found note
	jsonNote, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s`, jsonNote)
}

// Add a new note into the database.
// The note content will be found in the request body,
// if the JSON found in the body is malformated or missing
// an error will be returned over JSON.
// Otherwise, the new note will be returned.
func (sh *serviceHandler) addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	note := database.Note{}

	// Decode and get the note from the request body
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Add the note in the database
	note, err = sh.dbHandler.AddNote(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// Return the new note in JSON
	jsonNote, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonNote))
}

// Delete the note corresponding to the given ID in the database.
// If the ID is missing or malformated an error will be returned,
// otherwise the note will be deleted and a success message will
// be returned over JSON.
// If no data is found, no ID match, then a NoDataFound error will
// be returned.
func (sh *serviceHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get ID from the URL
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Check the ID is not malformated
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Delete the note and check for errors
	err = sh.dbHandler.DeleteNote(ID)
	if err != nil {
		if message.IsNoDataErr(err) {
			// If non nil error and error is a NoDataFound error
			w.WriteHeader(400)
			fmt.Fprintf(w, `{"success": false, "error": %s}`, message.NoDataFoundError.JSON())
		} else {
			// If it is an internal error
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		}
		return
	}
	fmt.Fprint(w, `{"success": true}`)
}

// Patch a note in the database. The new content will be found
// in the request body. If the body is incorect or missing an
// error will be returned.
func (sh *serviceHandler) patchNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	note := database.Note{}

	// Get ID from the URL
	vars := mux.Vars(r)
	varID, ok := vars["id"]

	// Decode the note inside the body
	err := json.NewDecoder(r.Body).Decode(&note)
	if !ok || err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Check the ID is correct and convert into int
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Patch the note in the database
	patchedNote, err := sh.dbHandler.PatchNote(ID, note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// Return the patched note
	jsonPatchedNote, err := json.Marshal(patchedNote)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonPatchedNote))
}
