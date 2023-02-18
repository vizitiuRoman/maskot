package handler

import (
	"net/http"

	"github.com/maskot/internal/helpers/use_cases"
	"github.com/maskot/internal/use_cases/seamless/get_balance"
	"github.com/maskot/internal/use_cases/seamless/rollback_transaction"
	"github.com/maskot/internal/use_cases/seamless/withdraw_and_deposit"
	"go.uber.org/zap"
)

type Rpc struct {
	gbUseCase  use_cases.UseCase[get_balance.Input, get_balance.Output]
	wadUseCase use_cases.UseCase[withdraw_and_deposit.Input, withdraw_and_deposit.Output]
	rtUseCase  use_cases.UseCase[rollback_transaction.Input, rollback_transaction.Output]

	log *zap.Logger
}

type RpcConfig struct {
	GbUseCase  use_cases.UseCase[get_balance.Input, get_balance.Output]
	WadUseCase use_cases.UseCase[withdraw_and_deposit.Input, withdraw_and_deposit.Output]
	RtUseCase  use_cases.UseCase[rollback_transaction.Input, rollback_transaction.Output]

	log *zap.Logger
}

func NewRpc(conf *RpcConfig) *Rpc {
	if conf.log == nil {
		conf.log = zap.NewNop()
	}

	return &Rpc{
		gbUseCase:  conf.GbUseCase,
		wadUseCase: conf.WadUseCase,
		rtUseCase:  conf.RtUseCase,
		log:        conf.log.With(zap.String("go.component", "handler.Rpc")),
	}
}

func (h *Rpc) WithdrawAndDeposit(r *http.Request, input *WithdrawAndDepositInput, output *WithdrawAndDepositOutput) error {
	log := h.log.With(zap.String("go.method", "WithdrawAndDeposit"))

	result, err := h.wadUseCase.Execute(r.Context(), &withdraw_and_deposit.Input{
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
		log.Error("failed to withdraw and deposit", zap.Error(err))
		return err
	}

	*output = WithdrawAndDepositOutput{
		NewBalance:    result.Balance,
		TransactionID: result.TransactionID,
	}
	return nil
}

func (h *Rpc) GetBalance(r *http.Request, input *GetBalanceInput, output *GetBalanceOutput) error {
	log := h.log.With(zap.String("go.method", "GetBalance"))

	result, err := h.gbUseCase.Execute(r.Context(), &get_balance.Input{
		PlayerName: input.PlayerName,
	})

	if err != nil {
		log.Error("failed to get balance", zap.Error(err))
		return err
	}

	*output = GetBalanceOutput{
		Balance: result.Balance,
	}
	return nil
}

func (h *Rpc) RollbackTransaction(r *http.Request, input *RollbackTransactionInput, output *RollbackTransactionOutput) error {
	log := h.log.With(zap.String("go.method", "RollbackTransaction"))

	_, err := h.rtUseCase.Execute(r.Context(), &rollback_transaction.Input{
		TransactionRef: input.TransactionRef,
		PlayerName:     input.PlayerName,
	})

	if err != nil {
		log.Error("failed to rollback transaction", zap.Error(err))
		return err
	}

	*output = RollbackTransactionOutput{
		Result: nil,
	}
	return nil
}
