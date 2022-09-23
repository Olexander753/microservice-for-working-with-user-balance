package repository

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	Replenishment(replenishment schema.Balance) (schema.Balance, error)
	GetBalance(id int) (schema.Balance, error)
	WriteOff(writeOff schema.Balance) (schema.Balance, error)
}

type Transaction interface {
	GetHistory(id int) ([]schema.Transaction, error)
	Transaction(transaction schema.Transaction) (schema.Balance, error)
}

type Repository struct {
	Balance
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance:     NewBalancePostgres(db),
		Transaction: NewTransactionPostgres(db),
	}
}
