package model

import "time"

type Transaction struct {
	ID                   int
	PlayerName           string
	Withdraw             int
	Deposit              int
	Currency             string
	TransactionRef       string
	IsRolledBack         bool
	SpinDetailsBetType   string
	SpinDetailsWinType   string
	GameRoundRef         *string
	GameID               *string
	Reason               *string
	SessionID            *string
	SessionAlternativeID *string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
