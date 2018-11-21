package database

import (
	"time"
)

// Notes model
type Notes struct {
	ID       int       `json:"ID"`
	Title    string    `json:"Title"`
	Content  string    `json:"Content"`
	LastEdit time.Time `json:"LastEdit"`
}
