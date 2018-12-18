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

// Get all the todos in the database and return them
func (sh *serviceHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get all the todos in ID descending order
	todos, err := sh.dbHandler.GetAllTodos()
	if err != nil {
		// TODO:
		// Add a NoDataFound condition.

		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// JSONify and return all the todos found
	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s, "number_of_record": %d}`, string(jsonTodos), len(todos))

}

// Return the todo corresponding to the given ID in the URL.
// If the ID is malformated or no data is found and error
// will be returned over JSON.
// Otherwise, return the todo over JSON.
func (sh *serviceHandler) getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get the URL variables
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Check if the given ID is well formated and
	// transform it into an int.
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// Get te todo by its ID
	todo, err := sh.dbHandler.GetTodoByID(ID)
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

	// Return the JSON of the todo we found
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonTodo))
}

// Add a new todo into the database.
// The todo content will be found in the request body,
// if the JSON found in the body is malformated or missing
// an error will be returned over JSON.
// Otherwise, the new todo will be returned.
func (sh *serviceHandler) addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todo := database.Todo{}

	// Decode and check the request body
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Add the new todo into the database and get the
	// newly added todo.
	newTodo, err := sh.dbHandler.AddTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"%s": false, "error": %s}`, err, message.InternalError.JSON())
		return
	}

	// Return the new todo over JSON
	jsonNewTodo, err := json.Marshal(newTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonNewTodo))
}

// Delete the todo corresponding to the given ID in the database.
// If the ID is missing or malformated an error will be returned,
// otherwise the todo will be deleted and a success message will
// be returned over JSON.
// If no data is found, no ID match, then a NoDataFound error will
// be returned.
func (sh *serviceHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get URL variables
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Convert ID found in URL into a int and verify
	// that it is not malformated.
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Delete the note in the database and check the error
	err = sh.dbHandler.DeleteTodo(ID)
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
	fmt.Fprintf(w, `{"success": true}`)
}

// Patch a todo in the database. The new content will be found
// in the request body. If the body is incorect or missing an
// error will be returned.
func (sh *serviceHandler) patchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todo := database.Todo{}
	vars := mux.Vars(r)

	// Get the URL variable
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	// Get the body and put it in a new todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// check the URL's ID if not incorect
	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	// Patch the todo in database and check for errors
	patchedTodo, err := sh.dbHandler.PatchTodo(ID, todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	// Return the patched todo
	jsonPatchedTodo, err := json.Marshal(patchedTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonPatchedTodo))
}
