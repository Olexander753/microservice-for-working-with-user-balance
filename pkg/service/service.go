package service

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type Balance interface {
	GetHistory(ctx context.Context, id int) ([]schema.Transaction, error)
	Replenishment(ctx context.Context, replenishment schema.Balance) (schema.Balance, error)
	GetBalance(ctx context.Context, id int) (schema.Balance, error)
	WriteOff(ctx context.Context, writeOff schema.Balance) (schema.Balance, error)
	Transaction(ctx context.Context, transaction schema.Transaction) (schema.Balance, error)
}

type Service struct {
	Balance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Balance: NewBalanceService(repo.Balance),
	}
}
