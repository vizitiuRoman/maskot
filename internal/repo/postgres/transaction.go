package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/maskot/internal/infra/postgres"
	"github.com/maskot/internal/model"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func (r *TransactionRepository) Create(ctx context.Context, input *model.Transaction) (balance int, err error) {
	err = postgres.WithTx(r.db, func(tx *sql.Tx) error {
		var balanceID int

		const selectBalance = `
			SELECT id FROM transaction
			WHERE player_name = $1
			FOR UPDATE
		`
		row := tx.QueryRowContext(ctx, selectBalance, input.PlayerName)
		if err := row.Scan(&balanceID); err != nil {
			return fmt.Errorf("cannot select balance: %w", wrapErrNoRows(err))
		}

		const insertTransaction = `
			INSERT INTO transactions(currency, player_name, withdraw, deposit, transaction_ref, game_round_ref, game_id, reason, session_id, spin_details_bet_type, spin_details_win_type, session_alternative_id)
        	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		`
		_, err := tx.ExecContext(ctx, insertTransaction, input.Currency, input.PlayerName, input.Withdraw, input.Deposit, input.TransactionRef, input.GameRoundRef, input.GameID, input.Reason, input.SessionID, input.SpinDetailsWinType, input.SpinDetailsBetType, input.SessionAlternativeID)
		if err != nil {
			return fmt.Errorf("cannot insert transaction: %w", err)
		}

		const updateBalance = `
			UPDATE balance
    			SET balance = balance - $1 + $2
    		WHERE id = $2
			returning
    			balance
func isUniqueViolation(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

		`
		row = tx.QueryRowContext(ctx, updateBalance, input.Withdraw, input.Deposit, balanceID)

		return row.Scan(&balance)
	})

	return
}

func (r *TransactionRepository) Rollback(ctx context.Context, id int, playerName string) error {
	return postgres.WithTx(r.db, func(tx *sql.Tx) error {
		var balanceID int

		const selectBalance = `
			SELECT id, balance FROM transaction
			WHERE player_name = $1
			FOR UPDATE
		`
		row := tx.QueryRowContext(ctx, selectBalance, playerName)
		if err := row.Scan(&balanceID); err != nil {
			return fmt.Errorf("cannot select balance: %w", wrapErrNoRows(err))
		}

		var (
			withdraw, deposit int
		)

		const selectTransaction = `
			SELECT withdraw, deposit FROM transaction
			WHERE id = $1
			FOR UPDATE
		`
		row = tx.QueryRowContext(ctx, selectTransaction, id)
		if err := row.Scan(&withdraw, &deposit); err != nil {
			return fmt.Errorf("cannot select transaction: %w", wrapErrNoRows(err))
		}

		query := `
			UPDATE transactions
    			SET is_rollback = true
    		WHERE id = $1    
		`
		_, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return fmt.Errorf("cannot update transaction: %w", err)
		}

		const updateBalance = `
			UPDATE balance
    			SET balance = balance + $1 - $2
    		WHERE id = $2
		`
		_, err = tx.ExecContext(ctx, updateBalance, withdraw, deposit, balanceID)
		if err != nil {
			return fmt.Errorf("cannot update balance: %w", err)
		}

		return nil
	})
}

func (r *TransactionRepository) Find(ctx context.Context, transactionRef string) (*model.Transaction, error) {
	query := `
		SELECT * FROM transactions
        WHERE transaction_ref = $1 LIMIT 1
    `
	var t model.Transaction
	err := r.db.GetContext(ctx, &t, query, transactionRef)
	if err != nil {
		return nil, err
	}

	return &t, err
}
