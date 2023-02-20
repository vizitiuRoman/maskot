package withdrawanddeposit

type Input struct {
	ID                   int
	PlayerName           string
	Withdraw             int
	Deposit              int
	Currency             string
	TransactionRef       string
	GameRoundRef         *string
	GameID               *string
	Reason               *string
	SessionID            *string
	SessionAlternativeID *string
	SpinDetailsBetType   string
	SpinDetailsWinType   string
}

type Output struct {
	Balance       int
	TransactionID int
}
