package event

import (
	"database/sql"
	"fmt"

	"github.com/jonathanmeij/go-reservation/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetEventByID(id int) (*types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanRowsIntoEvent(rows)
	}

	return nil, fmt.Errorf("account not found")
}

func (s *Store) GetEvents() ([]*types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	events := make([]*types.Event, 0)
	for rows.Next() {
		e, err := scanRowsIntoEvent(rows)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func (s *Store) CreateEvent(event types.Event) error {
	_, err := s.db.Exec(`INSERT INTO events 
	(title, description, image_url, date, created_at) 
	VALUES ($1, $2, $3, $4, $5)`, event.Title, event.Description, event.ImageUrl, event.Date, event.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteEvent(id int) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateEvent(event types.Event) error {
	oldEvent, err := s.GetEventByID(event.ID)

	if err != nil {
		return err
	}

	if event.Title != "" {
		oldEvent.Title = event.Title
	}
	if event.Description != "" {
		oldEvent.Description = event.Description
	}
	if event.ImageUrl != "" {
		oldEvent.ImageUrl = event.ImageUrl
	}
	if !event.Date.IsZero() {
		oldEvent.Date = event.Date
	}

	_, err = s.db.Exec(`UPDATE events SET 
	title = $1, description = $2, image_url = $3, date = $4, created_at = $5 
	WHERE id = $6`, oldEvent.Title, oldEvent.Description, oldEvent.ImageUrl, oldEvent.Date, oldEvent.CreatedAt, oldEvent.ID)

	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoEvent(rows *sql.Rows) (*types.Event, error) {
	event := new(types.Event)

	err := rows.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.ImageUrl,
		&event.Date,
		&event.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return event, nil
}
