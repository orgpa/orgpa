package mongo

import (
	"../../database"
	mgo "gopkg.in/mgo.v2"
)

const (
	// DB : database name
	DB = "sover"

	// NOTES : name of the collection containing the notes
	NOTES = "notes"
)

// MongoDBLayer interfacing with DatabaseHandler
type MongoDBLayer struct {
	session *mgo.Session
}

// NewMongoLayer return a new connection to a mongoDB
// database. Return a non nil error if can't connect.
func NewMongoLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		session: s,
	}, nil
}

// GetAllNotes return all the notes in the NOTES collection.
func (mgl *MongoDBLayer) GetAllNotes() ([]database.Notes, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	notes := []database.Notes{}
	err := s.DB(DB).C(NOTES).Find(nil).All(&notes)
	return notes, err
}

// Create a fresh new session and return it.
func (mgl *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgl.session.Copy()
}
