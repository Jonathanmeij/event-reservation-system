package plannedEvents

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

func (s *Store) GetPlannedEventByID(id int) (*types.PlannedEventEntity, error) {
	var plannedEvent types.PlannedEventEntity
	result := s.db.First(&plannedEvent, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get planned event: %w", result.Error)
	}

	return &plannedEvent, nil
}

func (s *Store) CreatePlannedEvent(PlannedEvent types.PlannedEventEntity) error {
	result := s.db.Create(&PlannedEvent)
	if result.Error != nil {
		return fmt.Errorf("failed to create planned event: %w", result.Error)
	}

	return nil
}

func (s *Store) DeletePlannedEvent(id int) error {
	result := s.db.Delete(&types.PlannedEventEntity{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete planned event: %w", result.Error)
	}

	return nil
}

func (s *Store) UpdatePlannedEvent(updatedPlannedEvent types.PlannedEventEntity) error {
	result := s.db.Model(&types.EventEntity{}).Where("id = ?", updatedPlannedEvent.ID).Updates(updatedPlannedEvent)
	if result.Error != nil {
		return fmt.Errorf("failed to update planned event: %w", result.Error)
	}

	return nil
}
