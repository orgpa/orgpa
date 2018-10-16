package database

// DatabaseHandler interface all the database functions
// useful if ever we want to use a different database.
type DatabaseHandler interface {
	GetAllNotes() ([]Notes, error)
}
