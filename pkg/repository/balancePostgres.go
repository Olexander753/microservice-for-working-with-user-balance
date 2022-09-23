package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

// структура BalancePostgres, включает в себя sqlx.DB
type BalancePostgres struct {
	db *sqlx.DB
}

// констурктор для структуры BalancePostgres, принимает sqlx.DB, возвращяет экземпляр структуры BalancePostgres
func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

// Реализаци интерфейса Balance
// функция пополнения баланса пользователя
func (b *BalancePostgres) Replenishment(replenishment schema.Balance) (schema.Balance, error) {
	var balance schema.Balance

	// проверяем наличие данного пользователя в базе даннх
	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		replenishment.Id)

	fmt.Println(replenishment.Amount)

	// если запись о пользователе отсутствует в бд, то записываем в таблицы данные
	if err == sql.ErrNoRows {
		// запись данных о пользователе в таблицу Users
		query = fmt.Sprintln("INSERT INTO Users VALUES($1, $2) RETURNING Users.amount;")
		row := b.db.QueryRow(
			query,
			replenishment.Id,
			replenishment.Amount)

		// считывание записанного баланса, если ошибка, то возвращем ее и пустую структуру
		if err = row.Scan(&balance.Amount); err != nil {
			log.Println(err)
			return balance, err
		}

		// запись данных о пополнении в таблицу Transactions
		query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation, date_) VALUES($1, $2, $3, $4, $5);")
		row = b.db.QueryRow(
			query,
			replenishment.Id,
			"'Bank'",
			replenishment.Amount,
			fmt.Sprintf("'Replenishment %v RUB'", replenishment.Amount),
			"NOW()")
		// проверка на ошибки
		if err = row.Err(); err != nil {
			log.Println(err)
		}

	} else if err != nil && err != sql.ErrNoRows {
		// если ошибка и она не соотвествует ошибке sql.ErrNoRows, то возвращем ее и пустую структуру
		log.Println(err)
		return balance, err

	} else {
		// если не прошло ни одно из прошлых условий, значит запись о пользователе уже есть
		// изменяем баланс пользователя прибавив сумму из запроса к балансу пользователя
		query := fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
		row := b.db.QueryRow(
			query,
			replenishment.Amount+balance.Amount,
			replenishment.Id)

		// считывание записанного баланса, если ошибка, то возвращем ее и пустую структуру
		if err = row.Scan(&balance.Amount); err != nil {
			log.Println(err)
			return balance, err
		}

		// запись данных о пополнении в таблицу Transactions
		query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation, date_) VALUES($1, $2, $3, $4, $5);")
		row = b.db.QueryRow(
			query,
			replenishment.Id, "'Bank'",
			balance.Amount,
			fmt.Sprintf("'Replenishment %v RUB'", replenishment.Amount),
			"NOW()")

		// проверка на ошибки
		if err := row.Err(); err != nil {
			log.Println(err)
		}
	}
	balance.Id = replenishment.Id
	// при успешном пополнении возвращем экземпляр стуктуры Balance и нил
	return balance, nil
}

// функция получения текущего баланса пользователя
func (b *BalancePostgres) GetBalance(id string) (schema.Balance, error) {
	var balance schema.Balance

	// считывание данных из бд
	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		id)
	// возвращем экземпляр стуктуры Balance и ошибку, при успеном считвании она нил
	return balance, err
}

// функция списаниея средст с баланса пользователя
func (b *BalancePostgres) WriteOff(writeOff schema.Balance) (schema.Balance, error) {
	var balance schema.Balance

	// считывание данных о пользователе
	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		writeOff.Id)

	// если ошибка то возвращается пустая структура и ошибка
	if err != nil {
		log.Println(err)
		return balance, err
	}

	// проверка наличия необходимой суммы на балансе пользователя
	if balance.Amount < writeOff.Amount {
		return balance, fmt.Errorf("Недостаточно средств для списания")
	}

	// запрос изменения данных о балнсе пользователя списав с его баланса сумму из запроса
	query = fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
	row := b.db.QueryRow(
		query,
		balance.Amount-writeOff.Amount,
		writeOff.Id)

	// считывание записанного баланса, если ошибка, то возвращем ее и пустую структуру
	if err = row.Scan(&balance.Amount); err != nil {
		log.Println(err)
		return balance, err
	}

	// запись данных о пополнении в таблицу Transactions
	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation, date_) VALUES($1, $2, $3, $4, $5);")
	row = b.db.QueryRow(
		query,
		writeOff.Id,
		"'Bank'",
		balance.Amount,
		fmt.Sprintf("'Write-off %v RUB'", writeOff.Amount),
		"NOW()")

	// проверка на наличие ошибки
	if err = row.Err(); err != nil {
		log.Println(err)
	}

	return balance, nil
}
