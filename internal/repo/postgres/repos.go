package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/maskot/internal/repo"
)

type Repos struct {
	Transaction repo.Transaction
	Balance     repo.Balance
}

func NewRepos(db *sqlx.DB) *Repos {
	return &Repos{
		Transaction: &TransactionRepository{db},
		Balance:     &BalanceRepository{db},
	}
}
