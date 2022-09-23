package repository

import (
	"fmt"
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) GetHistory(id string) ([]schema.Transaction, error) {
	var transactions []schema.Transaction

	query := fmt.Sprintln("SELECT * FROM Transactions WHERE recipient = $1")
	err := t.db.Select(
		&transactions,
		query,
		id)

	return transactions, err
}

func (t *TransactionPostgres) Transaction(transaction schema.Transaction) (schema.Balance, error) {
	var balance_sender, balance_recipient schema.Balance

	// прверем на наличие отправителя в базе
	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := t.db.Get(
		&balance_sender,
		query,
		transaction.Sender)

	if err != nil {
		log.Println(err)
		return balance_sender, err
	}

	// проверка на наличие средств для перевода
	if balance_sender.Amount < transaction.Amount {
		return balance_sender, fmt.Errorf("Недостаточно средств для перевода!")
	}

	//проверка на наличие получателя в базе
	query = fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err = t.db.Get(
		&balance_recipient,
		query,
		transaction.Recipient)

	if err != nil {
		log.Println(err)
		return balance_sender, err
	}

	// изменяем количество денег на счету у отправителя
	query = fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
	row := t.db.QueryRow(
		query,
		balance_sender.Amount-transaction.Amount,
		transaction.Sender)

	if err = row.Err(); err != nil {
		log.Println(err)
		return balance_sender, err
	}

	//записываем данную операцию в таблицу транзакций
	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation) VALUES($1, $2, $3, $4);")
	row = t.db.QueryRow(
		query,
		transaction.Sender,
		"''",
		balance_sender.Amount-transaction.Amount,
		fmt.Sprintf("'Transaction write-off %v'", transaction.Amount))

	if err = row.Err(); err != nil {
		log.Println(err)
	}

	// изменяем количество денег на счету у получателя
	query = fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
	row = t.db.QueryRow(
		query,
		balance_recipient.Amount+transaction.Amount,
		transaction.Recipient)

	if err = row.Err(); err != nil {
		log.Println(err)
		return balance_sender, err
	}

	//записываем данную операцию в таблицу транзакций
	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation) VALUES($1, $2, $3, $4);")
	row = t.db.QueryRow(
		query,
		transaction.Recipient,
		transaction.Sender,
		balance_recipient.Amount+transaction.Amount,
		fmt.Sprintf("'Transaction replenishment %v'", transaction.Amount))

	if err = row.Err(); err != nil {
		log.Println(err)
	}

	balance_sender.Amount = balance_sender.Amount - transaction.Amount

	return balance_sender, nil
}
