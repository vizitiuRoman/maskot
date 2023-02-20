package rollbacktransaction

import (
	"context"
	"fmt"

	"github.com/maskot/internal/helpers/usecase"
	"github.com/maskot/internal/repo"
)

type UseCase[In, Out any] struct {
	balanceRepo     repo.Balance
	transactionRepo repo.Transaction
}

func NewUseCase(balanceRepo repo.Balance, transactionRepo repo.Transaction) usecase.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		balanceRepo:     balanceRepo,
		transactionRepo: transactionRepo,
	}
}

func (uc *UseCase[In, Out]) Execute(ctx context.Context, input *Input) (*Output, error) {
	t, err := uc.transactionRepo.Find(ctx, input.TransactionRef)
	if err != nil {
		return nil, fmt.Errorf("cannot find transaction: %w", err)
	}

	if t.IsRolledBack {
		return &Output{}, nil
	}

	err = uc.transactionRepo.Rollback(ctx, t.ID, input.PlayerName)
	if err != nil {
		return nil, fmt.Errorf("cannot rollback transaction: %w", err)
	}

	return &Output{}, nil
}
