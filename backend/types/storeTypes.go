package types

type EventStore interface {
	GetEventByID(id int) (*Event, error)
	GetEvents() ([]*Event, error)
	CreateEvent(event Event) error
	DeleteEvent(id int) error
	UpdateEvent(event Event) error
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}
