package seamless

import (
	"github.com/maskot/pkg/helpers/use_cases"
	"github.com/maskot/pkg/repositories"
	"github.com/maskot/pkg/use_cases/seamless/get_balance"
	"github.com/maskot/pkg/use_cases/seamless/rollback_transaction"
	"github.com/maskot/pkg/use_cases/seamless/withdraw_and_deposit"
)

type UseCases struct {
	GetBalance          use_cases.UseCase[get_balance.Input, get_balance.Output]
	WithdrawAndDeposit  use_cases.UseCase[withdraw_and_deposit.Input, withdraw_and_deposit.Output]
	RollbackTransaction use_cases.UseCase[rollback_transaction.Input, rollback_transaction.Output]
}

func NewUseCases(repos *repositories.Repository) *UseCases {
	return &UseCases{
		GetBalance:          get_balance.NewUseCase(repos.Balance),
		WithdrawAndDeposit:  withdraw_and_deposit.NewUseCase(repos),
		RollbackTransaction: rollback_transaction.NewUseCase(repos),
	}
}
