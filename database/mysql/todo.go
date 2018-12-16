package mysql

import (
	"database/sql"
	"orgpa-database-api/database"
	"strconv"
)

// GetAllTodos returns all the todo in the todos table.
func (msql *MysqlDBLayer) GetAllTodos() ([]database.Todo, error) {
	resp, err := msql.session.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		return []database.Todo{}, err
	}
	defer resp.Close()

	allTodos := make([]database.Todo, 0)
	for resp.Next() {
		var todo database.Todo
		err = resp.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.DueDate, &todo.LastEdit)
		if err != nil {
			return allTodos, err
		}
		allTodos = append(allTodos, todo)
	}
	return allTodos, nil
}

// AddTodo inserts a new Todo struct into the todo table.
func (msql *MysqlDBLayer) AddTodo(todo database.Todo) (database.Todo, error) {
	query, err := msql.session.Prepare("INSERT INTO todos (title,content,due_date) VALUES(?,?,?)")
	if err != nil {
		panic(err)
	}
	defer query.Close()

	result, err := query.Exec(todo.Title, todo.Content, todo.DueDate)
	if err != nil {
		return database.Todo{}, err
	}

	// Get the last inserted ID in order to retreive the new todo and return it.
	newID, err := result.LastInsertId()
	if err != nil {
		return database.Todo{}, err
	}

	newTodo, err := msql.GetTodoByID(int(newID))
	if err != nil {
		return database.Todo{}, err
	}
	return newTodo, nil
}

// GetTodoByID will return the Todo related to the given ID.
func (msql *MysqlDBLayer) GetTodoByID(ID int) (database.Todo, error) {
	resp, err := msql.session.Query("SELECT * FROM todos WHERE id = ?", string(ID))
	if err != nil {
		return database.Todo{}, err
	}

	defer resp.Close()
	var todo database.Todo

	if resp.Next() {
		err = resp.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.LastEdit)
		if err != nil {
			return todo, err
		}
	}
	return todo, nil
}

// DeleteTodo will remove the given todo's ID from the todos table.
func (msql *MysqlDBLayer) DeleteTodo(ID int) error {
	query, err := msql.session.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(string(ID))
	if err != nil {
		return err
	}
	return nil
}

// PatchTodo will modify the given note in the database.
// If the note does not exist, it will be created
func (msql *MysqlDBLayer) PatchTodo(ID int, todo database.Todo) (database.Todo, error) {
	// Create the todo if it already exists.
	if msql.todoExist(ID) == false {
		return msql.AddTodo(todo)
	}

	// Otherwise, update the todo
	query, err := msql.session.Prepare("UPDATE todos SET title=?, content=?, due_date=?, last_edit=? WHERE id=?")
	if err != nil {
		return database.Todo{}, err
	}
	defer query.Close()

	_, err = query.Exec(todo.Title, todo.Content, todo.DueDate, todo.LastEdit, ID)
	if err != nil {
		return database.Todo{}, err
	}
	return todo, nil
}

// todoExist check if the given ID is linked to a Todo in the database.
// If it is linked then we return "true" and if no row is found or
// an error is found then we return "false".
func (msql *MysqlDBLayer) todoExist(ID int) bool {
	row := msql.session.QueryRow("SELECT id FROM todos WHERE id=?", strconv.Itoa(ID))
	err := row.Scan(&ID)
	if err != nil || err == sql.ErrNoRows {
		return false
	}
	return true
}
