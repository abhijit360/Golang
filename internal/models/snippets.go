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
}

// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error){
	statement := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
	
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}

	// ensure that the sql.Rows result set is always properly closed before the Latest() method returns
	// statement should come *after* you check for an error from the Query() method
	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		// create a pointer to a new zeoed Snippet struct
		var s Snippet

		// use rpws.Scan() to copy the values from each field in the row to the 
		err = rows.Scan(&s.ID,&s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		
		//apend it to the slice of snippets
		snippets = append(snippets, s)
	}

	return snippets, nil
}