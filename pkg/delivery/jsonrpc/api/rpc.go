package api

import (
	"net/http"

	"github.com/maskot/pkg/use_cases/seamless"
	"github.com/maskot/pkg/use_cases/seamless/get_balance"
	"github.com/maskot/pkg/use_cases/seamless/rollback_transaction"
	"github.com/maskot/pkg/use_cases/seamless/withdraw_and_deposit"
)

type Rpc struct {
	seamlessUseCases *seamless.UseCases
}

func NewRpc(seamlessUseCases *seamless.UseCases) *Rpc {
	return &Rpc{
		seamlessUseCases,
	}
}

func (rpc *Rpc) WithdrawAndDeposit(r *http.Request, input *WithdrawAndDepositInput, output *WithdrawAndDepositOutput) error {
	result, err := rpc.seamlessUseCases.WithdrawAndDeposit.Execute(r.Context(), &withdraw_and_deposit.Input{
		Currency:             input.Currency,
		PlayerName:           input.PlayerName,
		Withdraw:             input.Withdraw,
		Deposit:              input.Deposit,
		TransactionRef:       input.TransactionRef,
		GameRoundRef:         input.GameRoundRef,
		GameID:               input.GameID,
		Reason:               input.Reason,
		SessionID:            input.SessionID,
		SessionAlternativeID: input.SessionAlternativeID,
		SpinDetailsBetType:   input.SpinDetails.BetType,
		SpinDetailsWinType:   input.SpinDetails.WinType,
	})

	if err != nil {
		return err
	}

	*output = WithdrawAndDepositOutput{
		NewBalance:    result.Balance,
		TransactionID: result.TransactionID,
	}
	return nil
}

func (rpc *Rpc) GetBalance(r *http.Request, input *GetBalanceInput, output *GetBalanceOutput) error {
	result, err := rpc.seamlessUseCases.GetBalance.Execute(r.Context(), &get_balance.Input{
		PlayerName: input.PlayerName,
	})

	if err != nil {
		return err
	}

	*output = GetBalanceOutput{
		Balance: result.Balance,
	}
	return nil
}

func (rpc *Rpc) RollbackTransaction(r *http.Request, input *RollbackTransactionInput, output *RollbackTransactionOutput) error {
	_, err := rpc.seamlessUseCases.RollbackTransaction.Execute(r.Context(), &rollback_transaction.Input{
		PlayerName:     input.PlayerName,
		TransactionRef: input.TransactionRef,
	})

	if err != nil {
		return err
	}

	*output = RollbackTransactionOutput{
		Result: nil,
	}
	return nil
}
