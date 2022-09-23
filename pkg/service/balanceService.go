package service

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) Replenishment(replenishment schema.Balance) (schema.Balance, error) {
	return s.repo.Replenishment(replenishment)
}

func (s *BalanceService) GetBalance(id int) (schema.Balance, error) {
	return s.repo.GetBalance(id)
}

func (s *BalanceService) WriteOff(writeOff schema.Balance) (schema.Balance, error) {
	return s.repo.WriteOff(writeOff)
}
