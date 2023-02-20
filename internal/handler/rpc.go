package handler

import (
	"net/http"

	"github.com/maskot/internal/helpers/usecase"
	gb "github.com/maskot/internal/use_cases/seamless/get_balance"
	rb "github.com/maskot/internal/use_cases/seamless/rollback_transaction"
	wad "github.com/maskot/internal/use_cases/seamless/withdraw_and_deposit"
	"go.uber.org/zap"
)

type RPC struct {
	gbUseCase  usecase.UseCase[gb.Input, gb.Output]
	wadUseCase usecase.UseCase[wad.Input, wad.Output]
	rtUseCase  usecase.UseCase[rb.Input, rb.Output]

	log *zap.Logger
}

type RPCConfig struct {
	GbUseCase  usecase.UseCase[gb.Input, gb.Output]
	WadUseCase usecase.UseCase[wad.Input, wad.Output]
	RtUseCase  usecase.UseCase[rb.Input, rb.Output]

	log *zap.Logger
}

func NewRPC(conf *RPCConfig) *RPC {
	if conf.log == nil {
		conf.log = zap.NewNop()
	}

	return &RPC{
		gbUseCase:  conf.GbUseCase,
		wadUseCase: conf.WadUseCase,
		rtUseCase:  conf.RtUseCase,
		log:        conf.log.With(zap.String("go.component", "handler.RPC")),
	}
}

func (h *RPC) WithdrawAndDeposit(r *http.Request, input *WithdrawAndDepositInput, output *WithdrawAndDepositOutput) error {
	log := h.log.With(zap.String("go.method", "WithdrawAndDeposit"))

	result, err := h.wadUseCase.Execute(r.Context(), &wad.Input{
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

func (h *RPC) GetBalance(r *http.Request, input *GetBalanceInput, output *GetBalanceOutput) error {
	log := h.log.With(zap.String("go.method", "GetBalance"))

	result, err := h.gbUseCase.Execute(r.Context(), &gb.Input{
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

func (h *RPC) RollbackTransaction(r *http.Request, input *RollbackTransactionInput, output *RollbackTransactionOutput) error {
	log := h.log.With(zap.String("go.method", "RollbackTransaction"))

	_, err := h.rtUseCase.Execute(r.Context(), &rb.Input{
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
