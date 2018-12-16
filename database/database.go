package database

// DatabaseHandler interface all the database functions
// useful if ever we want to use a different database.
type DatabaseHandler interface {
	// Notes
	GetAllNotes() ([]Note, error)
	AddNote(note Note) (Note, error)
	GetNoteByID(ID int) (Note, error)
	DeleteNote(ID int) error
	PatchNote(ID int, content string) error

	// Todos
	GetAllTodos() ([]Todo, error)
	AddTodo(todo Todo) (Todo, error)
	GetTodoByID(ID int) (Todo, error)
	DeleteTodo(ID int) error
	PatchTodo(ID int, todo Todo) (Todo, error)
}
