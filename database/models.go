package database

import (
	"time"
)

// Note represent one row of the "notes" table in the database
type Note struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	LastEdit  *time.Time `json:"last_edit"`
	CreatedAt *time.Time `json:"created_at"`
}

// Todo represent one row of the "todos" table in the database.
type Todo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	DueDate   *time.Time `json:"due_date"`
	LastEdit  *time.Time `json:"last_edit"`
	CreatedAt *time.Time `json:"created_at"`
}
