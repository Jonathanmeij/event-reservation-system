package types

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateEventRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ImageUrl    string    `json:"imageUrl" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

func (r *CreateEventRequest) ToEvent() Event {
	return Event{
		Title:       r.Title,
		Description: r.Description,
		ImageUrl:    r.ImageUrl,
		Date:        r.Date,
		CreatedAt:   time.Now(),
	}
}

type EventStore interface {
	GetEventByID(ID int) error
	GetEvents() ([]*Event, error)
	CreateEvent(event Event) error
	DeleteEvent() error
	UpdateEvent() error
}
