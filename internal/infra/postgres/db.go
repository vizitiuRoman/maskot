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

func NewPool(cfg *Config) (*sqlx.DB, func() error, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s dbname=%s password=%s port=%s host=%s sslmode=%s",
			cfg.Username, cfg.DBName, cfg.Password, cfg.Port, cfg.Host, cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, db.Close, err
	}

	if err := db.Ping(); err != nil {
		return nil, db.Close, err
	}

	return db, db.Close, nil
}
