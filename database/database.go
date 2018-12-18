package database

// DatabaseHandler interface all the database functions
type DatabaseHandler interface {
	// Notes
	GetAllNotes() ([]Note, error)
	AddNote(note Note) (Note, error)
	GetNoteByID(ID int) (Note, error)
	DeleteNote(ID int) error
	PatchNote(ID int, note Note) (Note, error)

	// Todos
	GetAllTodos() ([]Todo, error)
	AddTodo(todo Todo) (Todo, error)
	GetTodoByID(ID int) (Todo, error)
	DeleteTodo(ID int) error
	PatchTodo(ID int, todo Todo) (Todo, error)
}
