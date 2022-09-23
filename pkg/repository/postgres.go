package repository

import (
	"fmt"
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/config"
	"github.com/jmoiron/sqlx"
)

// функци подключения к базе данных postgres, возращает указатель на sqlx.DB и ошибку
func NewPostgresBD(cfg *config.Config) (*sqlx.DB, error) {
	log.Println("Connect to postgres db")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username,
		cfg.Postgres.DBName, cfg.Postgres.Password, cfg.Postgres.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
