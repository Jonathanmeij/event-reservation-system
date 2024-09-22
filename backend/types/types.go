package types

import "time"

// event
type Event struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	ImageUrl      string         `json:"imageUrl"`
	CreatedAt     time.Time      `json:"createdAt"`
	PlannedEvents []PlannedEvent `json:"plannedEvents"`
}

type CreateEventRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ImageUrl    string    `json:"imageUrl" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

func (r *CreateEventRequest) ToEvent() EventEntity {
	return EventEntity{
		Title:       r.Title,
		Description: r.Description,
		ImageUrl:    r.ImageUrl,
	}
}

type UpdateEventRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ImageUrl    string    `json:"imageUrl" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

func (r *UpdateEventRequest) ToEvent() EventEntity {
	return EventEntity{
		Title:       r.Title,
		Description: r.Description,
		ImageUrl:    r.ImageUrl,
	}
}

type PlannedEvent struct {
	ID         uint      `json:"id"`
	EventID    int       `json:"eventId"`
	LocationID int       `json:"locationId"`
	Date       time.Time `json:"date"`
	Location   Location  `json:"location"`
}

type CreatePlannedEventRequest struct {
	EventID    int       `json:"eventId" validate:"required"`
	LocationID int       `json:"locationId" validate:"required"`
	Date       time.Time `json:"date" validate:"required"`
}

func (r *CreatePlannedEventRequest) ToPlannedEvent() PlannedEventEntity {
	return PlannedEventEntity{
		EventID:    r.EventID,
		LocationID: r.LocationID,
		Date:       r.Date,
	}
}

// Location
type Location struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AmountOfPeople int    `json:"amountOfPeople"`
}

type CreateLocationRequest struct {
	Name           string `json:"name" validate:"required"`
	AmountOfPeople int    `json:"amountOfPeople" validate:"required"`
}

func NewLocation(locationEntity LocationEntity) *Location {
	return &Location{
		ID:             locationEntity.ID,
		Name:           locationEntity.Name,
		AmountOfPeople: locationEntity.AmountOfPeople,
	}
}

func (r *CreateLocationRequest) ToLocationEntity() LocationEntity {
	return LocationEntity{
		Name:           r.Name,
		AmountOfPeople: r.AmountOfPeople,
	}
}

// users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role" tstype:"'admin' | 'user'"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

func NewUser(firstName string, lastName string, email string, password string) *UserEntity {
	return &UserEntity{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Role:      "user",
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}
