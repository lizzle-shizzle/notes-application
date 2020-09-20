package api

import (
	"fmt"
	"log"
	"time"
)

type Note struct {
	Text string `json:"text"`
}

type NoteRecord struct {
	ID               int       `json:"id"`
	Text             string    `json:"text"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
}

// CreateNote adds a new note to the note table
func CreateNote(c *Client, note Note) (NoteRecord, error) {
	currentTimestamp := time.Now()
	sqlStatement := `
	INSERT INTO note (text, created_timestamp) VALUES ($1, $2)
	RETURNING id`

	id := 0
	err := c.DB.QueryRow(sqlStatement, note.Text, currentTimestamp).Scan(&id)
	if err != nil {
		fmt.Printf("Failed to insert note: %s\n", err)
		return NoteRecord{}, err
	}

	return NoteRecord{
		id,
		note.Text,
		currentTimestamp,
	}, nil
}

// FetchNotes reads all entries in the note table
func FetchNotes(c *Client) ([]NoteRecord, error) {
	notes := []NoteRecord{}
	var id int
	var text string
	var createdTimestamp time.Time
	loc, _ := time.LoadLocation("Africa/Cairo")

	res, err := c.DB.Query("SELECT * FROM note")
	if err != nil {
		log.Println("Failed to fetch notes")
		return notes, err
	}

	defer res.Close()
	for res.Next() {
		err = res.Scan(&id, &text, &createdTimestamp)
		if err != nil {
			log.Println("Failed database scan")
			return notes, err
		}

		note := NoteRecord{
			ID:               id,
			Text:             text,
			CreatedTimestamp: createdTimestamp.In(loc),
		}

		notes = append(notes, note)
	}

	return notes, nil
}
