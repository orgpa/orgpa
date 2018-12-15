package mysql

import (
	"orgpa-database-api/database"
)

// GetAllTodos returns all the todo in the todos table.
func (msql *MysqlDBLayer) GetAllTodos() ([]database.Todo, error) {
	resp, err := msql.session.Query("SELECT * FROM todos ORDER BY created_at DESC")
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
