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

func (sh *serviceHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	todos, err := sh.dbHandler.GetAllTodos()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s, "number_of_record": %d}`, string(jsonTodos), len(todos))

}

func (sh *serviceHandler) getTodoByID(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	todo, err := sh.dbHandler.GetNoteByID(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonTodo))
}

func (sh *serviceHandler) addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todo := database.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	newTodo, err := sh.dbHandler.AddTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"%s": false, "error": %s}`, err, message.InternalError.JSON())
		return
	}

	jsonNewTodo, err := json.Marshal(newTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonNewTodo))
}

func (sh *serviceHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
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

	err = sh.dbHandler.DeleteTodo(ID)
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

func (sh *serviceHandler) patchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todo := database.Todo{}
	vars := mux.Vars(r)

	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MissingInformationError.JSON())
		return
	}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.MalformatedDataError.JSON())
		return
	}

	patchedTodo, err := sh.dbHandler.PatchTodo(ID, todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}

	jsonPatchedTodo, err := json.Marshal(patchedTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"success": false, "error": %s}`, message.InternalError.JSON())
		return
	}
	fmt.Fprintf(w, `{"success": true, "data": %s}`, string(jsonPatchedTodo))
}
