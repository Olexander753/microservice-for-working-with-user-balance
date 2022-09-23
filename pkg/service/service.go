package service

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type Balance interface {
	Replenishment(replenishment schema.Balance) (schema.Balance, error)
	GetBalance(id string) (schema.Balance, error)
	WriteOff(writeOff schema.Balance) (schema.Balance, error)
}

type Transaction interface {
	GetHistory(id string) ([]schema.Transaction, error)
	Transaction(transaction schema.Transaction) (schema.Balance, error)
}

type Service struct {
	Balance
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Balance:     NewBalanceService(repo.Balance),
		Transaction: NewTansctionService(repo.Transaction),
	}
}
