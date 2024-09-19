package types

type EventStore interface {
	GetEventByID(id int) (*EventEntity, error)
	GetEvents() ([]*EventEntity, error)
	CreateEvent(event EventEntity) error
	DeleteEvent(id int) error
	UpdateEvent(event *EventEntity) error
}

type UserStore interface {
	GetUserByEmail(email string) (*UserEntity, error)
	GetUserByID(id int) (*UserEntity, error)
	CreateUser(UserEntity) error
}

type PlannedEventStore interface {
	GetPlannedEventByID(id int) (*PlannedEventEntity, error)
	CreatePlannedEvent(PlannedEventEntity) error
	DeletePlannedEvent(id int) error
	UpdatePlannedEvent(PlannedEventEntity) error
}
