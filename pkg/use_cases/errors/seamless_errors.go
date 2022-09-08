package errors

import "errors"

var (
	ErrNotEnoughMoneyCode     = errors.New("ErrNotEnoughMoneyCode")
	ErrNegativeDepositCode    = errors.New("ErrNegativeDepositCode")
	ErrNegativeWithdrawalCode = errors.New("ErrNegativeWithdrawalCode")
)
