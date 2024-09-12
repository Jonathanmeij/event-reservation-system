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

// users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

func NewUser(firstName string, lastName string, email string, password string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
