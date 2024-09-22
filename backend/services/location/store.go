package location

import (
	"github.com/jonathanmeij/go-reservation/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetLocationByID(id int) (*types.LocationEntity, error) {
	var location types.LocationEntity
	result := s.db.First(&location, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &location, nil
}

func (s *Store) GetLocations() ([]*types.LocationEntity, error) {
	var locations []*types.LocationEntity
	result := s.db.Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}

	return locations, nil
}

func (s *Store) CreateLocation(location types.LocationEntity) error {
	result := s.db.Create(&location)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) DeleteLocation(id int) error {
	result := s.db.Delete(&types.LocationEntity{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) UpdateLocation(updatedLocation types.LocationEntity) error {
	result := s.db.Model(&types.LocationEntity{}).Where("id = ?", updatedLocation.ID).Updates(updatedLocation)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
