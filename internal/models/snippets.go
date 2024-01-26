package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// a snippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert a new snipper into the db
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	statement := `INSERT INTO snippets (title, content, created, expires) value (?, ?, UTC_timestamp(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(statement,title,content,expires)

	if err != nil{
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}
	return int(id), nil

}

// return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	statement := `SELECt id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() and id = ?`
	row := m.DB.QueryRow(statement, id)

	var s Snippet // variable to store the result

	// roow.Scan() to copy the values from each field in sql.Row to the corresponding field in the Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	
	if err != nil{
		// errors.Is() to compare errors
		if errors.Is(err, sql.ErrNoRows){
			return Snippet{} , ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil

	return Snippet{}, nil
}

// return the 10 most recently created snippets
func (m *SnippetModel) latest() ([]Snippet, error){
	return nil, nil
}