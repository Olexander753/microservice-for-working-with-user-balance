package service

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
)

// интерфейс Balance описывает поведение работы с балансом
type Balance interface {
	Replenishment(replenishment schema.Balance) (schema.Balance, error)
	GetBalance(id string) (schema.Balance, error)
	WriteOff(writeOff schema.Balance) (schema.Balance, error)
}

// интерфейс Transaction описывает поведение работы с транзакциями
type Transaction interface {
	GetHistory(id string) ([]schema.Transaction, error)
	Transaction(transaction schema.Transaction) (schema.Balance, error)
}

// структура Service включает в себя интерфейсы Balance и Transaction
// сервис нужен для будущего подключения стороних сервисов,
// по типу брокера сообщений nats-striming
// на данный момент через него будут проходить HTTP запросы,
// а сервис будет уже обращаться к репозиторию,
// который в своб очередь обращается к хранилищу дыннх
// в данном случае бд postgres
type Service struct {
	Balance
	Transaction
}

// конструктор для структуры Service принимает указатель на экземпляр стуктуры Repository и возвращает указатель на экземпляр структуры Service
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Balance:     NewBalanceService(repo.Balance),
		Transaction: NewTansctionService(repo.Transaction),
	}
}
