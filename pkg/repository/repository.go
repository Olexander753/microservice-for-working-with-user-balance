package repository

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
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

// структура Repository включает в себя интерфейсы Balance и Transaction
type Repository struct {
	Balance
	Transaction
}

// конструктор для структуры Repository принимает указатель sqlx.DB и возвращает указатель на экземпляр структуры Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance:     NewBalancePostgres(db),
		Transaction: NewTransactionPostgres(db),
	}
}
