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

type UserEntity struct {
	gorm.Model
	FirstName string         `gorm:"not null"`
	LastName  string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Role      string         `gorm:"not null"`
	Tickets   []TicketEntity `gorm:"foreignKey:UserID"`
}
