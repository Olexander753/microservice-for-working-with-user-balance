package service

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetHistory(ctx context.Context, id int) ([]schema.Transaction, error) {
	return s.repo.GetHistory(ctx, id)
}

func (s *BalanceService) Replenishment(ctx context.Context, replenishment schema.Balance) (schema.Balance, error) {
	return s.repo.Replenishment(ctx, replenishment)
}

func (s *BalanceService) GetBalance(ctx context.Context, id int) (schema.Balance, error) {
	return s.repo.GetBalance(ctx, id)
}

func (s *BalanceService) WriteOff(ctx context.Context, writeOff schema.Balance) (schema.Balance, error) {
	return s.repo.WriteOff(ctx, writeOff)
}

func (s *BalanceService) Transaction(ctx context.Context, transaction schema.Transaction) (schema.Balance, error) {
	return s.repo.Transaction(ctx, transaction)
}
