package service

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type TansctionService struct {
	repo repository.Transaction
}

func NewTansctionService(repo repository.Transaction) *TansctionService {
	return &TansctionService{repo: repo}
}

func (s *TansctionService) GetHistory(ctx context.Context, id int) ([]schema.Transaction, error) {
	return s.repo.GetHistory(ctx, id)
}

func (s *TansctionService) Transaction(ctx context.Context, transaction schema.Transaction) (schema.Balance, error) {
	return s.repo.Transaction(ctx, transaction)
}
