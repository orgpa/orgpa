package database

// DatabaseHandler interface all the database functions
// useful if ever we want to use a different database.
type DatabaseHandler interface {
	GetAllNotes() ([]Notes, error)
	AddNote(note Notes) (Notes, error)
	GetNoteByID(ID []byte) (Notes, error)
	DeleteNote(ID []byte) error
}
