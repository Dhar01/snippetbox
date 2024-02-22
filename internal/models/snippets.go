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

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// executing statement, returns a sql.result type
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get the ID of newly inserted record in the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// returns specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// QueryRow() method on connection pool will execute SQL statement.
	// row := m.DB.QueryRow(stmt, id)

	// initialize a pointer to a new zeroed Snippet struct
	s := &Snippet{}

	// err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, ErrNoRecord
	// 	} else {
	// 		return nil, err
	// 	}
	// }

	// rewriting above code
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// if everything OK, return snippet object
	return s, nil

}

// return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
