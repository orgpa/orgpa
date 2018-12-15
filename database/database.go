package database

// DatabaseHandler interface all the database functions
// useful if ever we want to use a different database.
type DatabaseHandler interface {
	GetAllNotes() ([]Note, error)
	AddNote(note Note) (Note, error)
	GetNoteByID(ID []byte) (Note, error)
	DeleteNote(ID []byte) error
	PatchNote(ID []byte, content string) error
}
