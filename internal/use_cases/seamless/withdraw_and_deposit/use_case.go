package withdrawanddeposit

import (
	"context"
	"fmt"

	"github.com/maskot/internal/helpers/usecase"
	"github.com/maskot/internal/model"
	"github.com/maskot/internal/repo"
	"github.com/maskot/internal/use_cases/errors"
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
	if input.Withdraw < 0 {
		return nil, errors.ErrNegativeWithdrawalCode
	}
	if input.Deposit < 0 {
		return nil, errors.ErrNegativeDepositCode
	}

	b, err := uc.balanceRepo.FindOrCreate(ctx, input.PlayerName)
	if err != nil {
		return nil, fmt.Errorf("cannot find or create balance: %w", err)
	}

	t, err := uc.transactionRepo.Find(ctx, input.TransactionRef)
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

	balance, err := uc.transactionRepo.Create(ctx, &model.Transaction{
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
		return nil, fmt.Errorf("cannot save transaction: %w", err)
	}

	return &Output{
		Balance:       balance,
		TransactionID: t.ID,
	}, nil
}
