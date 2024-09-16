package user

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

func (s *Store) CreateUser(user types.User) error {
	// _, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, role, password) VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Role, user.Password)

	return err
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	// if err != nil {
	// 	return nil, err
	// }

	// u := new(types.User)
	// for rows.Next() {
	// 	u, err = scanRowsIntoUser(rows)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if u.ID == 0 {
	// 	return nil, fmt.Errorf("user not found")
	// }

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	// rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return nil, err
	// }

	// u := new(types.User)
	// for rows.Next() {
	// 	u, err = scanRowsIntoUser(rows)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if u.ID == 0 {
	// 	return nil, fmt.Errorf("user not found")
	// }

	return u, nil
}
