package get_balance

import (
	"context"
	"errors"

	"github.com/maskot/internal/helpers/use_cases"
	"github.com/maskot/internal/repo"
)

type UseCase[In, Out any] struct {
	balanceRepo repo.Balance
}

func NewUseCase(balanceRepo repo.Balance) use_cases.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		balanceRepo: balanceRepo,
	}
}

func (uc *UseCase[In, Out]) Execute(ctx context.Context, input *Input) (*Output, error) {
	balance, err := uc.balanceRepo.Find(ctx, input.PlayerName)
	if err != nil {
		return nil, errors.New("balance not found")
	}

	return &Output{
		Balance: balance.Balance,
	}, nil
}
