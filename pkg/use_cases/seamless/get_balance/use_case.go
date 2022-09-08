package get_balance

import (
	"context"
	"errors"

	"github.com/maskot/pkg/helpers/use_cases"
	"github.com/maskot/pkg/repositories"
)

type UseCase[In, Out any] struct {
	repo repositories.Balance
}

func NewUseCase(repo repositories.Balance) use_cases.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		repo: repo,
	}
}

func (uc *UseCase[In, Out]) Execute(ctx context.Context, input *Input) (*Output, error) {
	balance, err := uc.repo.Get(ctx, input.PlayerName)
	if err != nil {
		return nil, errors.New("balance not found")
	}

	return &Output{
		Balance: balance.Balance,
	}, nil
}
