package event

import (
	"database/sql"

	"github.com/jonathanmeij/go-reservation/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetEventByID(ID int) error {
	return nil
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
	(title, description, imageUrl, date, createdAt) 
	VALUES (?,?,?,?,?)`, event.Title, event.Description, event.ImageUrl, event.Date, event.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteEvent() error {
	return nil
}

func (s *Store) UpdateEvent() error {
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
