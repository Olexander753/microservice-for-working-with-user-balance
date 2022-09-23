package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (b *BalancePostgres) Replenishment(replenishment schema.Balance) (schema.Balance, error) {
	var balance schema.Balance

	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		replenishment.Id)

	fmt.Println(replenishment.Amount)

	if err == sql.ErrNoRows {
		query = fmt.Sprintln("INSERT INTO Users VALUES($1, $2) RETURNING Users.amount;")
		row := b.db.QueryRow(
			query,
			replenishment.Id,
			replenishment.Amount)

		if err = row.Scan(&balance.Amount); err != nil {
			log.Println(err)
			return balance, err
		}

		fmt.Println(replenishment.Amount)

		query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation) VALUES($1, $2, $3, $4);")
		row = b.db.QueryRow(
			query,
			replenishment.Id,
			"'Bank'",
			replenishment.Amount,
			fmt.Sprintf("'Replenishment %v'", replenishment.Amount))

		if err = row.Err(); err != nil {
			log.Println(err)
		}

	} else if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return balance, err

	} else {
		query := fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2;")
		row := b.db.QueryRow(
			query,
			replenishment.Amount+balance.Amount,
			replenishment.Id)

		if err = row.Err(); err != nil {
			log.Println(err)
			return balance, err
		}

		query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation) VALUES($1, $2, $3, $4);")
		row = b.db.QueryRow(
			query,
			replenishment.Id, "'Bank'",
			replenishment.Amount+balance.Amount,
			fmt.Sprintf("'Replenishment %v'", replenishment.Amount))

		if err := row.Err(); err != nil {
			log.Println(err)
		}

		balance.Amount = replenishment.Amount + balance.Amount

	}
	balance.Id = replenishment.Id
	return balance, nil
}

func (b *BalancePostgres) GetBalance(id string) (schema.Balance, error) {
	var balance schema.Balance

	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		id)

	return balance, err
}

func (b *BalancePostgres) WriteOff(writeOff schema.Balance) (schema.Balance, error) {
	var balance schema.Balance

	query := fmt.Sprintln("SELECT * FROM Users WHERE ID = $1")
	err := b.db.Get(
		&balance,
		query,
		writeOff.Id)

	if err != nil {
		log.Println(err)
		return balance, err
	}

	if balance.Amount < writeOff.Amount {
		return balance, fmt.Errorf("Недостаточно средств для списания!")
	}

	query = fmt.Sprintln("UPDATE Users SET amount = $1 WHERE id = $2 RETURNING Users.amount;")
	row := b.db.QueryRow(
		query,
		balance.Amount-writeOff.Amount,
		writeOff.Id)

	if err = row.Err(); err != nil {
		log.Println(err)
		return balance, err
	}

	query = fmt.Sprintln("INSERT INTO Transactions(recipient, sender, amount, operation) VALUES($1, $2, $3, $4);")
	row = b.db.QueryRow(
		query,
		writeOff.Id,
		"'Bank'",
		balance.Amount-writeOff.Amount,
		fmt.Sprintf("'Write-off %v'", writeOff.Amount))

	if err = row.Err(); err != nil {
		log.Println(err)
	}

	balance.Amount = balance.Amount - writeOff.Amount

	return balance, nil
}
