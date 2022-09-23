package repository

import (
	"context"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (p *TransactionPostgres) GetHistory(ctx context.Context, id int) ([]schema.Transaction, error) {
	var transactions []schema.Transaction
	var err error
	return transactions, err
}

func (p *TransactionPostgres) Transaction(ctx context.Context, transaction schema.Transaction) (schema.Balance, error) {
	var balance schema.Balance
	var err error
	return balance, err
}