package mysql

import (
	"database/sql"
	"errors"
	"orgpa-database-api/database"
	"orgpa-database-api/message"
	"strconv"
)

// GetAllNotes return all the notes found in the database.
// If there is any error during the query the function will
// return an error.
func (msql *MysqlDBLayer) GetAllNotes() ([]database.Note, error) {
	resp, err := msql.session.Query("SELECT * FROM notes ORDER BY created_at DESC")
	if err != nil {
		return []database.Note{}, err
	}
	defer resp.Close()

	allNotes := make([]database.Note, 0)

	for resp.Next() {
		var note database.Note
		err = resp.Scan(&note.ID, &note.Title, &note.Content, &note.LastEdit, &note.CreatedAt)
		if err != nil {
			return allNotes, err
		}
		allNotes = append(allNotes, note)
	}
	return allNotes, nil
}

// AddNote will insert the given note in the database.
func (msql *MysqlDBLayer) AddNote(note database.Note) (database.Note, error) {
	query, err := msql.session.Prepare("INSERT INTO notes (title,content) VALUES(?,?)")
	if err != nil {
		panic(err)
	}
	defer query.Close()

	result, err := query.Exec(note.Title, note.Content)
	if err != nil {
		return database.Note{}, err
	}
	newID, err := result.LastInsertId()
	if err != nil {
		return database.Note{}, err
	}

	newNote, err := msql.GetNoteByID(int(newID))
	if err != nil {
		return database.Note{}, err
	}

	return newNote, nil
}

// GetNoteByID returns the note corresponding to the given ID.
// Returns an error if there is any when querying the database.
func (msql *MysqlDBLayer) GetNoteByID(ID int) (database.Note, error) {
	resp, err := msql.session.Query("SELECT * FROM notes WHERE id = ?", strconv.Itoa(ID))
	if err != nil {
		return database.Note{}, err
	}

	defer resp.Close()
	var note database.Note

	if resp.Next() {
		err = resp.Scan(&note.ID, &note.Title, &note.Content, &note.LastEdit, &note.CreatedAt)
		if err != nil {
			return note, err
		}
	}
	return note, nil
}

// DeleteNote deletes the given ID into the notes table.
// Returns an error if any.
func (msql *MysqlDBLayer) DeleteNote(ID int) error {
	if msql.noteExist(ID) == false {
		return errors.New(message.NoDataFoundError.Message)
	}

	query, err := msql.session.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		return err
	}

	defer query.Close()
	_, err = query.Exec(strconv.Itoa(ID))
	if err != nil {
		return err
	}
	return nil
}

func (myql *MysqlDBLayer) PatchNote(ID int, content string) error {
	return nil
}

func (msql *MysqlDBLayer) noteExist(ID int) bool {
	row := msql.session.QueryRow("SELECT id FROM notes WHERE id=?", strconv.Itoa(ID))
	err := row.Scan(&ID)
	if err != nil || err == sql.ErrNoRows {
		return false
	}
	return true
}
