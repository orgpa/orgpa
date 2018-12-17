package database

import (
	"time"
)

// Note model
type Note struct {
	ID       int        `json:"ID"`
	Title    string     `json:"Title"`
	Content  string     `json:"Content"`
	LastEdit *time.Time `json:"LastEdit"`
}

// Todo model
type Todo struct {
	ID       int        `json:"ID"`
	Title    string     `json:"Title"`
	Content  string     `json:"Content"`
	DueDate  *time.Time `json:"DueDate"`
	LastEdit *time.Time `json:"LastEdit"`
}
