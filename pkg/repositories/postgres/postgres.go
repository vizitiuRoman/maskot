package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s dbname=%s password=%s port=%s host=%s sslmode=%s",
			cfg.Username, cfg.DBName, cfg.Password, cfg.Port, cfg.Host, cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
