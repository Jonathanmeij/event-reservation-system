package types

import (
	"time"

	"gorm.io/gorm"
)

type EventEntity struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Description   string `gorm:"not null"`
	ImageUrl      string
	Date          time.Time            `gorm:"not null"`
	PlannedEvents []PlannedEventEntity `gorm:"foreignKey:EventID"`
}

func (e *EventEntity) ToEvent() Event {
	return Event{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		ImageUrl:    e.ImageUrl,
		Date:        e.Date,
		CreatedAt:   e.CreatedAt,
	}
}

type LocationEntity struct {
	ID             int    `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	AmountOfPeople int    `gorm:"not null"`
}

type TicketEntity struct {
	gorm.Model
	PlannedEventID int                `gorm:"not null"`
	PlannedEvent   PlannedEventEntity `gorm:"foreignKey:PlannedEventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PurchaseDate   time.Time          `gorm:"not null"`
	SeatNumber     int                `gorm:"not null"`
	UserID         int                `gorm:"not null"`
	User           UserEntity         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PlannedEventEntity struct {
	gorm.Model
	EventID    int            `gorm:"not null"`
	Event      EventEntity    `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LocationID int            `gorm:"not null"`
	Location   LocationEntity `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Date       time.Time      `gorm:"not null"`
	Tickets    []TicketEntity `gorm:"foreignKey:PlannedEventID"`
}

func (p *PlannedEventEntity) ToPlannedEvent() PlannedEvent {
	return PlannedEvent{
		ID:         p.ID,
		EventID:    p.EventID,
		LocationID: p.LocationID,
		Date:       p.Date,
		Location: Location{
			ID:             p.Location.ID,
			Name:           p.Location.Name,
			AmountOfPeople: p.Location.AmountOfPeople,
		},
	}
}

type UserEntity struct {
	gorm.Model
	FirstName string         `gorm:"not null"`
	LastName  string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Role      string         `gorm:"not null"`
	Tickets   []TicketEntity `gorm:"foreignKey:UserID"`
}
