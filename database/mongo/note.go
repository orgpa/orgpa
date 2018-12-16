package mongo

import (
	"orgpa-database-api/database"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// GetAllNotes return all the notes in the NOTES collection.
func (mgl *MongoDBLayer) GetAllNotes() ([]database.Note, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	notes := []database.Note{}
	err := s.DB(DB).C(NOTES).Find(nil).All(&notes)
	return notes, err
}

// AddNote insert the given note into the database.
func (mgl *MongoDBLayer) AddNote(note database.Note) (database.Note, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	note.ID, _ = strconv.Atoi(string([]byte(bson.NewObjectId())))
	note.LastEdit = time.Now().UTC()
	return note, s.DB(DB).C(NOTES).Insert(note)
}

// GetNoteByID will get the given ID and return the corresponding note.
func (mgl *MongoDBLayer) GetNoteByID(ID []byte) (database.Note, error) {
	s := mgl.getFreshSession()
	defer s.Close()
	note := database.Note{}
	err := s.DB(DB).C(NOTES).FindId(bson.ObjectId(ID)).One(&note)
	return note, err
}

// DeleteNote delete the note corresponding to the given ID.
func (mgl *MongoDBLayer) DeleteNote(ID []byte) error {
	s := mgl.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(NOTES).RemoveId(bson.ObjectId(ID))
}

// PatchNote patch the given note with the given content
func (mgl *MongoDBLayer) PatchNote(ID []byte, content string) error {
	s := mgl.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(NOTES).UpdateId(bson.ObjectId(ID), bson.M{"$set": bson.M{"Content": content}})
}
