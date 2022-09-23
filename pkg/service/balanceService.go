package service

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

// структура BalanceService включает в себя интерфейс Balance
type BalanceService struct {
	repo repository.Balance
}

// констурктор для структуры BalanceService, принимает repository.Balance, возвращяет экземпляр структуры BalanceService
func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

// реализация интерфейса Balance
// функция пополнения баланса пользователя
func (s *BalanceService) Replenishment(replenishment schema.Balance) (schema.Balance, error) {
	return s.repo.Replenishment(replenishment)
}

// функция получения текущего баланса пользователя
func (s *BalanceService) GetBalance(id string) (schema.Balance, error) {
	return s.repo.GetBalance(id)
}

// функция списаниея средст с баланса пользователя
func (s *BalanceService) WriteOff(writeOff schema.Balance) (schema.Balance, error) {
	return s.repo.WriteOff(writeOff)
}
