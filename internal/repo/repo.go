package repo

import (
	"context"
	"errors"

	"github.com/maskot/internal/model"
)

var (
	ErrNotFound      = errors.New("model was not found")
	ErrAlreadyExists = errors.New("model already exists")
)

//go:generate mockery --dir . --name Wallet --name Invoice --output ./mocks

type Transaction interface {
	Create(context.Context, *model.Transaction) (balance int, err error)
	Rollback(ctx context.Context, id int, playerName string) error
	Find(context.Context, string) (*model.Transaction, error)
}

type Balance interface {
	Save(context.Context, *model.Balance) (*model.Balance, error)
	Update(context.Context, *model.Balance) (*model.Balance, error)
	Find(context.Context, string) (*model.Balance, error)
	FindOrCreate(context.Context, string) (*model.Balance, error)
}
