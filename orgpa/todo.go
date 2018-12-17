package orgpa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orgpa-database-api/database"
	"strconv"

	"github.com/gorilla/mux"
)

func (sh *serviceHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todos, err := sh.dbHandler.GetAllTodos()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	err = json.NewEncoder(w).Encode(&todos)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode notes to JSON %s"}`, err)
	}
}

func (sh *serviceHandler) getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No ID found"}`)
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while decoding the ID: %s"}`, err)
		return
	}

	note, err := sh.dbHandler.GetNoteByID(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to get the data %s"}`, err)
		return
	}
	err = json.NewEncoder(w).Encode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode the note to JSON %s"}`, err)
	}
}

func (sh *serviceHandler) addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	todo := database.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to get the data %s"}`, err)
		return
	}

	newTodo, err := sh.dbHandler.AddTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to insert the data into the database %s"}`, err)
		return
	}

	err = json.NewEncoder(w).Encode(&newTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode the note to JSON %s"}`, err)
	}
}

func (sh *serviceHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	vars := mux.Vars(r)
	varID, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No ID found"}`)
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while decoding the ID: %s"}`, err)
		return
	}

	err = sh.dbHandler.DeleteTodo(ID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to delete the note %s"}`, err)
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
		fmt.Fprint(w, `{"error": "missing information"}`)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	ID, err := strconv.Atoi(varID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while decoding the ID: %s"}`, err)
		return
	}

	patchedTodo, err := sh.dbHandler.PatchTodo(ID, todo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to patching the note %s"}`, err)
		return
	}

	err = json.NewEncoder(w).Encode(&patchedTodo)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode the note to JSON %s"}`, err)
	}

}
