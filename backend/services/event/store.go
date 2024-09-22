package event

import (
	"fmt"

	"github.com/jonathanmeij/go-reservation/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetEventByID(id int) (*types.EventEntity, error) {
	var event types.EventEntity
	result := s.db.First(&event, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get event: %w", result.Error)
	}

	return &event, nil
}

func (s *Store) GetEvents() ([]*types.EventEntity, error) {
	var events []*types.EventEntity
	result := s.db.Find(&events)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get events: %w", result.Error)
	}

	return events, nil
}

func (s *Store) GetEventsWithPlannedEvents() ([]*types.EventEntity, error) {
	var events []*types.EventEntity
	result := s.db.Preload("PlannedEvents").Preload("Location").Find(&events)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get events with planned events: %w", result.Error)
	}

	return events, nil
}

func (s *Store) CreateEvent(event types.EventEntity) error {
	result := s.db.Create(&event)
	if result.Error != nil {
		return fmt.Errorf("failed to create event: %w", result.Error)
	}

	return nil
}

func (s *Store) DeleteEvent(id int) error {
	result := s.db.Delete(&types.EventEntity{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete event: %w", result.Error)
	}

	return nil
}

func (s *Store) UpdateEvent(updatedEvent *types.EventEntity) error {
	result := s.db.Model(&types.EventEntity{}).Where("id = ?", updatedEvent.ID).Updates(updatedEvent)
	if result.Error != nil {
		return fmt.Errorf("failed to update event: %w", result.Error)
	}

	return nil
}
