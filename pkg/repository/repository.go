package repository

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	Replenishment(ctx context.Context, replenishment schema.Balance) (schema.Balance, error)
	GetBalance(ctx context.Context, id int) (schema.Balance, error)
	WriteOff(ctx context.Context, writeOff schema.Balance) (schema.Balance, error)
}

type Transaction interface {
	GetHistory(ctx context.Context, id int) ([]schema.Transaction, error)
	Transaction(ctx context.Context, transaction schema.Transaction) (schema.Balance, error)
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
