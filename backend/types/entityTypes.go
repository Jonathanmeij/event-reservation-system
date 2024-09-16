package types

import (
	"time"
)

type EventEntity struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Title         string `gorm:"not null"`
	Description   string `gorm:"not null"`
	ImageUrl      string
	Date          time.Time            `gorm:"not null"`
	CreatedAt     time.Time            `gorm:"autoCreateTime"`
	PlannedEvents []PlannedEventEntity `gorm:"foreignKey:EventID"`
	Tickets       []TicketEntity       `gorm:"foreignKey:EventID"`
}

type LocationEntity struct {
	ID             int    `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	AmountOfPeople int    `gorm:"not null"`
}

type TicketEntity struct {
	ID           int         `gorm:"primaryKey;autoIncrement"`
	EventID      int         `gorm:"not null"`
	Event        EventEntity `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PurchaseDate time.Time   `gorm:"not null"`
	SeatNumber   int         `gorm:"not null"`
	UserID       int         `gorm:"not null"`
	User         UserEntity  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PlannedEventEntity struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	EventID    int            `gorm:"not null"`
	Event      EventEntity    `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LocationID int            `gorm:"not null"`
	Location   LocationEntity `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Date       time.Time      `gorm:"not null"`
}

type UserEntity struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	FirstName string         `gorm:"not null"`
	LastName  string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Role      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	Tickets   []TicketEntity `gorm:"foreignKey:UserID"`
}
