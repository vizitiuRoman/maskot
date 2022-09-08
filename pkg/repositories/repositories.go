package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/maskot/pkg/repositories/postgres/balance"
	"github.com/maskot/pkg/repositories/postgres/transaction"
)

//go:generate mockery --dir . --name Wallet --name Invoice --output ./mocks

type Transaction interface {
	Create(context.Context, *transaction.Transaction) (*transaction.Transaction, error)
	Rollback(ctx context.Context, id int) error
	Get(context.Context, string) (*transaction.Transaction, error)
}

type Balance interface {
	Create(context.Context, *balance.Balance) (*balance.Balance, error)
	Update(context.Context, *balance.Balance) (*balance.Balance, error)
	Get(context.Context, string) (*balance.Balance, error)
	GetOrCreate(context.Context, string) (*balance.Balance, error)
}

type Repository struct {
	Transaction Transaction
	Balance     Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: transaction.NewRepository(db),
		Balance:     balance.NewRepository(db),
	}
}
