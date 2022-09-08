package withdraw_and_deposit

import (
	"context"

	"github.com/maskot/pkg/helpers/use_cases"
	"github.com/maskot/pkg/repositories"
	"github.com/maskot/pkg/repositories/postgres/transaction"
	"github.com/maskot/pkg/use_cases/errors"
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
	if input.Withdraw < 0 {
		return nil, errors.ErrNegativeWithdrawalCode
	}
	if input.Deposit < 0 {
		return nil, errors.ErrNegativeDepositCode
	}

	b, err := uc.repo.Balance.GetOrCreate(ctx, input.PlayerName)
	if err != nil {
		return nil, err
	}

	t, err := uc.repo.Transaction.Get(ctx, input.TransactionRef)
	if t != nil && err == nil {
		return &Output{
			Balance:       b.Balance,
			TransactionID: t.ID,
		}, nil
	}

	b.Balance -= input.Withdraw
	if b.Balance < 0 {
		return nil, errors.ErrNotEnoughMoneyCode
	}
	b.Balance += input.Deposit

	t, err = uc.repo.Transaction.Create(ctx, &transaction.Transaction{
		ID:                   input.ID,
		PlayerName:           input.PlayerName,
		Currency:             input.Currency,
		Withdraw:             input.Withdraw,
		Deposit:              input.Deposit,
		TransactionRef:       input.TransactionRef,
		GameRoundRef:         input.GameRoundRef,
		GameID:               input.GameID,
		Reason:               input.Reason,
		SessionID:            input.SessionID,
		SessionAlternativeID: input.SessionAlternativeID,
		SpinDetailsBetType:   input.SpinDetailsBetType,
		SpinDetailsWinType:   input.SpinDetailsWinType,
	})
	if err != nil {
		return nil, err
	}

	updatedBalance, err := uc.repo.Balance.Update(ctx, b)
	if err != nil {
		return nil, err
	}

	return &Output{
		Balance:       updatedBalance.Balance,
		TransactionID: t.ID,
	}, nil
}
