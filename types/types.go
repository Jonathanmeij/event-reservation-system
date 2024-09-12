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

type UpdateEventRequest struct {
	ID          int       `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ImageUrl    string    `json:"imageUrl" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

func (r *UpdateEventRequest) ToEvent() Event {
	return Event{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		ImageUrl:    r.ImageUrl,
		Date:        r.Date,
	}
}

type EventStore interface {
	GetEventByID(id int) (*Event, error)
	GetEvents() ([]*Event, error)
	CreateEvent(event Event) error
	DeleteEvent(id int) error
	UpdateEvent(event Event) error
}

//users

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}
