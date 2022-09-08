package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, input *Transaction) (output *Transaction, error error) {
	query := `
		INSERT INTO transactions(currency, player_name, withdraw, deposit, transaction_ref, game_round_ref, game_id, reason, session_id, spin_details_bet_type, spin_details_win_type, session_alternative_id)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        returning
        	id, 
        	player_name, 
        	currency, 
        	withdraw, 
        	deposit, 
        	transaction_ref, 
        	game_round_ref, 
        	game_id, 
        	reason, 
        	session_id, 
        	spin_details_bet_type, 
        	spin_details_win_type
    
	`
	row := r.db.QueryRowContext(ctx, query, input.Currency, input.PlayerName, input.Withdraw, input.Deposit, input.TransactionRef, input.GameRoundRef, input.GameID, input.Reason, input.SessionID, input.SpinDetailsWinType, input.SpinDetailsBetType, input.SessionAlternativeID)

	var t Transaction
	err := row.Scan(
		&t.ID,
		&t.PlayerName,
		&t.Currency,
		&t.Withdraw,
		&t.Deposit,
		&t.TransactionRef,
		&t.GameRoundRef,
		&t.GameID,
		&t.Reason,
		&t.SessionID,
		&t.SpinDetailsWinType,
		&t.SpinDetailsBetType,
	)

	return &t, err
}

func (r *Repository) Rollback(ctx context.Context, id int) error {
	query := `
		UPDATE transactions
    		SET is_rollback = true
    	WHERE id = $1    
	`
	row := r.db.QueryRowContext(ctx, query, id)
	return row.Err()
}

func (r *Repository) Get(ctx context.Context, transactionRef string) (output *Transaction, error error) {
	query := `
		SELECT * FROM transactions
        WHERE transaction_ref = $1 LIMIT 1
    `
	var t Transaction
	err := r.db.GetContext(ctx, &t, query, transactionRef)
	if err != nil {
		return nil, err
	}

	return &t, err
}
