package transaction

import "time"

type Transaction struct {
	ID                   int       `db:"id"`
	PlayerName           string    `db:"player_name"`
	Withdraw             int       `db:"withdraw"`
	Deposit              int       `db:"deposit"`
	Currency             string    `db:"currency"`
	TransactionRef       string    `db:"transaction_ref"`
	IsRollback           bool      `db:"is_rollback"`
	SpinDetailsBetType   string    `db:"spin_details_bet_type"`
	SpinDetailsWinType   string    `db:"spin_details_win_type"`
	GameRoundRef         *string   `db:"game_round_ref"`
	GameID               *string   `db:"game_id"`
	Reason               *string   `db:"reason"`
	SessionID            *string   `db:"session_id"`
	SessionAlternativeID *string   `db:"session_alternative_id"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
