package service

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

type TansctionService struct {
	repo repository.Transaction
}

// констурктор для структуры TansctionService, принимает repository.Transaction, возвращяет экземпляр структуры TansctionService
func NewTansctionService(repo repository.Transaction) *TansctionService {
	return &TansctionService{repo: repo}
}

// реализация интерфейса Transaction
// функция получения истории операций с балансом пользователя
func (s *TansctionService) GetHistory(id string) ([]schema.Transaction, error) {
	return s.repo.GetHistory(id)
}

// функция перевода денег между пользователями
func (s *TansctionService) Transaction(transaction schema.Transaction) (schema.Balance, error) {
	return s.repo.Transaction(transaction)
}
