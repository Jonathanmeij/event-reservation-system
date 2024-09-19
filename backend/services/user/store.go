package user

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

func (s *Store) CreateUser(user types.UserEntity) error {
	result := s.db.Create(&user)

	return result.Error
}

func (s *Store) GetUserByEmail(email string) (*types.UserEntity, error) {
	var user types.UserEntity
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}

	return &user, nil
}

func (s *Store) GetUserByID(id int) (*types.UserEntity, error) {
	var u types.UserEntity
	result := s.db.First(&u, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}

	return &u, nil
}
