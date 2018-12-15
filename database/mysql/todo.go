package mysql

import "orgpa-database-api/database"

// GetAllTodos returns all the todo in the todos table.
// Return an error if there is any when querying the database.
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
