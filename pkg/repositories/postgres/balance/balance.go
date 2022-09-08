package balance

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, input *Balance) (output *Balance, error error) {
	query := `
		INSERT INTO balance(balance, player_name)
        VALUES($1, $2)
        returning
        	player_name, 
        	balance
    
	`
	row := r.db.QueryRowContext(ctx, query, input.Balance, input.PlayerName)

	var b Balance
	err := row.Scan(
		&b.PlayerName,
		&b.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *Repository) Update(ctx context.Context, input *Balance) (output *Balance, error error) {
	query := `
		UPDATE balance
    		SET balance = $1
    	WHERE player_name = $2
        returning
        	player_name, 
        	balance
    
	`
	row := r.db.QueryRowContext(ctx, query, input.Balance, input.PlayerName)

	var b Balance
	err := row.Scan(
		&b.PlayerName,
		&b.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *Repository) Get(ctx context.Context, id string) (output *Balance, error error) {
	query := `
		SELECT * FROM balance
        WHERE player_name = $1 LIMIT 1
    `
	var b Balance
	err := r.db.GetContext(ctx, &b, query, id)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *Repository) GetOrCreate(ctx context.Context, id string) (output *Balance, error error) {
	b, err := r.Get(ctx, id)
	if err == nil {
		return b, nil
	}

	if err == sql.ErrNoRows {
		createdBalance, err := r.Create(ctx, &Balance{PlayerName: id, Balance: 0})
		if err != nil {
			return nil, err
		}

		return createdBalance, nil
	}

	return nil, err
}
