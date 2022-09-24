package repository

import (
	"fmt"
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

// структура TransactionPostgres, включает в себя sqlx.DB
type TransactionPostgres struct {
	db *sqlx.DB
}

// констурктор для структуры TransactionPostgres, принимает sqlx.DB, возвращяет экземпляр структуры TransactionPostgres
func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

// реализация интерфейса Transaction
// функция получения истории операций с балансом пользователя
func (t *TransactionPostgres) GetHistory(id string) ([]schema.Transaction, error) {
	var transactions []schema.Transaction

	// считывание данных из бд
	query := fmt.Sprintln("SELECT * FROM Transactions WHERE recipient = $1")
	err := t.db.Select(
		&transactions,
		query,
		id)

	// возвращем слайс экземпляров стуктуры Transaction и ошибку, при успеном считвании она нил
	return transactions, err
}

// функция перевода денег между пользователями
func (t *TransactionPostgres) Transaction(transaction schema.Transaction) (schema.Balance, error) {
	var balance_sender, balance_recipient schema.Balance

	// прверем на наличие отправителя в базе
	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := t.db.Get(
		&balance_sender,
		query,
		transaction.Sender)

	// если ошибка то возвращается пустая структура и ошибка
	if err != nil {
		log.Println(err)
		return balance_sender, err
	}

	// проверка на наличие средств для перевода
	if balance_sender.Amount < transaction.Amount {
		return balance_sender, fmt.Errorf("Недостаточно средств для перевода")
	}

	//проверка на наличие получателя в базе
	query = fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err = t.db.Get(
		&balance_recipient,
		query,
		transaction.Recipient)

	// если ошибка то возвращается пустая структура и ошибка
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

	// если ошибка то возвращается пустая структура и ошибка
	if err = row.Err(); err != nil {
		log.Println(err)
		return balance_sender, err
	}

	//записываем данную операцию в таблицу транзакций
	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, balance, amount, operation, date_) VALUES($1, $2, $3, $4, $5, $6);")
	row = t.db.QueryRow(
		query,
		transaction.Recipient,
		transaction.Sender,
		balance_sender.Amount-transaction.Amount,
		transaction.Amount,
		"'Transaction write-off'",
		"NOW()")

	// если ошибка то возвращается пустая структура и ошибка
	if err = row.Err(); err != nil {
		log.Println(err)
	}

	// изменяем количество денег на счету у получателя
	query = fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
	row = t.db.QueryRow(
		query,
		balance_recipient.Amount+transaction.Amount,
		transaction.Recipient)

	// считывание записанного баланса, если ошибка, то возвращем ее и пустую структуру
	if err = row.Scan(&balance_sender.Amount); err != nil {
		log.Println(err)
		return balance_sender, err
	}

	//записываем данную операцию в таблицу транзакций
	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, balance, amount, operation, date_) VALUES($1, $2, $3, $4, $5, $6);")
	row = t.db.QueryRow(
		query,
		transaction.Recipient,
		transaction.Sender,
		balance_recipient.Amount,
		transaction.Amount,
		"'Transaction replenishment'",
		"NOW()")

	// если ошибка то возвращается пустая структура и ошибка
	if err = row.Err(); err != nil {
		log.Println(err)
	}

	balance_sender.Amount = balance_sender.Amount - transaction.Amount

	return balance_sender, nil
}
