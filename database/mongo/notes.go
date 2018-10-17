package mongo

import (
	"time"

	"github.com/frouioui/orgpa-database-api/database"
	"gopkg.in/mgo.v2/bson"
)

// GetAllNotes return all the notes in the NOTES collection.
func (mgl *MongoDBLayer) GetAllNotes() ([]database.Notes, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	notes := []database.Notes{}
	err := s.DB(DB).C(NOTES).Find(nil).All(&notes)
	return notes, err
}

// AddNote insert the given note into the database.
func (mgl *MongoDBLayer) AddNote(note database.Notes) (database.Notes, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	note.ID = bson.NewObjectId()
	note.LastEdit = time.Now().UTC()
	return note, s.DB(DB).C(NOTES).Insert(note)
}

// GetNoteByID will get the given ID and return the corresponding note.
func (mgl *MongoDBLayer) GetNoteByID(ID []byte) (database.Notes, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	note := database.Notes{}
	err := s.DB(DB).C(NOTES).FindId(bson.ObjectId(ID)).One(&note)
	return note, err
}

// DeleteNote delete the note corresponding to the given ID.
func (mgl *MongoDBLayer) DeleteNote(ID []byte) error {
	s := mgl.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(NOTES).RemoveId(bson.ObjectId(ID))
}
