package storage

import (
	"database/sql"
	"fmt"

	"github.com/jonathanmeij/go-reservation/pkg/models"
)

type PostgresAccountStorage struct {
	db *sql.DB
}

type AccountStorage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAcxcount(*models.Account) error
	GetAccountByID(int) (*models.Account, error)
	GetAccounts() ([]*models.Account, error)
}

func NewPostGresAccountStorage(db *sql.DB) *PostgresAccountStorage {
	return &PostgresAccountStorage{db: db}
}

func (s *PostgresAccountStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresAccountStorage) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account  (
		id SERIAL PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresAccountStorage) CreateAccount(acc *models.Account) error {
	query := `INSERT INTO account 
	(first_name, last_name, created_at) 
	VALUES ($1, $2, $3, $4, $5)`

	res, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresAccountStorage) UpdateAccount(a *models.Account) error {
	return nil
}

func (s *PostgresAccountStorage) DeleteAccount(id int) error {
	query := `DELETE FROM account WHERE id = $1`
	_, err := s.db.Query(query, id)
	return err
}

func (s *PostgresAccountStorage) GetAccountByID(id int) (*models.Account, error) {
	query := `SELECT * FROM account WHERE id = $1`
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account not found")
}

func (s *PostgresAccountStorage) GetAccounts() ([]*models.Account, error) {
	query := `SELECT * FROM account`
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.CreatedAt)

	return account, err
}
