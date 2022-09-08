package balance

import "time"

type Balance struct {
	ID         int       `db:"id"`
	PlayerName string    `db:"player_name"`
	Balance    int       `db:"balance"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
