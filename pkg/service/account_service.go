package service

import (
	"github.com/jonathanmeij/go-reservation/pkg/models"
	"github.com/jonathanmeij/go-reservation/pkg/storage"
)

type AccountService struct {
	accountStorage *storage.PostgresAccountStorage
}

func NewAccountService(accountStorage *storage.PostgresAccountStorage) *AccountService {
	return &AccountService{accountStorage: accountStorage}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	// Business logic before calling storage
	return s.accountStorage.CreateAccount(account)
}
