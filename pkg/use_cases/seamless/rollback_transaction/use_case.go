package rollback_transaction

import (
	"context"

	"github.com/maskot/pkg/helpers/use_cases"
	"github.com/maskot/pkg/repositories"
	"github.com/maskot/pkg/repositories/postgres/balance"
)

type UseCase[In, Out any] struct {
	repo *repositories.Repository
}

func NewUseCase(repos *repositories.Repository) use_cases.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		repo: repos,
	}
}

func (uc *UseCase[In, Out]) Execute(ctx context.Context, input *Input) (*Output, error) {
	b, err := uc.repo.Balance.Get(ctx, input.PlayerName)
	if err != nil {
		return nil, err
	}

	t, err := uc.repo.Transaction.Get(ctx, input.TransactionRef)
	if err != nil {
		return nil, err
	}
	if t.IsRollback {
		return &Output{}, nil
	}

	err = uc.repo.Transaction.Rollback(ctx, t.ID)
	if err != nil {
		return nil, err
	}

	_, err = uc.repo.Balance.Update(ctx, &balance.Balance{
		PlayerName: b.PlayerName,
		Balance:    b.Balance + t.Withdraw - t.Deposit,
	})
	if err != nil {
		return nil, err
	}

	return &Output{}, nil
}
