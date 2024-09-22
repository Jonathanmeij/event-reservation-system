package types

type EventStore interface {
	GetEventByID(id int) (*EventEntity, error)
	GetEvents() ([]*EventEntity, error)
	GetEventsWithPlannedEvents() ([]*EventEntity, error)
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
	GetPlannedEvents() ([]*PlannedEventEntity, error)
	CreatePlannedEvent(PlannedEvent PlannedEventEntity) error
	DeletePlannedEvent(id int) error
	UpdatePlannedEvent(updatedPlannedEvent PlannedEventEntity) error
}

type LocationStore interface {
	GetLocationByID(id int) (*LocationEntity, error)
	GetLocations() ([]*LocationEntity, error)
	CreateLocation(LocationEntity) error
	DeleteLocation(id int) error
	UpdateLocation(updatedLocation LocationEntity) error
}
