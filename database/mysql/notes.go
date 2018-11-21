package mysql

import (
	"orgpa-database-api/database"
	"strconv"
)

// GetAllNotes return all the notes found in the database.
// If there is any error during the query the function will
// return an error.
func (msql *MysqlDBLayer) GetAllNotes() ([]database.Notes, error) {
	resp, err := msql.session.Query("SELECT * FROM notes ORDER BY created_at DESC")
	if err != nil {
		return []database.Notes{}, err
	}

	defer resp.Close()
	allNotes := make([]database.Notes, 0)

	for resp.Next() {
		var note database.Notes
		err = resp.Scan(&note.ID, &note.Title, &note.Content, &note.LastEdit)
		if err != nil {
			return allNotes, err
		}
		allNotes = append(allNotes, note)
	}
	return allNotes, nil
}

func (msql *MysqlDBLayer) AddNote(note database.Notes) (database.Notes, error) {

	query, err := msql.session.Prepare("INSERT INTO notes (title,content) VALUES(?,?)")
	if err != nil {
		panic(err)
	}
	defer query.Close()

	result, err := query.Exec(note.Title, note.Content)
	if err != nil {
		return database.Notes{}, err
	}
	newID, err := result.LastInsertId()
	if err != nil {
		return database.Notes{}, err
	}

	byteID := []byte(strconv.Itoa(int(newID)))
	newNote, err := msql.GetNoteByID(byteID)
	if err != nil {
		return database.Notes{}, err
	}

	return newNote, nil
}

// GetNoteByID returns the note corresponding to the given ID.
// Returns an error if there is any when querying the database.
func (msql *MysqlDBLayer) GetNoteByID(ID []byte) (database.Notes, error) {
	resp, err := msql.session.Query("SELECT * FROM notes WHERE id = ?", string(ID))
	if err != nil {
		return database.Notes{}, err
	}

	defer resp.Close()
	var note database.Notes

	if resp.Next() {
		err = resp.Scan(&note.ID, &note.Title, &note.Content, &note.LastEdit)
		if err != nil {
			return note, err
		}
	}
	return note, nil
}

// DeleteNote deletes the given ID into the notes table.
// Returns an error if any.
func (msql *MysqlDBLayer) DeleteNote(ID []byte) error {
	query, err := msql.session.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		return err
	}

	defer query.Close()
	_, err = query.Exec(string(ID))
	if err != nil {
		return err
	}
	return nil
}

func (myql *MysqlDBLayer) PatchNote(ID []byte, content string) error {
	return nil
}
