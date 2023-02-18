package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/maskot/internal/model"
)

type BalanceRepository struct {
	db *sqlx.DB
}

func (r *BalanceRepository) Save(ctx context.Context, input *model.Balance) (output *model.Balance, error error) {
	query := `
		INSERT INTO balance(balance, player_name)
        VALUES($1, $2)
        returning
        	player_name, 
        	balance
    
	`
	row := r.db.QueryRowContext(ctx, query, input.Balance, input.PlayerName)

	var b model.Balance
	err := row.Scan(
		&b.PlayerName,
		&b.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *BalanceRepository) Update(ctx context.Context, input *model.Balance) (output *model.Balance, error error) {
	query := `
		UPDATE balance
    		SET balance = $1
    	WHERE player_name = $2
        returning
        	player_name, 
        	balance
	`
	row := r.db.QueryRowContext(ctx, query, input.Balance, input.PlayerName)

	var b model.Balance
	err := row.Scan(
		&b.PlayerName,
		&b.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *BalanceRepository) Find(ctx context.Context, id string) (output *model.Balance, error error) {
	query := `
		SELECT * FROM balance
        WHERE player_name = $1 LIMIT 1
    `
	var b model.Balance
	err := r.db.GetContext(ctx, &b, query, id)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *BalanceRepository) FindOrCreate(ctx context.Context, id string) (output *model.Balance, error error) {
	b, err := r.Find(ctx, id)
	if err == nil {
		return b, nil
	}

	if err == sql.ErrNoRows {
		createdBalance, err := r.Save(ctx, &model.Balance{PlayerName: id, Balance: 0})
		if err != nil {
			return nil, err
		}

		return createdBalance, nil
	}

	return nil, err
}
