package repository

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (b *BalancePostgres) Replenishment(ctx context.Context, replenishment schema.Balance) (schema.Balance, error) {
	var balance schema.Balance
	var err error
	return balance, err
}

func (b *BalancePostgres) GetBalance(ctx context.Context, id int) (schema.Balance, error) {
	var balance schema.Balance
	var err error
	return balance, err
}

func (b *BalancePostgres) WriteOff(ctx context.Context, replenishment schema.Balance) (schema.Balance, error) {
	var balance schema.Balance
	var err error
	return balance, err
}
