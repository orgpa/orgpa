package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Notes model
type Notes struct {
	ID       bson.ObjectId `bson:"_id"`
	Title    string        `bson:"Title"`
	Content  string        `bson:"Content"`
	LastEdit time.Time     `bson:"LastEdit"`
}
