package api

type GetBalanceInput struct {
	PlayerName string `json:"playerName"`
	Currency   string `json:"currency"`
	GameID     string `json:"gameId"`
}

type GetBalanceOutput struct {
	Balance int `json:"balance"`
}

type WithdrawAndDepositInput struct {
	PlayerName           string      `json:"playerName"`
	Withdraw             int         `json:"withdraw"`
	Deposit              int         `json:"deposit"`
	Currency             string      `json:"currency"`
	TransactionRef       string      `json:"transactionRef"`
	GameRoundRef         *string     `json:"gameRoundRef"`
	GameID               *string     `json:"gameId"`
	Reason               *string     `json:"reason"`
	SessionID            *string     `json:"sessionId"`
	SessionAlternativeID *string     `json:"sessionAlternativeId"`
	SpinDetails          spinDetails `json:"spinDetails"`
}

type spinDetails struct {
	BetType string `json:"betType"`
	WinType string `json:"winType"`
}

type WithdrawAndDepositOutput struct {
	NewBalance    int `json:"newBalance"`
	TransactionID int `json:"transactionId"`
}

type RollbackTransactionInput struct {
	PlayerName     string `json:"playerName"`
	TransactionRef string `json:"transactionRef"`
	GameRoundRef   string `json:"gameRoundRef"`
	GameID         string `json:"gameId"`
	SessionID      string `json:"sessionId"`
}

type RollbackTransactionOutput struct {
	Result *string `json:"result"`
}
